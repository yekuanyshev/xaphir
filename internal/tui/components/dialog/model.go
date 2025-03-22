package dialog

import (
	"strings"

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
	return c.view.render(
		c.focus,
		c.blurredTitle,
		c.title,
		c.slider,
		c.input,
	)
}

func (c *Component) HelpView() string {
	return c.help.View()
}
