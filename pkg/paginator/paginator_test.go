package paginator

import (
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestPaginatorInitialization(t *testing.T) {
	p := NewPaginator(100, 10)
	assert.Equal(t, 100, p.total)
	assert.Equal(t, 10, p.limit)
	assert.Equal(t, 0, p.page)
	assert.Equal(t, 10, p.totalPages)
}

func TestPaginatorNavigation(t *testing.T) {
	p := NewPaginator(25, 10)
	p.NextPage()
	assert.Equal(t, 1, p.page)

	p.NextPage()
	assert.Equal(t, 2, p.page)

	p.NextPage()
	assert.Equal(t, 2, p.page) // Shouldn't increment beyond last page

	p.PrevPage()
	assert.Equal(t, 1, p.page)

	p.PrevPage()
	assert.Equal(t, 0, p.page)

	p.PrevPage()
	assert.Equal(t, 0, p.page) // Shouldn't decrement below first page
}

func TestOnFirstPage(t *testing.T) {
	p := NewPaginator(50, 10)
	assert.True(t, p.OnFirstPage())

	p.NextPage()
	assert.False(t, p.OnFirstPage())
}

func TestOnLastPage(t *testing.T) {
	p := NewPaginator(25, 10)

	p.NextPage()
	p.NextPage()
	assert.True(t, p.OnLastPage())

	p.PrevPage()
	assert.False(t, p.OnLastPage())
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
		assert.Equal(t, tt.expectedStart, start)
		assert.Equal(t, tt.expectedEnd, end)
	}
}

func TestNumOfItemsOnPage(t *testing.T) {
	p := NewPaginator(35, 10)

	assert.Equal(t, 10, p.NumOfItemsOnPage())

	p.NextPage()
	assert.Equal(t, 10, p.NumOfItemsOnPage())

	p.NextPage()
	assert.Equal(t, 10, p.NumOfItemsOnPage())

	p.NextPage()
	assert.Equal(t, 5, p.NumOfItemsOnPage())
}

func TestCurrentPage(t *testing.T) {
	p := NewPaginator(100, 10)
	assert.Equal(t, 0, p.CurrentPage())

	p.NextPage()
	assert.Equal(t, 1, p.CurrentPage())
}

func TestTotalPages(t *testing.T) {
	p := NewPaginator(95, 10)
	assert.Equal(t, 10, p.TotalPages())

	p = NewPaginator(100, 10)
	assert.Equal(t, 10, p.TotalPages())

	p = NewPaginator(101, 10)
	assert.Equal(t, 11, p.TotalPages())
}

func TestSetLimit(t *testing.T) {
	p := NewPaginator(95, 10)
	assert.Equal(t, 95, p.total)
	assert.Equal(t, 10, p.limit)
	assert.Equal(t, 0, p.page)
	assert.Equal(t, 10, p.totalPages)

	p.SetLimit(20)
	assert.Equal(t, 95, p.total)
	assert.Equal(t, 20, p.limit)
	assert.Equal(t, 0, p.page)
	assert.Equal(t, 5, p.totalPages)

	p.SetLimit(0)
	assert.Equal(t, 95, p.total)
	assert.Equal(t, 0, p.limit)
	assert.Equal(t, 0, p.page)
	assert.Equal(t, 0, p.totalPages)
}

func TestSetTotal(t *testing.T) {
	p := NewPaginator(95, 10)
	assert.Equal(t, 95, p.total)
	assert.Equal(t, 10, p.limit)
	assert.Equal(t, 0, p.page)
	assert.Equal(t, 10, p.totalPages)

	p.SetTotal(50)
	assert.Equal(t, 50, p.total)
	assert.Equal(t, 10, p.limit)
	assert.Equal(t, 0, p.page)
	assert.Equal(t, 5, p.totalPages)
}

func TestString(t *testing.T) {
	p := NewPaginator(100, 10)
	dots := make([]string, 0, p.TotalPages())
	for range p.TotalPages() {
		dots = append(dots, inactiveDot)
	}

	for page := range p.TotalPages() {
		dots[page] = activeDot

		expected := strings.Join(dots, "")

		assert.Equal(t, expected, p.String())
		p.NextPage()

		dots[page] = inactiveDot
	}
}
