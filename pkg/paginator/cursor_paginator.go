package paginator

type CursorPaginator struct {
	*Paginator
	cursor int // index of the current item on page
}

func NewCursorPaginator(total, limit int) *CursorPaginator {
	return &CursorPaginator{
		Paginator: NewPaginator(total, limit),
		cursor:    0,
	}
}

func (cp *CursorPaginator) Increment() {
	cp.cursor++
	if cp.cursor >= cp.limit {
		if cp.OnLastPage() {
			cp.cursor--
			return
		}

		cp.NextPage()
		cp.cursor = 0
		return
	}

	n := cp.NumOfItemsOnPage()
	if cp.OnLastPage() && cp.cursor >= n-1 {
		cp.cursor = n - 1
	}
}

func (cp *CursorPaginator) Decrement() {
	cp.cursor--
	if cp.cursor < 0 {
		if cp.OnFirstPage() {
			cp.cursor = 0
			return
		}

		cp.PrevPage()
		cp.cursor = cp.NumOfItemsOnPage() - 1
	}
}

func (cp *CursorPaginator) SkipToNextPage() {
	if !cp.OnLastPage() {
		cp.cursor = 0
	}
	cp.NextPage()
}

func (cp *CursorPaginator) SkipToPrevPage() {
	if !cp.OnFirstPage() {
		cp.cursor = 0
	}
	cp.PrevPage()
}

func (cp *CursorPaginator) Cursor() int {
	return cp.cursor
}

func (cp *CursorPaginator) SetLimit(limit int) {
	cp.Paginator.SetLimit(limit)
	cp.cursor = 0
}

func (cp *CursorPaginator) SetTotal(total int) {
	cp.Paginator.SetTotal(total)
	cp.cursor = 0
}
