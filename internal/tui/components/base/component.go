package base

import "github.com/charmbracelet/lipgloss"

type Component struct {
	width  int
	height int
	focus  bool
	style  lipgloss.Style
}

func NewComponent(opts ...option) *Component {
	c := &Component{}

	for _, opt := range opts {
		opt(c)
	}

	return c
}

func (c *Component) SetWidth(width int) {
	c.width = width
	c.style = c.style.Width(width)
}

func (c *Component) SetHeight(height int) {
	c.height = height
	c.style = c.style.Height(height)
}

func (c *Component) Width() int                    { return c.width }
func (c *Component) Height() int                   { return c.height }
func (c *Component) InnerWidth() int               { return c.width - c.style.GetHorizontalFrameSize() }
func (c *Component) InnerHeight() int              { return c.height - c.style.GetVerticalFrameSize() }
func (c *Component) Style() lipgloss.Style         { return c.style }
func (c *Component) SetStyle(style lipgloss.Style) { c.style = style }
func (c *Component) Render(strs ...string) string  { return c.style.Render(strs...) }
func (c *Component) Focus()                        { c.focus = true }
func (c *Component) Blur()                         { c.focus = false }
func (c *Component) Focused() bool                 { return c.focus }
