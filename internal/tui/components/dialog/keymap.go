package dialog

import "github.com/charmbracelet/bubbles/key"

type KeyMap struct {
	CursorUp    key.Binding
	CursorDown  key.Binding
	SendMessage key.Binding
	BackToChats key.Binding
	ToggleHelp  key.Binding
}

func DefaultKeyMap() KeyMap {
	return KeyMap{
		CursorUp: key.NewBinding(
			key.WithKeys("up"),
			key.WithHelp("↑", "up"),
		),
		CursorDown: key.NewBinding(
			key.WithKeys("down"),
			key.WithHelp("↓", "down"),
		),
		SendMessage: key.NewBinding(
			key.WithKeys("enter"),
			key.WithHelp("enter", "send message"),
		),
		BackToChats: key.NewBinding(
			key.WithKeys("esc"),
			key.WithHelp("esc", "back to chats"),
		),
		ToggleHelp: key.NewBinding(
			key.WithKeys("?"),
			key.WithHelp("?", "show/close help"),
		),
	}
}

func (km KeyMap) ShortHelp() []key.Binding {
	return []key.Binding{
		km.CursorUp,
		km.CursorDown,
		km.SendMessage,
		km.BackToChats,
		km.ToggleHelp,
	}
}

func (km KeyMap) FullHelp() [][]key.Binding {
	return [][]key.Binding{
		{
			km.CursorUp,
			km.CursorDown,
			km.SendMessage,
			km.BackToChats,
			km.ToggleHelp,
		},
	}
}
