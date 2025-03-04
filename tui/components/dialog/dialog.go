package dialog

import (
	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
	"github.com/yekuanyshev/xaphir/pkg/utils"
)

type Component struct {
	width  int
	height int
	title  string
	items  []MessageItem

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

	itemsView := c.itemsView()
	sections = append(sections, itemsView)

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
		return NewMessageItem(message, c.width)
	})
}

func (c *Component) itemsView() string {
	avalHeight := c.height - c.titleStyle.GetHeight() - c.titleStyle.GetVerticalBorderSize()

	items := make([]string, 0, 20)
	block := ""
	h := 0

	for _, item := range c.items {
		itemView := item.View()
		items = append(items, itemView)

		h += lipgloss.Height(itemView)

		if h >= avalHeight {
			block = lipgloss.JoinVertical(lipgloss.Left, items[:len(items)-1]...)
			break
		}
	}

	if block == "" {
		block = lipgloss.JoinVertical(lipgloss.Left, items...)
	}

	return block
}
