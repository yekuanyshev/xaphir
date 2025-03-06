package item

import (
	"time"

	"github.com/charmbracelet/lipgloss"
)

type Message struct {
	Content  string
	SendTime time.Time
	IsFromMe bool
}

type Item struct {
	Message
	width int

	style     lipgloss.Style
	timeStyle lipgloss.Style
}

func NewItem(message Message, width int) Item {
	return Item{
		Message: message,
		width:   width,

		style: lipgloss.NewStyle().
			MaxWidth(width / 2).
			PaddingLeft(1).PaddingRight(1).
			BorderStyle(lipgloss.RoundedBorder()),

		timeStyle: lipgloss.NewStyle().
			Faint(true),
	}
}

func (i Item) View() string {
	s := i.style.Render(
		i.Content,
		i.timeStyle.Render(i.SendTime.Format("15:04")),
	)

	if i.IsFromMe {
		return lipgloss.PlaceHorizontal(i.width, lipgloss.Right, s)
	}

	return lipgloss.PlaceHorizontal(i.width, lipgloss.Left, s)

}
