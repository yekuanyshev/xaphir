package dialog

import (
	"time"

	"github.com/charmbracelet/lipgloss"
)

type Message struct {
	Content  string
	SendTime time.Time
	IsFromMe bool
}

type MessageItem struct {
	Message
	width int

	style     lipgloss.Style
	timeStyle lipgloss.Style
}

func NewMessageItem(message Message, width int) MessageItem {
	return MessageItem{
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

func (mi MessageItem) View() string {
	s := mi.style.Render(
		mi.Content,
		mi.timeStyle.Render(mi.SendTime.Format("04:05")),
	)

	w := mi.width - 2

	view := lipgloss.PlaceHorizontal(w, lipgloss.Left, s)

	if mi.IsFromMe {
		view = lipgloss.PlaceHorizontal(w, lipgloss.Right, s)
	}

	return view
}
