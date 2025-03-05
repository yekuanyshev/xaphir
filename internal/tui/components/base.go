package components

import (
	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
	"github.com/yekuanyshev/xaphir/internal/tui/components/chatlist"
	"github.com/yekuanyshev/xaphir/internal/tui/components/dialog"
)

type Base struct {
	chatList *chatlist.Component
	dialog   *dialog.Component
}

func NewBase(
	chatList *chatlist.Component,
	dialog *dialog.Component,
) *Base {
	return &Base{
		chatList: chatList,
		dialog:   dialog,
	}
}

func (b *Base) Init() tea.Cmd {
	return nil
}

func (b *Base) Update(msg tea.Msg) (tea.Model, tea.Cmd) {

	switch msg := msg.(type) {
	case tea.KeyMsg:
		switch msg.String() {
		case "ctrl+c", "q":
			return b, tea.Quit
		}
	case tea.WindowSizeMsg:
		chatListWidth := int(float64(msg.Width) * 0.2)
		dialogWidth := int(float64(msg.Width)*0.8 - 3)
		height := int(float64(msg.Height) * 0.9)

		b.chatList.SetWidth(chatListWidth)
		b.chatList.SetHeight(height)

		b.dialog.SetWidth(dialogWidth)
		b.dialog.SetHeight(height)
	}

	model, chatListCmd := b.chatList.Update(msg)
	b.chatList = model.(*chatlist.Component)

	model, dialogCmd := b.dialog.Update(msg)
	b.dialog = model.(*dialog.Component)

	return b, tea.Sequence(
		chatListCmd,
		dialogCmd,
	)
}

func (b *Base) View() string {
	return lipgloss.JoinHorizontal(
		lipgloss.Top,
		b.chatList.View(),
		b.dialog.View(),
	)
}
