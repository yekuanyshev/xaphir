package chatlist

import (
	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
	"github.com/yekuanyshev/xaphir/internal/tui/components/dialog"
	"github.com/yekuanyshev/xaphir/pkg/paginator"
	"github.com/yekuanyshev/xaphir/pkg/utils"
)

type Component struct {
	width     int
	height    int
	items     []ChatItem
	paginator *Paginator[ChatItem]

	dialog *dialog.Component

	style      lipgloss.Style
	titleStyle lipgloss.Style
}

func NewComponent(
	chats []Chat,
	dialog *dialog.Component,
) *Component {
	items := utils.SliceMap(chats, func(chat Chat) ChatItem {
		return NewChatItem(chat)
	})
	paginatorLimit := 15

	return &Component{
		dialog: dialog,

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
		}
	}

	c.items[previousItemIdx].SetSelected(false)
	c.items[c.paginator.CurrentIndex()].SetSelected(true)

	c.dialog.SetTitle(c.paginator.CurrentItem().Username)
	c.dialog.SetItems(c.paginator.CurrentItem().Messages)

	return c, nil
}

func (c *Component) View() string {
	c.style = c.style.Width(c.width).Height(c.height)
	c.paginator.SetWidth(c.width - c.style.GetHorizontalFrameSize())

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

func (c *Component) SetWidth(w int) {
	c.width = w
}

func (c *Component) SetHeight(h int) {
	c.height = h
}

func (c *Component) itemsView() string {
	w := c.width - c.style.GetHorizontalFrameSize()

	itemsOnPage := c.paginator.ItemsOnCurrentPage()

	items := make([]string, 0, len(itemsOnPage))

	for _, chatItem := range itemsOnPage {
		items = append(items, chatItem.View(w))
	}

	return lipgloss.JoinVertical(lipgloss.Left, items...)
}
