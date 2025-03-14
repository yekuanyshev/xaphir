package dialog

import (
	"time"

	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
	"github.com/yekuanyshev/xaphir/internal/tui/components/common"
	"github.com/yekuanyshev/xaphir/internal/tui/components/dialog/item"
	"github.com/yekuanyshev/xaphir/internal/tui/components/events"
	"github.com/yekuanyshev/xaphir/pkg/utils"
)

func (c *Component) Init() tea.Cmd {
	return nil
}

func (c *Component) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	if !c.Focused() {
		return c, nil
	}

	switch msg := msg.(type) {
	case tea.KeyMsg:
		switch msg.String() {
		case "esc":
			c.Blur()
			return c, events.ChatListFocusCMD()
		case "enter":
			inputValue := c.input.Value()
			if inputValue == "" {
				return c, nil
			}

			message := item.Message{
				Content:  inputValue,
				SendTime: time.Now(),
				IsFromMe: true,
			}
			c.slider.AppendMessage(message)
			c.input.SetValue("")
		case "down":
			c.slider.Increment()
		case "up":
			c.slider.Decrement()
		}
	}

	var inputCMD tea.Cmd
	c.input, inputCMD = c.input.Update(msg)

	return c, tea.Batch(
		inputCMD,
	)
}

func (c *Component) View() string {
	if !c.Focused() {
		return c.blurredView()
	}

	titleView := c.titleStyle.Render(c.title)
	itemsView := c.itemsView()
	inputView := c.inputStyle.Render(c.input.View())
	availableHeight := common.CalculateAvailableHeight(
		c.InnerHeight(), titleView, itemsView, inputView,
	)
	emptySpace := common.FillWithEmptySpace(availableHeight)

	sections := []string{
		titleView,
		emptySpace,
		itemsView,
		inputView,
	}

	return c.Render(
		lipgloss.JoinVertical(
			lipgloss.Left,
			sections...,
		),
	)
}

func (c *Component) blurredView() string {
	return c.Render(
		lipgloss.Place(
			c.InnerWidth(), c.InnerHeight(),
			lipgloss.Center, lipgloss.Center,
			c.blurredTitleStyle.Render(c.blurredTitle),
		),
	)
}

func (c *Component) itemsView() string {
	itemViews := utils.SliceMap(
		c.slider.GetItems(),
		func(item item.Item) string {
			return item.View()
		},
	)

	return lipgloss.JoinVertical(
		lipgloss.Left,
		itemViews...,
	)
}
