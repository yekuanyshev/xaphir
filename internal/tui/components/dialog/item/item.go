package item

import (
	"time"

	"github.com/charmbracelet/lipgloss"
	"github.com/charmbracelet/x/ansi"
)

type MessageStatus int

const (
	MessageStatusUnknown = iota
	MessageStatusSent
	MessageStatusRead
)

type Message struct {
	Content  string
	SendTime time.Time
	IsFromMe bool
	Status   int
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
			PaddingLeft(1).PaddingRight(1).
			BorderStyle(lipgloss.RoundedBorder()),

		timeStyle: lipgloss.NewStyle().
			Faint(true),
	}
}

func (i Item) View() string {
	w := i.width/2 - i.style.GetHorizontalFrameSize()
	content := ansi.Wrap(i.Content, w, "")

	status := ""
	switch i.Status {
	case MessageStatusUnknown:
	case MessageStatusSent:
		status = lipgloss.NewStyle().Faint(true).Render("✔✔")
	case MessageStatusRead:
		status = lipgloss.NewStyle().Faint(false).Render("✔✔")
	}

	s := i.style.Render(
		content,
		i.timeStyle.Render(i.SendTime.Format("15:04")),
		status,
	)

	if i.IsFromMe {
		return lipgloss.PlaceHorizontal(i.width, lipgloss.Right, s)
	}

	return lipgloss.PlaceHorizontal(i.width, lipgloss.Left, s)
}

func (i *Item) SetWidth(width int) {
	i.width = width
}
