package dialog

import (
	"github.com/charmbracelet/bubbles/textinput"
	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
	"github.com/yekuanyshev/xaphir/pkg/utils"
)

type Component struct {
	width  int
	height int
	title  string
	items  []MessageItem

	input textinput.Model

	style      lipgloss.Style
	titleStyle lipgloss.Style
	inputStyle lipgloss.Style
}

func NewComponent() *Component {
	input := textinput.New()
	input.Placeholder = "Write a message..."
	input.Blur()

	return &Component{
		input: input,

		style: lipgloss.NewStyle().
			PaddingLeft(1).PaddingRight(1).
			BorderStyle(lipgloss.RoundedBorder()).
			BorderForeground(lipgloss.Color("36")),

		titleStyle: lipgloss.NewStyle().
			PaddingLeft(1).PaddingRight(1).
			MarginBottom(1).
			Foreground(lipgloss.Color("#fff")).
			Background(lipgloss.Color("62")).
			Bold(true),

		inputStyle: lipgloss.NewStyle().
			PaddingLeft(1).PaddingRight(1).
			BorderStyle(lipgloss.RoundedBorder()).
			BorderForeground(lipgloss.Color("36")),
	}
}

func (c *Component) Init() tea.Cmd {
	return nil
}

func (c *Component) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	return c, nil
}

func (c *Component) View() string {
	c.style = c.style.Width(c.width).Height(c.height)
	c.inputStyle = c.inputStyle.Width(c.style.GetWidth() - c.style.GetHorizontalFrameSize() - 2)
	c.input.Width = c.inputStyle.GetWidth()

	var sections []string
	availHeight := c.style.GetHeight() - c.style.GetVerticalFrameSize()

	titleView := c.titleStyle.Render(c.title)
	sections = append(sections, titleView)
	availHeight -= lipgloss.Height(titleView)

	itemsView := c.itemsView()
	sections = append(sections, itemsView)
	availHeight -= lipgloss.Height(itemsView)

	inputView := c.inputStyle.Render(c.input.View())
	availHeight -= lipgloss.Height(inputView)

	// append empty space
	sections = append(sections, lipgloss.NewStyle().Height(availHeight).Render(""))

	sections = append(sections, inputView)

	return c.style.Render(
		lipgloss.JoinVertical(
			lipgloss.Left,
			sections...,
		),
	)
}

func (c *Component) SetWidth(w int) {
	c.width = w
}

func (c *Component) SetHeight(h int) {
	c.height = h
}

func (c *Component) SetTitle(title string) {
	c.title = title
}

func (c *Component) SetItems(items []Message) {
	c.items = utils.SliceMap(items, func(message Message) MessageItem {
		return NewMessageItem(message, c.style.GetWidth()-c.style.GetHorizontalFrameSize())
	})
}

func (c *Component) itemsView() string {
	availHeight := c.style.GetHeight() - c.style.GetVerticalFrameSize()
	availHeight -= lipgloss.Height(c.titleStyle.Render(c.title))
	availHeight -= lipgloss.Height(c.inputStyle.Render(c.input.View()))

	items := make([]string, 0, 20)
	h := 0

	for i := range c.items {
		itemView := c.items[i].View()
		items = append(items, itemView)

		h += lipgloss.Height(itemView)

		if h >= availHeight {
			items = items[:len(items)-1]
			break
		}
	}

	return lipgloss.JoinVertical(
		lipgloss.Left,
		items...,
	)
}
