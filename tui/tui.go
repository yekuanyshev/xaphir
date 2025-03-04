package tui

import (
	"log"

	tea "github.com/charmbracelet/bubbletea"
	"github.com/yekuanyshev/xaphir/tui/components"
	"github.com/yekuanyshev/xaphir/tui/stubs"
)

func Run() {
	stubs, err := stubs.Load()
	if err != nil {
		log.Fatal(err)
	}

	clc := components.NewChatList(50, 50, stubs.Chats)
	p := tea.NewProgram(clc, tea.WithAltScreen())
	_, err = p.Run()
	if err != nil {
		log.Fatal(err)
	}

}
