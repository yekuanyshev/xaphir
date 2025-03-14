package chatlist

import (
	"github.com/charmbracelet/lipgloss"
	"github.com/yekuanyshev/xaphir/internal/tui/components/base"
	"github.com/yekuanyshev/xaphir/internal/tui/components/chatlist/item"
	"github.com/yekuanyshev/xaphir/internal/tui/components/common"
	"github.com/yekuanyshev/xaphir/pkg/utils"
)

type Component struct {
	*base.Component

	title     string
	items     []item.Item
	paginator *Paginator[item.Item]

	titleStyle lipgloss.Style
}

func NewComponent(
	chats []item.Chat,
) *Component {
	items := utils.SliceMap(chats, item.NewItem)
	paginator := NewPaginator(items)

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

	return &Component{
		Component:  base.NewComponent(base.WithStyle(style)),
		title:      "Chats",
		items:      items,
		paginator:  paginator,
		titleStyle: titleStyle,
	}
}

func (c *Component) SetWidth(width int) {
	c.Component.SetWidth(width)
	c.paginator.SetWidth(c.InnerWidth())
}

func (c *Component) SetHeight(height int) {
	c.Component.SetHeight(height)
	c.paginator.SetLimit(c.calculateLimit())
	for i := range c.items {
		c.items[i].Blur()
	}
}

func (c *Component) Focus() {
	c.Component.Focus()
	c.SetStyle(c.Style().Faint(false))
	c.titleStyle = c.titleStyle.Faint(false)
}

func (c *Component) Blur() {
	c.Component.Blur()
	c.SetStyle(c.Style().Faint(true))
	c.titleStyle = c.titleStyle.Faint(true)
}

func (c *Component) calculateLimit() int {
	availableHeight := common.CalculateAvailableHeight(
		c.InnerHeight(),
		c.titleStyle.Render(c.title),
		c.paginator.View(),
	)
	h := 0
	num := 0

	for _, item := range c.items {
		viewHeight := lipgloss.Height(item.View(c.InnerWidth()))
		if h+viewHeight >= availableHeight {
			return num
		}
		h += viewHeight
		num++
	}

	return num
}
