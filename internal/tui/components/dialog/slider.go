package dialog

import (
	"github.com/charmbracelet/lipgloss"
	"github.com/yekuanyshev/xaphir/internal/tui/components/dialog/item"
	"github.com/yekuanyshev/xaphir/pkg/utils"
)

type Slider struct {
	width  int
	height int

	items []item.Item
	start int
	end   int
}

func NewSlider() *Slider {
	return &Slider{}
}

func (s *Slider) SetWidth(width int) {
	s.width = width
	for i := range s.items {
		s.items[i].SetWidth(width)
	}
}
func (s *Slider) SetHeight(height int) { s.height = height }

func (s *Slider) SetMessages(messages []item.Message) {
	converter := func(message item.Message) item.Item {
		return item.NewItem(message, s.width)
	}
	s.items = utils.SliceMap(messages, converter)
	s.end = len(s.items)
	s.start = s.calculateStart(max(s.end-1, 0))
}

func (s *Slider) AppendMessage(message item.Message) {
	item := item.NewItem(message, s.width)
	s.items = append(s.items, item)
	s.end = len(s.items)
	s.start = s.calculateStart(max(s.end-1, 0))
}

func (s *Slider) Increment() {
	if s.end < len(s.items) {
		s.end = min(s.end+1, len(s.items))
		s.start = s.calculateStart(s.end - 1)
	}
}

func (s *Slider) Decrement() {
	s.end = max(s.end-1, s.end-s.start)
	s.start = s.calculateStart(s.end - 1)
}

func (s *Slider) GetItems() []item.Item {
	return s.items[s.start:s.end]
}

func (s *Slider) calculateStart(end int) int {
	if end <= 0 {
		return 0
	}

	availHeight := s.height
	h := 0
	i := end

	for i >= 0 {
		itemViewHeight := lipgloss.Height(s.items[i].View())

		if h+itemViewHeight >= availHeight {
			return i + 1
		}
		h += itemViewHeight
		i--
	}

	return max(i, 0)
}
