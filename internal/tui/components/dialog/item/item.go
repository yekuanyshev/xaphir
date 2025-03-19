package item

import (
	"github.com/charmbracelet/lipgloss"
	"github.com/charmbracelet/x/ansi"
	"github.com/yekuanyshev/xaphir/internal/tui/components/models"
)

type Item struct {
	models.ChatMessage

	style     lipgloss.Style
	timeStyle lipgloss.Style
}

func NewItem(message models.ChatMessage) Item {
	return Item{
		ChatMessage: message,

		style: lipgloss.NewStyle().
			PaddingLeft(1).PaddingRight(1).
			BorderStyle(lipgloss.RoundedBorder()),

		timeStyle: lipgloss.NewStyle().
			Faint(true),
	}
}

func (i Item) View(width int) string {
	w := width/2 - i.style.GetHorizontalFrameSize()
	content := ansi.Wrap(i.Content, w, "")

	status := ""
	switch {
	case i.IsStatusUnknown():
	case i.IsStatusSent():
		status = lipgloss.NewStyle().Faint(true).Render("✔✔")
	case i.IsStatusRead():
		status = lipgloss.NewStyle().Faint(false).Render("✔✔")
	}

	s := i.style.Render(
		content,
		i.timeStyle.Render(i.SentTime.Format("15:04")),
		status,
	)

	if i.IsFromMe {
		return lipgloss.PlaceHorizontal(width, lipgloss.Right, s)
	}

	return lipgloss.PlaceHorizontal(width, lipgloss.Left, s)
}
