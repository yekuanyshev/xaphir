package dialog

import (
	"github.com/charmbracelet/bubbles/textinput"
	"github.com/charmbracelet/lipgloss"
	"github.com/yekuanyshev/xaphir/internal/tui/components/base"
	"github.com/yekuanyshev/xaphir/internal/tui/components/common"
	"github.com/yekuanyshev/xaphir/internal/tui/components/dialog/item"
	"github.com/yekuanyshev/xaphir/internal/tui/components/help"
)

type Component struct {
	*base.Component

	title        string
	blurredTitle string

	slider *Slider

	input                   textinput.Model
	focusedInputPlaceholder string
	blurredInputPlaceholder string

	titleStyle        lipgloss.Style
	blurredTitleStyle lipgloss.Style
	inputStyle        lipgloss.Style

	keyMap KeyMap
	help   *help.Component
}

func NewComponent() *Component {
	input := textinput.New()
	focusedInputPlaceholder := "Write a message..."
	blurredInputPlaceholder := "Press enter to type a message..."
	input.Placeholder = blurredInputPlaceholder
	input.Blur()

	blurredTitle := "Select a chat to start messaging..."

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

	blurredTitleStyle := lipgloss.NewStyle().
		Faint(true)

	inputStyle := lipgloss.NewStyle().
		PaddingLeft(1).PaddingRight(1).
		BorderStyle(lipgloss.RoundedBorder())

	keyMap := DefaultKeyMap()
	help := help.New(keyMap)

	return &Component{
		Component:               base.NewComponent(base.WithStyle(style)),
		title:                   "",
		blurredTitle:            blurredTitle,
		slider:                  NewSlider(),
		input:                   input,
		focusedInputPlaceholder: focusedInputPlaceholder,
		blurredInputPlaceholder: blurredInputPlaceholder,
		titleStyle:              titleStyle,
		blurredTitleStyle:       blurredTitleStyle,
		inputStyle:              inputStyle,
		keyMap:                  keyMap,
		help:                    help,
	}
}

func (c *Component) SetWidth(width int) {
	c.Component.SetWidth(width)
	c.slider.SetWidth(c.InnerWidth())
	c.inputStyle = c.inputStyle.Width(c.InnerWidth() - c.inputStyle.GetHorizontalFrameSize())
}

func (c *Component) SetHeight(height int) {
	c.Component.SetHeight(height)

	availableHeight := common.CalculateAvailableHeight(
		c.InnerHeight(),
		c.titleStyle.Render(c.title),
		c.inputStyle.Render(c.input.View()),
	)

	c.slider.SetHeight(availableHeight)
}

func (c *Component) Blur() {
	c.Component.Blur()
	c.input.SetValue("")
}

func (c *Component) SetTitle(title string) {
	c.title = title
}

func (c *Component) SetSliderMessages(messages []item.Message) {
	c.slider.SetMessages(messages)
}

func (c *Component) inputFocus() {
	c.input.Focus()
	c.input.Placeholder = c.focusedInputPlaceholder
	c.inputStyle = c.inputStyle.BorderForeground(lipgloss.Color("36"))
}

func (c *Component) inputBlur() {
	c.input.Blur()
	c.input.Placeholder = c.blurredInputPlaceholder
	c.inputStyle = c.inputStyle.UnsetBorderForeground()
}

func (c *Component) IsTypingMessage() bool {
	return c.input.Focused()
}
