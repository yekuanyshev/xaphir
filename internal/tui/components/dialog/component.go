package dialog

import (
	"github.com/charmbracelet/bubbles/textinput"
	"github.com/yekuanyshev/xaphir/internal/tui/components/common"
	"github.com/yekuanyshev/xaphir/internal/tui/components/help"
	"github.com/yekuanyshev/xaphir/internal/tui/components/models"
)

type Component struct {
	chatID int64

	focus bool

	title        string
	blurredTitle string

	slider *Slider

	input                   textinput.Model
	focusedInputPlaceholder string
	blurredInputPlaceholder string

	keyMap KeyMap
	help   *help.Component
	view   *view
}

func NewComponent() *Component {
	input := textinput.New()
	focusedInputPlaceholder := "Write a message..."
	blurredInputPlaceholder := "Press enter to type a message..."
	input.Placeholder = blurredInputPlaceholder
	input.Blur()

	blurredTitle := "Select a chat to start messaging..."

	keyMap := DefaultKeyMap()
	help := help.New(keyMap)

	return &Component{
		title:                   "",
		blurredTitle:            blurredTitle,
		slider:                  NewSlider(),
		input:                   input,
		focusedInputPlaceholder: focusedInputPlaceholder,
		blurredInputPlaceholder: blurredInputPlaceholder,
		keyMap:                  keyMap,
		help:                    help,
		view:                    newView(),
	}
}

func (c *Component) SetWidth(width int) {
	c.view.setWidth(width)
	c.slider.SetWidth(c.view.innerWidth())
}

func (c *Component) SetHeight(height int) {
	c.view.setHeight(height)

	availableHeight := common.CalculateAvailableHeight(
		c.view.innerHeight(),
		c.view.renderTitle(c.title),
		c.view.renderInput(c.input),
	)

	c.slider.SetHeight(availableHeight)
}

func (c *Component) Focus() {
	c.focus = true
}

func (c *Component) Blur() {
	c.focus = false
	c.input.SetValue("")
}

func (c *Component) Focused() bool {
	return c.focus
}

func (c *Component) SetTitle(title string) {
	c.title = title
}

func (c *Component) SetSliderMessages(messages []models.ChatMessage) {
	c.slider.SetMessages(messages)
}

func (c *Component) SetChatID(chatID int64) {
	c.chatID = chatID
}

func (c *Component) inputFocus() {
	c.input.Focus()
	c.input.Placeholder = c.focusedInputPlaceholder
	c.view.inputFocus()
}

func (c *Component) inputBlur() {
	c.input.Blur()
	c.input.Placeholder = c.blurredInputPlaceholder
	c.view.inputBlur()
}

func (c *Component) IsTypingMessage() bool {
	return c.input.Focused()
}
