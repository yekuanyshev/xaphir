package chatlist

import (
	"github.com/charmbracelet/bubbles/key"
	tea "github.com/charmbracelet/bubbletea"
	"github.com/yekuanyshev/xaphir/internal/tui/components/events"
)

func (c *Component) Init() tea.Cmd {
	return nil
}

func (c *Component) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	if !c.focus {
		return c, nil
	}

	var cmds []tea.Cmd

	if c.filter.enabled {
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
		case key.Matches(msg, c.keyMap.CloseSearch) && c.filter.enabled:
			c.disableFiltering()
		}
	}

	currentItemIdx := c.paginator.CurrentIndex()

	if !c.paginator.IsEmpty() {
		previousItem := c.paginator.ItemByIndex(previousItemIdx)
		currentItem := c.paginator.ItemByIndex(currentItemIdx)
		previousItem.blur()
		currentItem.focus()
		c.paginator.SetItemOn(previousItemIdx, previousItem)
		c.paginator.SetItemOn(currentItemIdx, currentItem)
	}

	return c, tea.Batch(cmds...)
}

func (c *Component) View() string {
	return c.view.render(
		c.title,
		c.filter,
		c.paginator,
	)
}

func (c *Component) HelpView() string {
	return c.help.View()
}

func (c *Component) handleFiltering(msg tea.Msg) tea.Cmd {
	var cmd tea.Cmd

	previousFilteredItems := c.filter.filterItems(c.items)

	c.filter, cmd = c.filter.update(msg)

	c.applyFiltering(previousFilteredItems)

	return cmd
}
