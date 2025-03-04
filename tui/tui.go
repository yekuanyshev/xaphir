package tui

import (
	"log"

	tea "github.com/charmbracelet/bubbletea"
	"github.com/yekuanyshev/xaphir/tui/components"
	"github.com/yekuanyshev/xaphir/tui/components/chatlist"
	"github.com/yekuanyshev/xaphir/tui/stubs"
)

func Run() {
	stubs, err := stubs.Load()
	if err != nil {
		log.Fatal(err)
	}

	chatList := chatlist.NewComponent(stubs.Chats)
	base := components.NewBase(chatList)

	p := tea.NewProgram(base, tea.WithAltScreen())
	_, err = p.Run()
	if err != nil {
		log.Fatal(err)
	}

}
