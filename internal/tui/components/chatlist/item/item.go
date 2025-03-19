package item

import (
	"github.com/charmbracelet/lipgloss"
	"github.com/charmbracelet/x/ansi"
	"github.com/yekuanyshev/xaphir/internal/tui/components/common"
	"github.com/yekuanyshev/xaphir/internal/tui/components/models"
)

type Item struct {
	models.Chat
	focus bool

	style                    lipgloss.Style
	titleStyle               lipgloss.Style
	descriptionStyle         lipgloss.Style
	selectedTitleStyle       lipgloss.Style
	selectedDescriptionStyle lipgloss.Style
	lastMessageTimeStyle     lipgloss.Style
}

func NewItem(chat models.Chat) Item {
	return Item{
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

		lastMessageTimeStyle: lipgloss.NewStyle().Faint(true),
	}
}

func (i Item) View(width int) string {
	var lastMessageContent string
	var lastMessageSentTime string

	if i.LastMessage != nil {
		lastMessageContent = ansi.Truncate(i.LastMessage.Content, width, "...")
		lastMessageSentTime = i.LastMessage.FormatSentTime()
	}

	availableWidth := common.CalculateAvailableWidth(
		width,
		i.selectedTitleStyle.Render(i.Username),
		i.lastMessageTimeStyle.Render(lastMessageSentTime),
	)

	emptySpace := lipgloss.NewStyle().Width(availableWidth).Render("")

	if i.focus {
		return i.style.Render(
			lipgloss.JoinVertical(
				lipgloss.Left,
				lipgloss.JoinHorizontal(
					lipgloss.Top,
					i.selectedTitleStyle.Render(i.Username),
					emptySpace,
					i.lastMessageTimeStyle.Render(lastMessageSentTime),
				),
				i.selectedDescriptionStyle.Render(lastMessageContent),
			),
		)
	}

	return i.style.Render(
		lipgloss.JoinVertical(
			lipgloss.Left,
			lipgloss.JoinHorizontal(
				lipgloss.Top,
				i.titleStyle.Render(i.Username),
				emptySpace,
				i.lastMessageTimeStyle.Render(lastMessageSentTime),
			),
			i.descriptionStyle.Render(lastMessageContent),
		),
	)
}

func (i *Item) Focus() { i.focus = true }
func (i *Item) Blur()  { i.focus = false }
