package chatlist

import (
	"github.com/charmbracelet/lipgloss"
	"github.com/yekuanyshev/xaphir/pkg/paginator"
)

type Paginator[T any] struct {
	*paginator.ItemPaginator[T]
	style lipgloss.Style
}

func NewPaginator[T any](items []T) *Paginator[T] {
	return &Paginator[T]{
		ItemPaginator: paginator.NewItemPaginator(items, len(items)),
		style: lipgloss.NewStyle().
			AlignHorizontal(lipgloss.Center),
	}
}

func (p *Paginator[T]) View() string {
	return p.style.Render(p.String())
}

func (p *Paginator[T]) SetWidth(w int) {
	p.style = p.style.Width(w)
}

func (p *Paginator[T]) String() string {
	activeDot := "○"
	inactiveDot := "•"

	view := ""
	for page := range p.TotalPages() {
		if page == p.CurrentPage() {
			view += activeDot
		} else {
			view += inactiveDot
		}
	}

	return view
}
