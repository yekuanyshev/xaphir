package chatlist

import (
	"github.com/yekuanyshev/xaphir/internal/tui/components/models"
)

type Item struct {
	models.Chat
	focused bool

	itemView *itemView
}

func newItem(chat models.Chat) Item {
	return Item{
		Chat:     chat,
		focused:  false,
		itemView: newItemView(),
	}
}

func (i Item) view(width int) string {
	return i.itemView.render(width, i.Chat, i.focused)
}

func (i *Item) focus() { i.focused = true }
func (i *Item) blur()  { i.focused = false }
