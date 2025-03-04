package chatlist

import "github.com/charmbracelet/lipgloss"

type Chat struct {
	Username    string
	LastMessage string
}

type ChatItem struct {
	Chat

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

func (ci ChatItem) View(isSelected bool) string {
	view := lipgloss.JoinVertical(
		lipgloss.Left,
		ci.titleStyle.Render(ci.Username),
		ci.descriptionStyle.Render(ci.LastMessage),
	)

	if isSelected {
		view = lipgloss.JoinVertical(
			lipgloss.Left,
			ci.selectedTitleStyle.Render(ci.Username),
			ci.selectedDescriptionStyle.Render(ci.LastMessage),
		)
	}

	return ci.style.Render(view)
}
