package chatlist

import (
	"slices"

	"github.com/charmbracelet/lipgloss"
	"github.com/yekuanyshev/xaphir/internal/tui/components/common"
	"github.com/yekuanyshev/xaphir/internal/tui/components/help"
	"github.com/yekuanyshev/xaphir/internal/tui/components/models"
	"github.com/yekuanyshev/xaphir/pkg/paginator"
	"github.com/yekuanyshev/xaphir/pkg/utils"
)

type Component struct {
	focus bool

	title     string
	items     []Item
	paginator *paginator.ItemPaginator[Item]

	keyMap KeyMap
	help   *help.Component

	filter *filter
	view   *view
}

func NewComponent() *Component {
	keyMap := DefaultKeyMap()

	return &Component{
		title:     "Chats",
		items:     nil,
		paginator: nil,
		keyMap:    keyMap,
		help:      help.New(keyMap),
		filter:    newFilter(),
		view:      newView(),
	}
}

func (c *Component) SetWidth(width int) {
	c.view.setWidth(width)
	c.filter.setWidth(c.view.innerWidth())
}

func (c *Component) SetHeight(height int) {
	c.view.setHeight(height)
	c.paginator.SetLimit(c.calculateLimit())
	c.blurItems()
}

func (c *Component) Focus() {
	c.focus = true
	c.view.focus()
}

func (c *Component) Blur() {
	c.focus = false
	c.view.blur()
}

func (c *Component) Focused() bool {
	return c.focus
}

func (c *Component) SetItems(chats []models.Chat) {
	if c.filter.enabled {
		c.disableFiltering()
	}

	c.items = utils.SliceMap(chats, newItem)
	if c.paginator == nil {
		// limit for initializing, correct limit will be after c.SetHeight()
		initialLimit := 0
		c.paginator = paginator.NewItemPaginator(c.items, initialLimit)
	}
	c.paginator.SetItems(c.items)
	if len(c.items) > 0 {
		c.items[0].focus()
	}
}

func (c *Component) calculateLimit() int {
	availableHeight := common.CalculateAvailableHeight(
		c.view.innerHeight(),
		c.view.renderTitle(c.title),
		c.view.renderPaginator(c.paginator),
	)
	h := 0
	num := 0

	for _, item := range c.items {
		viewHeight := lipgloss.Height(c.view.renderItem(item))
		if h+viewHeight >= availableHeight {
			return num
		}
		h += viewHeight
		num++
	}

	return num
}

func (c *Component) applyFiltering(previousFilteredItems []Item) {
	filteredItems := c.filter.filterItems(c.items)

	if !slices.EqualFunc(
		previousFilteredItems,
		filteredItems,
		func(item1, item2 Item) bool { return item1.ID == item2.ID },
	) {
		c.paginator.SetItems(filteredItems)
	}
}

func (c *Component) enableFiltering() {
	c.filter.enable()
	c.paginator.SetItems(nil)
	c.blurItems()
}

func (c *Component) disableFiltering() {
	c.filter.disable()
	c.paginator.SetItems(c.items)
	c.blurItems()
}

func (c *Component) blurItems() {
	for i := range c.items {
		c.items[i].blur()
	}
}
