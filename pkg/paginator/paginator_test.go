package paginator

import "testing"

func TestPaginatorInitialization(t *testing.T) {
	p := NewPaginator(100, 10)
	if p.total != 100 {
		t.Errorf("expected total to be 100, got %d", p.total)
	}
	if p.limit != 10 {
		t.Errorf("expected limit to be 10, got %d", p.limit)
	}
	if p.page != 0 {
		t.Errorf("expected initial page to be 0, got %d", p.page)
	}
	if p.totalPages != 10 {
		t.Errorf("expected total pages to be 10, got %d", p.totalPages)
	}
}

func TestPaginatorNavigation(t *testing.T) {
	p := NewPaginator(25, 10)
	p.NextPage()
	if p.page != 1 {
		t.Errorf("expected page to be 1 after NextPage, got %d", p.page)
	}
	p.NextPage()
	if p.page != 2 {
		t.Errorf("expected page to be 2 after NextPage, got %d", p.page)
	}
	p.NextPage() // Shouldn't increment beyond last page
	if p.page != 2 {
		t.Errorf("expected page to remain 2 after NextPage at last page, got %d", p.page)
	}
	p.PrevPage()
	if p.page != 1 {
		t.Errorf("expected page to be 1 after PrevPage, got %d", p.page)
	}
	p.PrevPage()
	if p.page != 0 {
		t.Errorf("expected page to be 0 after PrevPage, got %d", p.page)
	}
	p.PrevPage() // Shouldn't decrement below first page
	if p.page != 0 {
		t.Errorf("expected page to remain 0 after PrevPage at first page, got %d", p.page)
	}
}

func TestOnFirstPage(t *testing.T) {
	p := NewPaginator(50, 10)
	if !p.OnFirstPage() {
		t.Errorf("expected OnFirstPage to be true on first page")
	}
	p.NextPage()
	if p.OnFirstPage() {
		t.Errorf("expected OnFirstPage to be false after NextPage")
	}
}

func TestOnLastPage(t *testing.T) {
	p := NewPaginator(25, 10)
	p.NextPage()
	p.NextPage()
	if !p.OnLastPage() {
		t.Errorf("expected OnLastPage to be true on last page")
	}
	p.PrevPage()
	if p.OnLastPage() {
		t.Errorf("expected OnLastPage to be false after PrevPage")
	}
}

func TestGetIndexBounds(t *testing.T) {
	tests := []struct {
		total, limit, page         int
		expectedStart, expectedEnd int
	}{
		{100, 10, 0, 0, 10},
		{100, 10, 1, 10, 20},
		{35, 10, 3, 30, 35},
	}

	for _, tt := range tests {
		p := NewPaginator(tt.total, tt.limit)
		p.page = tt.page
		start, end := p.GetIndexBounds()
		if start != tt.expectedStart || end != tt.expectedEnd {
			t.Errorf("For total=%d, limit=%d, page=%d, expected bounds (%d, %d), got (%d, %d)",
				tt.total, tt.limit, tt.page, tt.expectedStart, tt.expectedEnd, start, end)
		}
	}
}

func TestNumOfItemsOnPage(t *testing.T) {
	p := NewPaginator(35, 10)
	if p.NumOfItemsOnPage() != 10 {
		t.Errorf("expected 10 items on first page, got %d", p.NumOfItemsOnPage())
	}
	p.NextPage()
	if p.NumOfItemsOnPage() != 10 {
		t.Errorf("expected 10 items on second page, got %d", p.NumOfItemsOnPage())
	}
	p.NextPage()
	if p.NumOfItemsOnPage() != 10 {
		t.Errorf("expected 10 items on third page, got %d", p.NumOfItemsOnPage())
	}
	p.NextPage()
	if p.NumOfItemsOnPage() != 5 {
		t.Errorf("expected 5 items on last page, got %d", p.NumOfItemsOnPage())
	}
}

func TestCurrentPage(t *testing.T) {
	p := NewPaginator(100, 10)
	if p.CurrentPage() != 0 {
		t.Errorf("expected current page to be 0, got %d", p.CurrentPage())
	}
	p.NextPage()
	if p.CurrentPage() != 1 {
		t.Errorf("expected current page to be 1, got %d", p.CurrentPage())
	}
}

func TestTotalPages(t *testing.T) {
	p := NewPaginator(95, 10)
	if p.TotalPages() != 10 {
		t.Errorf("expected total pages to be 10, got %d", p.TotalPages())
	}
	p = NewPaginator(100, 10)
	if p.TotalPages() != 10 {
		t.Errorf("expected total pages to be 10, got %d", p.TotalPages())
	}
	p = NewPaginator(101, 10)
	if p.TotalPages() != 11 {
		t.Errorf("expected total pages to be 11, got %d", p.TotalPages())
	}
}
