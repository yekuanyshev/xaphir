package chatlist

import (
	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
	"github.com/yekuanyshev/xaphir/internal/tui/components/base"
	"github.com/yekuanyshev/xaphir/internal/tui/components/chatlist/item"
	"github.com/yekuanyshev/xaphir/internal/tui/components/events"
	"github.com/yekuanyshev/xaphir/pkg/paginator"
	"github.com/yekuanyshev/xaphir/pkg/utils"
)

type Component struct {
	*base.Component

	items     []item.Item
	paginator *Paginator[item.Item]

	style      lipgloss.Style
	titleStyle lipgloss.Style
}

func NewComponent(
	chats []item.Chat,
) *Component {
	items := utils.SliceMap(chats, func(chat item.Chat) item.Item {
		return item.NewItem(chat)
	})
	paginatorLimit := 15

	return &Component{
		Component: base.NewComponent(),

		style: lipgloss.NewStyle().
			PaddingLeft(1).PaddingRight(1).
			BorderStyle(lipgloss.RoundedBorder()).
			BorderForeground(lipgloss.Color("36")),

		titleStyle: lipgloss.NewStyle().
			PaddingLeft(1).PaddingRight(1).
			MarginBottom(1).
			Foreground(lipgloss.Color("#fff")).
			Background(lipgloss.Color("62")).
			Bold(true),

		items:     items,
		paginator: NewPaginator(paginator.NewItemPaginator(items, paginatorLimit)),
	}
}

func (c *Component) Init() tea.Cmd {
	return nil
}

func (c *Component) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	if c.isFocusCMD(msg) {
		c.Focus()
		return c, nil
	}

	if !c.Focused() {
		return c, nil
	}

	previousItemIdx := c.paginator.CurrentIndex()

	switch msg := msg.(type) {
	case tea.KeyMsg:
		switch msg.String() {
		case "down":
			c.paginator.Increment()
		case "up":
			c.paginator.Decrement()
		case "right":
			c.paginator.SkipToNextPage()
		case "left":
			c.paginator.SkipToPrevPage()
		case "enter":
			c.Blur()
			currentItem := c.paginator.CurrentItem()
			return c, events.DialogFocusCMD(
				currentItem.Username,
				currentItem.Messages,
			)
		}
	}

	c.items[previousItemIdx].SetSelected(false)
	c.items[c.paginator.CurrentIndex()].SetSelected(true)

	return c, nil
}

func (c *Component) View() string {
	c.style = c.style.Width(c.Width()).Height(c.Height())
	c.paginator.SetWidth(c.Width() - c.style.GetHorizontalFrameSize())

	if c.Focused() {
		c.style = c.style.Faint(false)
		c.titleStyle = c.titleStyle.Faint(false)
	} else {
		c.style = c.style.Faint(true)
		c.titleStyle = c.titleStyle.Faint(true)
	}

	var sections []string
	availHeight := c.style.GetHeight() - c.style.GetVerticalFrameSize()

	titleView := c.titleStyle.Render("Chats")
	sections = append(sections, titleView)
	availHeight -= lipgloss.Height(titleView)

	itemsView := c.itemsView()
	sections = append(sections, itemsView)
	availHeight -= lipgloss.Height(itemsView)

	paginatorView := c.paginator.View()
	availHeight -= lipgloss.Height(paginatorView)

	// append empty space
	sections = append(sections, lipgloss.NewStyle().Height(availHeight).Render(""))

	sections = append(sections, paginatorView)

	return c.style.Render(
		lipgloss.JoinVertical(
			lipgloss.Left,
			sections...,
		),
	)
}

func (c *Component) itemsView() string {
	w := c.Width() - c.style.GetHorizontalFrameSize()

	itemsOnPage := c.paginator.ItemsOnCurrentPage()

	items := make([]string, 0, len(itemsOnPage))

	for _, chatItem := range itemsOnPage {
		items = append(items, chatItem.View(w))
	}

	return lipgloss.JoinVertical(lipgloss.Left, items...)
}

func (c *Component) isFocusCMD(msg tea.Msg) bool {
	_, ok := msg.(events.ChatListFocus)
	return ok
}
