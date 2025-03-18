package components

import (
	"log"

	"github.com/charmbracelet/bubbles/key"
	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
	"github.com/yekuanyshev/xaphir/internal/service"
	"github.com/yekuanyshev/xaphir/internal/tui/components/chatlist"
	"github.com/yekuanyshev/xaphir/internal/tui/components/common"
	"github.com/yekuanyshev/xaphir/internal/tui/components/dialog"
	"github.com/yekuanyshev/xaphir/internal/tui/components/events"
	"github.com/yekuanyshev/xaphir/pkg/utils"
)

type Main struct {
	width  int
	height int

	srv *service.Service

	chatList         *chatlist.Component
	dialog           *dialog.Component
	showChatListHelp bool
	showDialogHelp   bool

	keyMap KeyMap
}

func NewMain(
	srv *service.Service,
	chatList *chatlist.Component,
	dialog *dialog.Component,
) *Main {
	return &Main{
		srv:      srv,
		chatList: chatList,
		dialog:   dialog,
		keyMap:   DefaultKeyMap(),
	}
}

func (m *Main) Init() tea.Cmd {
	chats, err := m.srv.ListChats()
	if err != nil {
		log.Fatal(err)
	}

	items := utils.SliceMap(chats, service.Chat.ToComponentModel)

	m.chatList.SetItems(items)
	m.chatList.Focus()
	m.dialog.Blur()
	return nil
}

func (m *Main) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	switch msg := msg.(type) {
	case tea.KeyMsg:
		switch {
		case key.Matches(msg, m.keyMap.Quit):
			return m, tea.Quit
		case key.Matches(msg, m.keyMap.ToggleHelp):
			if !m.dialog.IsTypingMessage() {
				m.toggleChatListHelp()
				m.toggleDialogHelp()
				return m, nil
			}
		}
	case tea.WindowSizeMsg:
		m.width, m.height = msg.Width, msg.Height

		chatListWidth := int(float64(msg.Width) * 0.2)
		dialogWidth := int(float64(msg.Width)*0.8 - 3)
		height := int(float64(msg.Height) - 2)

		m.chatList.SetWidth(chatListWidth)
		m.chatList.SetHeight(height)

		m.dialog.SetWidth(dialogWidth)
		m.dialog.SetHeight(height)
	case events.ChatListFocus:
		m.chatList.Focus()
		return m, nil
	case events.DialogFocus:
		m.handleDialogFocus(msg)
		return m, nil
	case events.SendMessage:
		m.handleSendMessage(msg)
		return m, nil
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
	view := lipgloss.JoinHorizontal(
		lipgloss.Top,
		m.chatList.View(),
		m.dialog.View(),
	)

	showHelp := false
	helpView := ""

	if m.showChatListHelp {
		showHelp = true
		helpView = m.chatList.HelpView()
	}

	if m.showDialogHelp {
		showHelp = true
		helpView = m.dialog.HelpView()
	}

	if showHelp {
		centerX := m.width / 2
		centerY := m.height / 2
		helpViewWidth, helpViewHeight := lipgloss.Size(helpView)
		x := centerX - helpViewWidth/2
		y := centerY - helpViewHeight/2

		return common.PlaceOverlay(
			x, y, helpView, view,
		)
	}

	return view
}

func (m *Main) handleDialogFocus(msg events.DialogFocus) {
	chat, err := m.srv.GetChat(msg.ChatID)
	if err != nil {
		log.Fatal(err)
	}

	items := utils.SliceMap(chat.Messages, service.ChatMessage.ToComponentModel)

	m.dialog.SetChatID(msg.ChatID)
	m.dialog.SetTitle(chat.Member.Username)
	m.dialog.SetSliderMessages(items)
	m.dialog.Focus()
}

func (m *Main) handleSendMessage(msg events.SendMessage) {
	err := m.srv.SendMessage(msg.ChatID, msg.Content)
	if err != nil {
		log.Fatal(err)
	}

	chat, err := m.srv.GetChat(msg.ChatID)
	if err != nil {
		log.Fatal(err)
	}

	items := utils.SliceMap(chat.Messages, service.ChatMessage.ToComponentModel)

	m.dialog.SetSliderMessages(items)
}

func (m *Main) toggleChatListHelp() {
	if m.showChatListHelp {
		m.showChatListHelp = false
		m.chatList.Focus()
		return
	}

	if m.chatList.Focused() {
		m.showChatListHelp = true
		m.chatList.Blur()
	}
}

func (m *Main) toggleDialogHelp() {
	if m.showDialogHelp {
		m.showDialogHelp = false
		m.dialog.Focus()
		return
	}

	if m.dialog.Focused() {
		m.showDialogHelp = true
		m.dialog.Blur()
	}
}
