package paginator

type ItemPaginator[T any] struct {
	*CursorPaginator
	items []T
}

func NewItemPaginator[T any](items []T, limit int) *ItemPaginator[T] {
	return &ItemPaginator[T]{
		CursorPaginator: NewCursorPaginator(len(items), limit),
		items:           items,
	}
}

func (p *ItemPaginator[T]) ItemsOnCurrentPage() []T {
	start, end := p.GetIndexBounds()
	return p.items[start:end]
}

func (p *ItemPaginator[T]) CurrentItem() T {
	return p.items[p.CurrentIndex()]
}

func (p *ItemPaginator[T]) CurrentIndex() int {
	return p.page*p.limit + p.cursor
}
