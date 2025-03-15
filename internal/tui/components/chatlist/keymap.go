package chatlist

import "github.com/charmbracelet/bubbles/key"

type KeyMap struct {
	CursorUp   key.Binding
	CursorDown key.Binding
	NextPage   key.Binding
	PrevPage   key.Binding
	GoToDialog key.Binding
	ToggleHelp key.Binding
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
		NextPage: key.NewBinding(
			key.WithKeys("right"),
			key.WithHelp("→", "next page"),
		),
		PrevPage: key.NewBinding(
			key.WithKeys("left"),
			key.WithHelp("←", "prev page"),
		),
		GoToDialog: key.NewBinding(
			key.WithKeys("enter"),
			key.WithHelp("enter", "go to dialog"),
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
		km.NextPage,
		km.PrevPage,
		km.GoToDialog,
		km.ToggleHelp,
	}
}

func (km KeyMap) FullHelp() [][]key.Binding {
	return [][]key.Binding{
		{
			km.CursorUp,
			km.CursorDown,
			km.NextPage,
			km.PrevPage,
			km.GoToDialog,
			km.ToggleHelp,
		},
	}
}
