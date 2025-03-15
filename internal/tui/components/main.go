package components

import (
	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
	"github.com/yekuanyshev/xaphir/internal/tui/components/chatlist"
	"github.com/yekuanyshev/xaphir/internal/tui/components/dialog"
	"github.com/yekuanyshev/xaphir/internal/tui/components/events"
)

type Main struct {
	chatList *chatlist.Component
	dialog   *dialog.Component
}

func NewMain(
	chatList *chatlist.Component,
	dialog *dialog.Component,
) *Main {
	return &Main{
		chatList: chatList,
		dialog:   dialog,
	}
}

func (m *Main) Init() tea.Cmd {
	m.chatList.Focus()
	m.dialog.Blur()
	return nil
}

func (m *Main) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	if events.IsChatListFocusCMD(msg) {
		m.chatList.Focus()
		return m, nil
	}

	if msg, ok := events.IsDialogFocusCMD(msg); ok {
		m.dialog.SetTitle(msg.Title)
		m.dialog.SetSliderMessages(msg.Items)
		m.dialog.Focus()
		return m, nil
	}

	switch msg := msg.(type) {
	case tea.KeyMsg:
		switch msg.String() {
		case "ctrl+c":
			return m, tea.Quit
		}
	case tea.WindowSizeMsg:
		chatListWidth := int(float64(msg.Width) * 0.2)
		dialogWidth := int(float64(msg.Width)*0.8 - 3)
		height := int(float64(msg.Height) - 2)

		m.chatList.SetWidth(chatListWidth)
		m.chatList.SetHeight(height)

		m.dialog.SetWidth(dialogWidth)
		m.dialog.SetHeight(height)
	}

	model, chatListCmd := m.chatList.Update(msg)
	m.chatList = model.(*chatlist.Component)

	model, dialogCmd := m.dialog.Update(msg)
	m.dialog = model.(*dialog.Component)

	return m, tea.Batch(
		chatListCmd,
		dialogCmd,
	)
}

func (m *Main) View() string {
	return lipgloss.JoinHorizontal(
		lipgloss.Top,
		m.chatList.View(),
		m.dialog.View(),
	)
}
