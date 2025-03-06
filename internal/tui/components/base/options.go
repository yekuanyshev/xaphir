package base

import "github.com/charmbracelet/lipgloss"

type option func(c *Component)

func WithWidth(width int) option {
	return func(c *Component) {
		c.SetWidth(width)
	}
}

func WithHeight(height int) option {
	return func(c *Component) {
		c.SetHeight(height)
	}
}

func WithFocus(focus bool) option {
	return func(c *Component) {
		c.focus = true
	}
}

func WithStyle(style lipgloss.Style) option {
	return func(c *Component) {
		c.SetStyle(style)
	}
}
