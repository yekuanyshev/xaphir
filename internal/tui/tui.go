package tui

import (
	"context"
	"log"
	"os"
	"os/signal"
	"syscall"

	tea "github.com/charmbracelet/bubbletea"
	"github.com/yekuanyshev/xaphir/internal/stubs"
	"github.com/yekuanyshev/xaphir/internal/tui/components"
	"github.com/yekuanyshev/xaphir/internal/tui/components/chatlist"
	"github.com/yekuanyshev/xaphir/internal/tui/components/dialog"
)

func Run(ctx context.Context) {
	ctx, cancel := signal.NotifyContext(ctx, os.Interrupt, syscall.SIGTERM)
	defer cancel()

	stubs, err := stubs.Load()
	if err != nil {
		log.Fatal(err)
	}

	dialog := dialog.NewComponent()
	chatList := chatlist.NewComponent(stubs.Chats)
	main := components.NewMain(chatList, dialog)

	p := tea.NewProgram(
		main,
		tea.WithContext(ctx),
		tea.WithAltScreen(),
	)
	_, err = p.Run()
	if err != nil {
		log.Fatal(err)
	}
}
