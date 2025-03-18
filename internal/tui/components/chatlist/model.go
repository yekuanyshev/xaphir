package chatlist

import (
	"github.com/charmbracelet/bubbles/key"
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

	var cmds []tea.Cmd

	if c.filtering {
		cmd := c.handleFiltering(msg)
		cmds = append(cmds, cmd)
	}

	previousItemIdx := c.paginator.CurrentIndex()

	switch msg := msg.(type) {
	case tea.KeyMsg:
		switch {
		case key.Matches(msg, c.keyMap.CursorDown) && !c.paginator.IsEmpty():
			c.paginator.Increment()
		case key.Matches(msg, c.keyMap.CursorUp) && !c.paginator.IsEmpty():
			c.paginator.Decrement()
		case key.Matches(msg, c.keyMap.NextPage) && !c.paginator.IsEmpty():
			c.paginator.SkipToNextPage()
		case key.Matches(msg, c.keyMap.PrevPage) && !c.paginator.IsEmpty():
			c.paginator.SkipToPrevPage()
		case key.Matches(msg, c.keyMap.GoToDialog) && !c.paginator.IsEmpty():
			c.Blur()
			currentItem := c.paginator.CurrentItem()
			return c, events.DialogFocusCMD(currentItem.ID)
		case key.Matches(msg, c.keyMap.ShowSearch):
			c.enableFiltering()
		case key.Matches(msg, c.keyMap.CloseSearch) && c.filtering:
			c.disableFiltering()
		}
	}

	currentItemIdx := c.paginator.CurrentIndex()

	if !c.paginator.IsEmpty() {
		previousItem := c.paginator.ItemByIndex(previousItemIdx)
		currentItem := c.paginator.ItemByIndex(currentItemIdx)
		previousItem.Blur()
		currentItem.Focus()
		c.paginator.SetItemOn(previousItemIdx, previousItem)
		c.paginator.SetItemOn(currentItemIdx, currentItem)
	}

	return c, tea.Batch(cmds...)
}

func (c *Component) View() string {
	titleView := c.titleStyle.Render(c.title)
	headerView := titleView

	if c.filtering {
		filterInputView := c.filterInputStyle.Render(
			c.filterInput.View(),
		)
		headerView = filterInputView
	}

	itemsView := c.itemsView()
	paginatorView := c.paginator.View()
	availableHeight := common.CalculateAvailableHeight(
		c.InnerHeight(), headerView, itemsView, paginatorView,
	)
	emptySpace := common.FillWithEmptySpace(availableHeight)

	sections := []string{
		headerView,
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

func (c *Component) HelpView() string {
	return c.help.View()
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

func (c *Component) handleFiltering(msg tea.Msg) tea.Cmd {
	var cmd tea.Cmd

	previousFilteredItems := c.filterItems()

	c.filterInput, cmd = c.filterInput.Update(msg)

	c.applyFiltering(previousFilteredItems)

	return cmd
}
