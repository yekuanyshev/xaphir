package chatlist

import (
	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
	"github.com/yekuanyshev/xaphir/internal/tui/components/base"
	"github.com/yekuanyshev/xaphir/internal/tui/components/chatlist/item"
	"github.com/yekuanyshev/xaphir/internal/tui/components/common"
	"github.com/yekuanyshev/xaphir/internal/tui/components/events"
	"github.com/yekuanyshev/xaphir/pkg/paginator"
	"github.com/yekuanyshev/xaphir/pkg/utils"
)

type Component struct {
	*base.Component

	title     string
	items     []item.Item
	paginator *Paginator[item.Item]

	titleStyle lipgloss.Style
}

func NewComponent(
	chats []item.Chat,
) *Component {
	items := utils.SliceMap(chats, item.NewItem)
	paginator := NewPaginator(paginator.NewItemPaginator(items, 1))

	style := lipgloss.NewStyle().
		PaddingLeft(1).PaddingRight(1).
		BorderStyle(lipgloss.RoundedBorder()).
		BorderForeground(lipgloss.Color("36"))

	titleStyle := lipgloss.NewStyle().
		PaddingLeft(1).PaddingRight(1).
		MarginBottom(1).
		Foreground(lipgloss.Color("#fff")).
		Background(lipgloss.Color("62")).
		Bold(true)

	return &Component{
		Component:  base.NewComponent(base.WithStyle(style)),
		title:      "Chats",
		items:      items,
		paginator:  paginator,
		titleStyle: titleStyle,
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

func (c *Component) SetWidth(width int) {
	c.Component.SetWidth(width)
	c.paginator.SetWidth(c.InnerWidth())
}

func (c *Component) SetHeight(height int) {
	c.Component.SetHeight(height)
	c.paginator.SetLimit(c.calculateLimit())
	for i := range c.items {
		c.items[i].SetSelected(false)
	}
}

func (c *Component) View() string {
	if c.Focused() {
		c.SetStyle(c.Style().Faint(false))
		c.titleStyle = c.titleStyle.Faint(false)
	} else {
		c.SetStyle(c.Style().Faint(true))
		c.titleStyle = c.titleStyle.Faint(true)
	}

	titleView := c.titleStyle.Render(c.title)
	itemsView := c.itemsView()
	paginatorView := c.paginator.View()
	availableHeight := common.CalculateAvailableHeight(
		c.InnerHeight(), titleView, itemsView, paginatorView,
	)
	emptySpace := common.FillWithEmptySpace(availableHeight)

	sections := []string{
		titleView,
		itemsView,
		emptySpace,
		paginatorView,
	}

	return c.Render(
		lipgloss.JoinVertical(
			lipgloss.Left,
			sections...,
		),
	)
}

func (c *Component) itemsView() string {
	itemViews := utils.SliceMap(
		c.paginator.ItemsOnCurrentPage(),
		func(item item.Item) string {
			return item.View(c.InnerWidth())
		},
	)

	return lipgloss.JoinVertical(
		lipgloss.Left,
		itemViews...,
	)
}

func (c *Component) isFocusCMD(msg tea.Msg) bool {
	_, ok := msg.(events.ChatListFocus)
	return ok
}

func (c *Component) calculateLimit() int {
	availHeight := c.getItemsAvailableHeight()
	h := 0
	num := 0

	for _, item := range c.items {
		viewHeight := lipgloss.Height(item.View(c.InnerWidth()))
		if h+viewHeight >= availHeight {
			return num
		}
		h += viewHeight
		num++
	}

	return num
}

func (c *Component) getItemsAvailableHeight() int {
	return c.InnerHeight() -
		lipgloss.Height(c.titleStyle.Render(c.title)) -
		lipgloss.Height(c.paginator.View())
}
