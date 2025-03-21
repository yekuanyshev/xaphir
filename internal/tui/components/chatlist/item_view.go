package chatlist

import (
	"github.com/charmbracelet/lipgloss"
	"github.com/charmbracelet/x/ansi"
	"github.com/yekuanyshev/xaphir/internal/tui/components/common"
	"github.com/yekuanyshev/xaphir/internal/tui/components/models"
)

type itemView struct {
	style                    lipgloss.Style
	titleStyle               lipgloss.Style
	descriptionStyle         lipgloss.Style
	selectedTitleStyle       lipgloss.Style
	selectedDescriptionStyle lipgloss.Style
	lastMessageTimeStyle     lipgloss.Style
}

func newItemView() *itemView {
	style := lipgloss.NewStyle().
		MarginBottom(1)

	titleStyle := lipgloss.NewStyle().
		Foreground(lipgloss.Color("#fff")).
		Bold(true)

	descriptionStyle := lipgloss.NewStyle().
		Faint(true)

	selectedTitleStyle := lipgloss.NewStyle().
		Foreground(lipgloss.Color("62")).
		Bold(true)

	selectedDescriptionStyle := lipgloss.NewStyle().
		Faint(true)

	lastMessageTimeStyle := lipgloss.NewStyle().Faint(true)

	return &itemView{
		style:                    style,
		titleStyle:               titleStyle,
		descriptionStyle:         descriptionStyle,
		selectedTitleStyle:       selectedTitleStyle,
		selectedDescriptionStyle: selectedDescriptionStyle,
		lastMessageTimeStyle:     lastMessageTimeStyle,
	}
}

func (v *itemView) render(
	width int,
	chat models.Chat,
	focus bool,
) string {
	var lastMessageContent string
	var lastMessageSentTime string

	if chat.LastMessage != nil {
		lastMessageContent = ansi.Truncate(chat.LastMessage.Content, width, "...")
		lastMessageSentTime = chat.LastMessage.FormatSentTime()
	}

	availableWidth := common.CalculateAvailableWidth(
		width,
		v.selectedTitleStyle.Render(chat.Username),
		v.lastMessageTimeStyle.Render(lastMessageSentTime),
	)

	emptySpace := common.HorizontalGap(availableWidth)

	if focus {
		return v.style.Render(
			lipgloss.JoinVertical(
				lipgloss.Left,
				lipgloss.JoinHorizontal(
					lipgloss.Top,
					v.selectedTitleStyle.Render(chat.Username),
					emptySpace,
					v.lastMessageTimeStyle.Render(lastMessageSentTime),
				),
				v.selectedDescriptionStyle.Render(lastMessageContent),
			),
		)
	}

	return v.style.Render(
		lipgloss.JoinVertical(
			lipgloss.Left,
			lipgloss.JoinHorizontal(
				lipgloss.Top,
				v.titleStyle.Render(chat.Username),
				emptySpace,
				v.lastMessageTimeStyle.Render(lastMessageSentTime),
			),
			v.descriptionStyle.Render(lastMessageContent),
		),
	)
}
