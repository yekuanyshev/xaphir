package events

import (
	tea "github.com/charmbracelet/bubbletea"
	"github.com/yekuanyshev/xaphir/internal/tui/components/dialog/item"
)

type (
	DialogFocus struct {
		Title string
		Items []item.Message
	}
	ChatListFocus struct{}
)

func DialogFocusCMD(title string, items []item.Message) tea.Cmd {
	return func() tea.Msg {
		return DialogFocus{
			Title: title,
			Items: items,
		}
	}
}

func ChatListFocusCMD() tea.Cmd {
	return func() tea.Msg {
		return ChatListFocus{}
	}
}

func IsDialogFocusCMD(msg tea.Msg) (DialogFocus, bool) {
	dialogFocus, ok := msg.(DialogFocus)
	return dialogFocus, ok
}

func IsChatListFocusCMD(msg tea.Msg) bool {
	_, ok := msg.(ChatListFocus)
	return ok
}
