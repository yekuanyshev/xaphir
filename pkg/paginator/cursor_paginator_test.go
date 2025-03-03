package paginator

import "testing"

func TestCursorPaginatorInitialization(t *testing.T) {
	cp := NewCursorPaginator(100, 10)
	if cp.total != 100 {
		t.Errorf("expected total to be 100, got %d", cp.total)
	}
	if cp.limit != 10 {
		t.Errorf("expected limit to be 10, got %d", cp.limit)
	}
	if cp.page != 0 {
		t.Errorf("expected initial page to be 0, got %d", cp.page)
	}
	if cp.cursor != 0 {
		t.Errorf("expected initial cursor to be 0, got %d", cp.cursor)
	}
}

func TestCursorPaginatorIncrement(t *testing.T) {
	cp := NewCursorPaginator(25, 10)

	for i := 0; i < 9; i++ {
		cp.Increment()
		if cp.cursor != i+1 {
			t.Errorf("expected cursor to be %d, got %d", i+1, cp.cursor)
		}
	}

	cp.Increment() // Should move to next page
	if cp.page != 1 || cp.cursor != 0 {
		t.Errorf("expected cursor to reset and page to increment, got page %d, cursor %d", cp.page, cp.cursor)
	}
}

func TestCursorPaginatorDecrement(t *testing.T) {
	cp := NewCursorPaginator(25, 10)
	cp.cursor = 5
	cp.Decrement()
	if cp.cursor != 4 {
		t.Errorf("expected cursor to be 4, got %d", cp.cursor)
	}

	cp.cursor = 0
	cp.Decrement() // Should not move to last item on prev page
	if cp.cursor != 0 || cp.page != 0 {
		t.Errorf("expected cursor to be 9 on previous page, got page %d, cursor %d", cp.page, cp.cursor)
	}
}

func TestCursorPaginatorSkipToNextPage(t *testing.T) {
	cp := NewCursorPaginator(50, 10)
	cp.cursor = 5
	cp.SkipToNextPage()
	if cp.cursor != 0 || cp.page != 1 {
		t.Errorf("expected cursor to reset and page to be 1, got page %d, cursor %d", cp.page, cp.cursor)
	}
}

func TestCursorPaginatorSkipToPrevPage(t *testing.T) {
	cp := NewCursorPaginator(50, 10)
	cp.page = 2
	cp.cursor = 7
	cp.SkipToPrevPage()
	if cp.cursor != 0 || cp.page != 1 {
		t.Errorf("expected cursor to reset and page to be 1, got page %d, cursor %d", cp.page, cp.cursor)
	}
}

func TestCursorPaginatorBoundaryConditions(t *testing.T) {
	cp := NewCursorPaginator(25, 10)
	for i := 0; i < 25; i++ {
		cp.Increment()
	}
	if cp.cursor != 4 || cp.page != 2 {
		t.Errorf("expected cursor at last item of last page, got page %d, cursor %d", cp.page, cp.cursor)
	}

	for i := 0; i < 25; i++ {
		cp.Decrement()
	}
	if cp.cursor != 0 || cp.page != 0 {
		t.Errorf("expected cursor at first item of first page, got page %d, cursor %d", cp.page, cp.cursor)
	}
}
