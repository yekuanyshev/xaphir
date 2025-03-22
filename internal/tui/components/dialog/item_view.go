package dialog

import (
	"github.com/charmbracelet/lipgloss"
	"github.com/charmbracelet/x/ansi"
	"github.com/yekuanyshev/xaphir/internal/tui/components/models"
)

type itemView struct {
	style     lipgloss.Style
	timeStyle lipgloss.Style
}

func newItemView() *itemView {
	style := lipgloss.NewStyle().
		PaddingLeft(1).PaddingRight(1).
		BorderStyle(lipgloss.RoundedBorder())

	timeStyle := lipgloss.NewStyle().
		Faint(true)

	return &itemView{
		style:     style,
		timeStyle: timeStyle,
	}
}

func (v *itemView) render(
	width int,
	chatMessage models.ChatMessage,
) string {
	w := width/2 - v.style.GetHorizontalFrameSize()
	content := ansi.Wrap(chatMessage.Content, w, "")

	status := ""
	switch {
	case chatMessage.IsStatusUnknown():
	case chatMessage.IsStatusSent():
		status = lipgloss.NewStyle().Faint(true).Render("✔✔")
	case chatMessage.IsStatusRead():
		status = lipgloss.NewStyle().Faint(false).Render("✔✔")
	}

	s := v.style.Render(
		content,
		v.timeStyle.Render(chatMessage.SentTime.Format("15:04")),
		status,
	)

	if chatMessage.IsFromMe {
		return lipgloss.PlaceHorizontal(width, lipgloss.Right, s)
	}

	return lipgloss.PlaceHorizontal(width, lipgloss.Left, s)
}
