package chatlist

import (
	"strings"

	"github.com/charmbracelet/bubbles/textinput"
	tea "github.com/charmbracelet/bubbletea"
	"github.com/yekuanyshev/xaphir/pkg/utils"
)

type filter struct {
	enabled bool
	input   textinput.Model
}

func newFilter() *filter {
	input := textinput.New()
	input.Placeholder = "Search..."

	return &filter{
		enabled: false,
		input:   input,
	}
}

func (f *filter) setWidth(width int) {
	f.input.Width = width
	f.input.CharLimit = 32
}

func (f *filter) update(msg tea.Msg) (*filter, tea.Cmd) {
	if !f.enabled {
		return f, nil
	}

	var cmd tea.Cmd

	f.input, cmd = f.input.Update(msg)

	return f, cmd
}

func (f *filter) enable() {
	f.enabled = true
	f.input.Focus()
}

func (f *filter) disable() {
	f.enabled = false
	f.input.Blur()
	f.input.SetValue("")
}

func (f *filter) filterItems(items []Item) []Item {
	inputValue := f.input.Value()
	inputValue = strings.TrimSpace(inputValue)
	inputValue = strings.ToLower(inputValue)

	if inputValue == "" {
		return nil
	}

	return utils.SliceFilter(items, f.filterBy(inputValue))
}

func (f *filter) filterBy(search string) func(Item) bool {
	return func(item Item) bool {
		username := strings.ToLower(item.Username)
		return strings.Contains(username, search)
	}
}
