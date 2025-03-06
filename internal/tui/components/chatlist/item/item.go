package item

import (
	"github.com/charmbracelet/lipgloss"
	"github.com/charmbracelet/x/ansi"
	"github.com/yekuanyshev/xaphir/internal/tui/components/dialog/item"
)

type Chat struct {
	Username    string
	LastMessage string
	Messages    []item.Message
}

type Item struct {
	Chat
	selected bool

	style                    lipgloss.Style
	titleStyle               lipgloss.Style
	descriptionStyle         lipgloss.Style
	selectedTitleStyle       lipgloss.Style
	selectedDescriptionStyle lipgloss.Style
}

func NewItem(chat Chat) Item {
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
	}
}

func (i Item) View(width int) string {
	i.LastMessage = ansi.Truncate(i.LastMessage, width, "...")

	if i.selected {
		return i.style.Render(
			lipgloss.JoinVertical(
				lipgloss.Left,
				i.selectedTitleStyle.Render(i.Username),
				i.selectedDescriptionStyle.Render(i.LastMessage),
			),
		)
	}

	return i.style.Render(
		lipgloss.JoinVertical(
			lipgloss.Left,
			i.titleStyle.Render(i.Username),
			i.descriptionStyle.Render(i.LastMessage),
		),
	)
}

func (i *Item) SetSelected(selected bool) {
	i.selected = selected
}
