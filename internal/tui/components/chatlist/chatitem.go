package chatlist

import (
	"github.com/charmbracelet/lipgloss"
	"github.com/charmbracelet/x/ansi"
	"github.com/yekuanyshev/xaphir/internal/tui/components/dialog"
)

type Chat struct {
	Username    string
	LastMessage string
	Messages    []dialog.Message
}

type ChatItem struct {
	Chat
	selected bool

	style                    lipgloss.Style
	titleStyle               lipgloss.Style
	descriptionStyle         lipgloss.Style
	selectedTitleStyle       lipgloss.Style
	selectedDescriptionStyle lipgloss.Style
}

func NewChatItem(chat Chat) ChatItem {
	return ChatItem{
		Chat: chat,

		style: lipgloss.NewStyle().
			MarginBottom(1),

		titleStyle: lipgloss.NewStyle().
			Foreground(lipgloss.Color("#fff")).
			Bold(true),

		descriptionStyle: lipgloss.NewStyle().
			Faint(true),

		selectedTitleStyle: lipgloss.NewStyle().
			Foreground(lipgloss.Color("62")).
			Bold(true),

		selectedDescriptionStyle: lipgloss.NewStyle().
			Faint(true),
	}
}

func (ci ChatItem) View(width int) string {
	ci.LastMessage = ansi.Truncate(ci.LastMessage, width, "...")

	if ci.selected {
		return ci.style.Render(
			lipgloss.JoinVertical(
				lipgloss.Left,
				ci.selectedTitleStyle.Render(ci.Username),
				ci.selectedDescriptionStyle.Render(ci.LastMessage),
			),
		)
	}

	return ci.style.Render(
		lipgloss.JoinVertical(
			lipgloss.Left,
			ci.titleStyle.Render(ci.Username),
			ci.descriptionStyle.Render(ci.LastMessage),
		),
	)
}

func (ci *ChatItem) SetSelected(selected bool) {
	ci.selected = selected
}
