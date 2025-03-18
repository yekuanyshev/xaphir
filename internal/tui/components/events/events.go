package events

import (
	tea "github.com/charmbracelet/bubbletea"
)

type (
	DialogFocus struct {
		ChatID int64
	}
	ChatListFocus struct{}
	SendMessage   struct {
		ChatID  int64
		Content string
	}
)

func DialogFocusCMD(chatID int64) tea.Cmd {
	return func() tea.Msg {
		return DialogFocus{ChatID: chatID}
	}
}

func ChatListFocusCMD() tea.Cmd {
	return func() tea.Msg {
		return ChatListFocus{}
	}
}

func SendMessageCMD(chatID int64, content string) tea.Cmd {
	return func() tea.Msg {
		return SendMessage{
			ChatID:  chatID,
			Content: content,
		}
	}
}
