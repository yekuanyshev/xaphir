package help

import (
	"github.com/charmbracelet/bubbles/help"
	"github.com/charmbracelet/lipgloss"
	"github.com/yekuanyshev/xaphir/internal/tui/components/base"
)

type Component struct {
	*base.Component
	keyMap help.KeyMap
	help   help.Model
}

func New(keyMap help.KeyMap) *Component {
	help := help.New()
	help.ShowAll = true

	style := lipgloss.NewStyle().
		PaddingLeft(1).PaddingRight(1).
		BorderStyle(lipgloss.RoundedBorder()).
		BorderForeground(lipgloss.Color("36"))

	return &Component{
		Component: base.NewComponent(
			base.WithStyle(style),
		),
		keyMap: keyMap,
		help:   help,
	}
}
