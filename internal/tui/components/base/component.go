package base

type Component struct {
	width  int
	height int
	focus  bool
}

func NewComponent() *Component {
	return &Component{}
}

func (c *Component) SetWidth(width int) {
	c.width = width
}

func (c *Component) SetHeight(height int) {
	c.height = height
}

func (c *Component) Width() int {
	return c.width
}

func (c *Component) Height() int {
	return c.height
}

func (c *Component) Focus() {
	c.focus = true
}

func (c *Component) Blur() {
	c.focus = false
}

func (c *Component) Focused() bool {
	return c.focus
}
