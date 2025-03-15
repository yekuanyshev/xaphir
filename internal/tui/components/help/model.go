package help

import tea "github.com/charmbracelet/bubbletea"

func (c *Component) Init() tea.Cmd                       { return nil }
func (c *Component) Update(tea.Msg) (tea.Model, tea.Cmd) { return c, nil }

func (c *Component) View() string {
	return c.Render(c.help.View(c.keyMap))
}
