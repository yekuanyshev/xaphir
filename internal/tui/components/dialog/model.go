package dialog

import (
	"strings"

	"github.com/charmbracelet/bubbles/key"
	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
	"github.com/yekuanyshev/xaphir/internal/tui/components/common"
	"github.com/yekuanyshev/xaphir/internal/tui/components/dialog/item"
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

	switch msg := msg.(type) {
	case tea.KeyMsg:
		switch {
		case key.Matches(msg, c.keyMap.CursorDown):
			c.slider.Increment()
		case key.Matches(msg, c.keyMap.CursorUp):
			c.slider.Decrement()
		case key.Matches(msg, c.keyMap.BackToChats):
			if c.IsTypingMessage() {
				c.inputBlur()
				return c, nil
			}

			c.Blur()
			return c, events.ChatListFocusCMD()
		case key.Matches(msg, c.keyMap.SendMessage):
			if !c.IsTypingMessage() {
				c.inputFocus()
				return c, nil
			}

			inputValue := strings.TrimSpace(c.input.Value())
			if inputValue == "" {
				return c, nil
			}

			c.input.SetValue("")
			return c, events.SendMessageCMD(c.chatID, inputValue)
		}
	}

	var inputCMD tea.Cmd
	c.input, inputCMD = c.input.Update(msg)

	return c, tea.Batch(
		inputCMD,
	)
}

func (c *Component) View() string {
	if !c.Focused() {
		return c.blurredView()
	}

	titleView := c.titleStyle.Render(c.title)
	itemsView := c.itemsView()
	inputView := c.inputStyle.Render(c.input.View())
	availableHeight := common.CalculateAvailableHeight(
		c.InnerHeight(), titleView, itemsView, inputView,
	)
	emptySpace := common.FillWithEmptySpace(availableHeight)

	sections := []string{
		titleView,
		emptySpace,
		itemsView,
		inputView,
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

func (c *Component) blurredView() string {
	return c.Render(
		lipgloss.Place(
			c.InnerWidth(), c.InnerHeight(),
			lipgloss.Center, lipgloss.Center,
			c.blurredTitleStyle.Render(c.blurredTitle),
		),
	)
}

func (c *Component) itemsView() string {
	itemViews := utils.SliceMap(
		c.slider.GetItems(),
		func(item item.Item) string {
			return item.View(c.InnerWidth())
		},
	)

	return lipgloss.JoinVertical(
		lipgloss.Left,
		itemViews...,
	)
}
