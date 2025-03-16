package chatlist

import (
	"slices"
	"strings"

	"github.com/charmbracelet/bubbles/textinput"
	"github.com/charmbracelet/lipgloss"
	"github.com/yekuanyshev/xaphir/internal/tui/components/base"
	"github.com/yekuanyshev/xaphir/internal/tui/components/chatlist/item"
	"github.com/yekuanyshev/xaphir/internal/tui/components/common"
	"github.com/yekuanyshev/xaphir/internal/tui/components/help"
	"github.com/yekuanyshev/xaphir/pkg/utils"
)

type Component struct {
	*base.Component

	title     string
	items     []item.Item
	paginator *Paginator[item.Item]

	titleStyle lipgloss.Style

	keyMap KeyMap
	help   *help.Component

	filtering        bool
	filterInput      textinput.Model
	filterInputStyle lipgloss.Style
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

	keyMap := DefaultKeyMap()
	help := help.New(keyMap)

	filterInput := textinput.New()
	filterInput.Placeholder = "Search..."
	filterInputStyle := lipgloss.NewStyle().
		PaddingLeft(1).PaddingRight(1).
		BorderStyle(lipgloss.RoundedBorder()).
		BorderForeground(lipgloss.Color("36"))

	return &Component{
		Component:        base.NewComponent(base.WithStyle(style)),
		title:            "Chats",
		items:            items,
		paginator:        paginator,
		titleStyle:       titleStyle,
		keyMap:           keyMap,
		help:             help,
		filtering:        false,
		filterInput:      filterInput,
		filterInputStyle: filterInputStyle,
	}
}

func (c *Component) SetWidth(width int) {
	c.Component.SetWidth(width)
	c.paginator.SetWidth(c.InnerWidth())
	c.filterInput.Width = c.InnerWidth()
	c.filterInput.CharLimit = 32
	c.filterInputStyle = c.filterInputStyle.Width(
		c.InnerWidth() - c.filterInputStyle.GetHorizontalFrameSize(),
	)
}

func (c *Component) SetHeight(height int) {
	c.Component.SetHeight(height)
	c.paginator.SetLimit(c.calculateLimit())
	c.blurItems()
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

func (c *Component) applyFiltering(previousFilteredItems []item.Item) {
	filteredItems := c.filterItems()

	if !slices.EqualFunc(
		previousFilteredItems,
		filteredItems,
		item.ItemEquals,
	) {
		c.paginator.SetItems(filteredItems)
	}
}

func (c *Component) enableFiltering() {
	c.filtering = true
	c.filterInput.Focus()
	c.paginator.SetItems(c.filterItems())
	c.blurItems()
}

func (c *Component) disableFiltering() {
	c.filtering = false
	c.filterInput.Blur()
	c.filterInput.SetValue("")
	c.paginator.SetItems(c.items)
	c.blurItems()
}

func (c *Component) filterItems() []item.Item {
	inputValue := c.filterInput.Value()
	inputValue = strings.TrimSpace(inputValue)
	inputValue = strings.ToLower(inputValue)

	if inputValue == "" {
		return nil
	}

	return utils.SliceFilter(c.items, func(item item.Item) bool {
		username := strings.ToLower(item.Username)
		return strings.Contains(username, inputValue)
	})
}

func (c *Component) blurItems() {
	for i := range c.items {
		c.items[i] = c.items[i].Blur()
	}
}
