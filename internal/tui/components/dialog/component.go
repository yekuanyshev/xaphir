package dialog

import (
	"github.com/charmbracelet/bubbles/textinput"
	"github.com/charmbracelet/lipgloss"
	"github.com/yekuanyshev/xaphir/internal/tui/components/base"
	"github.com/yekuanyshev/xaphir/internal/tui/components/common"
	"github.com/yekuanyshev/xaphir/internal/tui/components/dialog/item"
)

type Component struct {
	*base.Component

	title        string
	blurredTitle string

	slider *Slider

	input textinput.Model

	titleStyle        lipgloss.Style
	blurredTitleStyle lipgloss.Style
	inputStyle        lipgloss.Style
}

func NewComponent() *Component {
	input := textinput.New()
	input.Placeholder = "Write a message..."
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
		BorderStyle(lipgloss.RoundedBorder()).
		BorderForeground(lipgloss.Color("36"))

	return &Component{
		Component:         base.NewComponent(base.WithStyle(style)),
		title:             "",
		blurredTitle:      blurredTitle,
		slider:            NewSlider(),
		input:             input,
		titleStyle:        titleStyle,
		blurredTitleStyle: blurredTitleStyle,
		inputStyle:        inputStyle,
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

func (c *Component) Focus() {
	c.Component.Focus()
	c.input.Focus()
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
