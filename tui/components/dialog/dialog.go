package dialog

import (
	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
)

type Component struct {
	width  int
	height int
	title  string

	style      lipgloss.Style
	titleStyle lipgloss.Style
}

func NewComponent() *Component {
	return &Component{
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

	var sections []string

	titleView := c.titleStyle.Render(c.title)
	sections = append(sections, titleView)

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
