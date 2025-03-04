package components

import (
	tea "github.com/charmbracelet/bubbletea"
	"github.com/yekuanyshev/xaphir/tui/components/chatlist"
)

type Base struct {
	chatList *chatlist.Component
}

func NewBase(chatList *chatlist.Component) *Base {
	return &Base{
		chatList: chatList,
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
		chatListHeight := int(float64(msg.Height) * 0.9)

		b.chatList.SetWidth(chatListWidth)
		b.chatList.SetHeight(chatListHeight)
	}

	model, chatListCmd := b.chatList.Update(msg)
	b.chatList = model.(*chatlist.Component)

	return b, tea.Sequence(
		chatListCmd,
	)
}

func (b *Base) View() string {
	return b.chatList.View()
}
