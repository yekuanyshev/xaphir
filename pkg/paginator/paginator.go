package paginator

import "math"

type Paginator struct {
	total      int // total number of items
	limit      int // number of items per page
	page       int // current page
	totalPages int // number of total pages
}

func NewPaginator(total, limit int) *Paginator {
	return &Paginator{
		total:      total,
		limit:      limit,
		page:       0,
		totalPages: int(math.Ceil(float64(total) / float64(limit))),
	}
}

func (p *Paginator) NextPage() {
	if !p.OnLastPage() {
		p.page++
	}
}

func (p *Paginator) PrevPage() {
	if !p.OnFirstPage() {
		p.page--
	}
}

func (p *Paginator) OnFirstPage() bool {
	return p.page == 0
}

func (p *Paginator) OnLastPage() bool {
	return p.page == p.totalPages-1
}

func (m *Paginator) GetIndexBounds() (start int, end int) {
	start = m.page * m.limit
	end = min(m.page*m.limit+m.limit, m.total)
	return start, end
}

func (m *Paginator) NumOfItemsOnPage() int {
	if m.total < 0 {
		return 0
	}

	start, end := m.GetIndexBounds()
	return end - start
}

func (p *Paginator) CurrentPage() int {
	return p.page
}

func (p *Paginator) TotalPages() int {
	return p.totalPages
}

func (p *Paginator) SetLimit(limit int) {
	p.limit = limit
	p.page = 0
	p.totalPages = int(math.Ceil(float64(p.total) / float64(limit)))
}

func (p *Paginator) SetTotal(total int) {
	p.total = total
	p.page = 0
	p.totalPages = int(math.Ceil(float64(p.total) / float64(p.limit)))
}
