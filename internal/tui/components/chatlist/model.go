package chatlist

import (
	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
	"github.com/yekuanyshev/xaphir/internal/tui/components/chatlist/item"
	"github.com/yekuanyshev/xaphir/internal/tui/components/common"
	"github.com/yekuanyshev/xaphir/internal/tui/components/events"
	"github.com/yekuanyshev/xaphir/pkg/utils"
)

func (c *Component) Init() tea.Cmd {
	return nil
}

func (c *Component) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
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

	currentItemIdx := c.paginator.CurrentIndex()

	c.items[previousItemIdx].Blur()
	c.items[currentItemIdx].Focus()

	return c, nil
}

func (c *Component) View() string {
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
