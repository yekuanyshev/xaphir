package dialog

import (
	"github.com/yekuanyshev/xaphir/internal/tui/components/models"
)

type Item struct {
	models.ChatMessage

	itemView *itemView
}

func newItem(message models.ChatMessage) Item {
	return Item{
		ChatMessage: message,
		itemView:    newItemView(),
	}
}

func (i Item) view(width int) string {
	return i.itemView.render(width, i.ChatMessage)
}
