package paginator

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestCursorPaginatorInitialization(t *testing.T) {
	cp := NewCursorPaginator(100, 10)
	assert.Equal(t, 100, cp.total)
	assert.Equal(t, 10, cp.limit)
	assert.Equal(t, 0, cp.page)
	assert.Equal(t, 0, cp.cursor)
}

func TestCursorPaginatorIncrement(t *testing.T) {
	cp := NewCursorPaginator(25, 10)

	for i := 0; i < 9; i++ {
		cp.Increment()
		assert.Equal(t, i+1, cp.cursor)
	}

	cp.Increment() // Should move to next page
	assert.Equal(t, 1, cp.page)
	assert.Equal(t, 0, cp.cursor)
}

func TestCursorPaginatorDecrement(t *testing.T) {
	cp := NewCursorPaginator(25, 10)
	cp.cursor = 5
	cp.Decrement()
	assert.Equal(t, 4, cp.cursor)

	cp.cursor = 0
	cp.Decrement() // Should not move to last item on prev page
	assert.Equal(t, 0, cp.cursor)
	assert.Equal(t, 0, cp.page)
}

func TestCursorPaginatorSkipToNextPage(t *testing.T) {
	cp := NewCursorPaginator(50, 10)
	cp.cursor = 5
	cp.SkipToNextPage()
	assert.Equal(t, 0, cp.cursor)
	assert.Equal(t, 1, cp.page)
}

func TestCursorPaginatorSkipToPrevPage(t *testing.T) {
	cp := NewCursorPaginator(50, 10)
	cp.page = 2
	cp.cursor = 7
	cp.SkipToPrevPage()
	assert.Equal(t, 0, cp.cursor)
	assert.Equal(t, 1, cp.page)
}

func TestCursorPaginatorBoundaryConditions(t *testing.T) {
	cp := NewCursorPaginator(25, 10)
	for i := 0; i < 25; i++ {
		cp.Increment()
	}

	assert.Equal(t, 4, cp.cursor)
	assert.Equal(t, 2, cp.page)

	for i := 0; i < 25; i++ {
		cp.Decrement()
	}

	assert.Equal(t, 0, cp.cursor)
	assert.Equal(t, 0, cp.page)
}

func TestCursorPaginatorSetLimit(t *testing.T) {
	cp := NewCursorPaginator(100, 10)
	assert.Equal(t, 100, cp.total)
	assert.Equal(t, 10, cp.limit)
	assert.Equal(t, 0, cp.page)
	assert.Equal(t, 10, cp.totalPages)
	assert.Equal(t, 0, cp.cursor)

	cp.SetLimit(20)
	assert.Equal(t, 100, cp.total)
	assert.Equal(t, 20, cp.limit)
	assert.Equal(t, 0, cp.page)
	assert.Equal(t, 5, cp.totalPages)
	assert.Equal(t, 0, cp.cursor)
}

func TestCursorPaginatorSetTotal(t *testing.T) {
	cp := NewCursorPaginator(100, 10)
	assert.Equal(t, 100, cp.total)
	assert.Equal(t, 10, cp.limit)
	assert.Equal(t, 0, cp.page)
	assert.Equal(t, 10, cp.totalPages)
	assert.Equal(t, 0, cp.cursor)

	cp.SetTotal(50)
	assert.Equal(t, 50, cp.total)
	assert.Equal(t, 10, cp.limit)
	assert.Equal(t, 0, cp.page)
	assert.Equal(t, 5, cp.totalPages)
	assert.Equal(t, 0, cp.cursor)
}
