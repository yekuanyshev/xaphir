package components

import "github.com/charmbracelet/bubbles/key"

type KeyMap struct {
	Quit       key.Binding
	ToggleHelp key.Binding
}

func DefaultKeyMap() KeyMap {
	return KeyMap{
		Quit: key.NewBinding(
			key.WithKeys("ctrl+c"),
		),
		ToggleHelp: key.NewBinding(
			key.WithKeys("?"),
		),
	}
}
