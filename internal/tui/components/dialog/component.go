package dialog

import (
	"time"

	"github.com/charmbracelet/bubbles/textinput"
	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
	"github.com/yekuanyshev/xaphir/internal/tui/components/base"
	"github.com/yekuanyshev/xaphir/internal/tui/components/dialog/item"
	"github.com/yekuanyshev/xaphir/internal/tui/components/events"
	"github.com/yekuanyshev/xaphir/pkg/utils"
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

	sliderAvailableHeight := c.InnerHeight() -
		lipgloss.Height(c.titleStyle.Render(c.title)) -
		lipgloss.Height(c.inputStyle.Render(c.input.View()))

	c.slider.SetHeight(sliderAvailableHeight)
}

func (c *Component) Init() tea.Cmd {
	return nil
}

func (c *Component) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	if msg, ok := c.isFocusCMD(msg); ok {
		c.SetTitle(msg.Title)
		c.slider.SetMessages(msg.Items)
		c.Focus()
		return c, nil
	}

	if !c.Focused() {
		return c, nil
	}

	switch msg := msg.(type) {
	case tea.KeyMsg:
		switch msg.String() {
		case "esc":
			c.Blur()
			return c, events.ChatListFocusCMD()
		case "enter":
			inputValue := c.input.Value()
			if inputValue == "" {
				return c, nil
			}

			message := item.Message{
				Content:  inputValue,
				SendTime: time.Now(),
				IsFromMe: true,
			}
			c.slider.AppendMessage(message)
			c.input.SetValue("")
		case "down":
			c.slider.Increment()
		case "up":
			c.slider.Decrement()
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
		return c.Render(
			lipgloss.Place(
				c.InnerWidth(), c.InnerHeight(),
				lipgloss.Center, lipgloss.Center,
				c.blurredTitleStyle.Render(c.blurredTitle),
			),
		)
	}

	var sections []string
	availHeight := c.InnerHeight()

	titleView := c.titleStyle.Render(c.title)
	sections = append(sections, titleView)
	availHeight -= lipgloss.Height(titleView)

	itemsView := c.itemsView()
	availHeight -= lipgloss.Height(itemsView)

	inputView := c.inputStyle.Render(c.input.View())
	availHeight -= lipgloss.Height(inputView)

	// append empty space
	sections = append(sections, lipgloss.NewStyle().Height(availHeight).Render(""))
	sections = append(sections, itemsView)

	sections = append(sections, inputView)

	return c.Render(
		lipgloss.JoinVertical(
			lipgloss.Left,
			sections...,
		),
	)
}

func (c *Component) SetTitle(title string) {
	c.title = title
}

func (c *Component) itemsView() string {
	itemViews := utils.SliceMap(
		c.slider.GetItems(),
		func(item item.Item) string {
			return item.View()
		},
	)

	return lipgloss.JoinVertical(
		lipgloss.Left,
		itemViews...,
	)
}

func (c *Component) Focus() {
	c.Component.Focus()
	c.input.Focus()
}

func (c *Component) Blur() {
	c.Component.Blur()
	c.input.SetValue("")
}

func (c *Component) isFocusCMD(msg tea.Msg) (events.DialogFocus, bool) {
	event, ok := msg.(events.DialogFocus)
	return event, ok
}
