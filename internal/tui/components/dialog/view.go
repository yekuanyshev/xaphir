package dialog

import (
	"github.com/charmbracelet/bubbles/textinput"
	"github.com/charmbracelet/lipgloss"
	"github.com/yekuanyshev/xaphir/internal/tui/components/common"
	"github.com/yekuanyshev/xaphir/pkg/utils"
)

type view struct {
	width  int
	height int

	style             lipgloss.Style
	titleStyle        lipgloss.Style
	blurredTitleStyle lipgloss.Style
	inputStyle        lipgloss.Style
}

func newView() *view {
	style := lipgloss.NewStyle().
		PaddingLeft(1).PaddingRight(1).
		BorderStyle(lipgloss.RoundedBorder()).
		BorderForeground(lipgloss.Color("36"))

	titleStyle := lipgloss.NewStyle().
		PaddingLeft(1).PaddingRight(1).
		MarginBottom(1).
		Foreground(lipgloss.Color("#fff")).
		Background(lipgloss.Color("62")).
		Bold(true)

	blurredTitleStyle := lipgloss.NewStyle().
		Faint(true)

	inputStyle := lipgloss.NewStyle().
		PaddingLeft(1).PaddingRight(1).
		BorderStyle(lipgloss.RoundedBorder())

	return &view{
		style:             style,
		titleStyle:        titleStyle,
		blurredTitleStyle: blurredTitleStyle,
		inputStyle:        inputStyle,
	}
}

func (v *view) setWidth(width int) {
	v.width = width
	v.style = v.style.Width(width)
	v.inputStyle = v.inputStyle.Width(v.innerWidth() - v.inputStyle.GetHorizontalFrameSize())
}

func (v *view) setHeight(height int) {
	v.height = height
	v.style = v.style.Height(height)
}

func (v *view) innerWidth() int {
	return v.width - v.style.GetHorizontalFrameSize()
}

func (v *view) innerHeight() int {
	return v.height - v.style.GetVerticalFrameSize()
}

func (v *view) render(
	focus bool,
	blurredTitle string,
	title string,
	slider *Slider,
	input textinput.Model,
) string {
	if !focus {
		return v.blurredView(blurredTitle)
	}

	titleView := v.renderTitle(title)
	itemsView := v.renderItems(slider)
	inputView := v.renderInput(input)
	availableHeight := common.CalculateAvailableHeight(
		v.innerHeight(), titleView, itemsView, inputView,
	)
	emptySpace := common.VerticalGap(availableHeight)

	sections := []string{
		titleView,
		emptySpace,
		itemsView,
		inputView,
	}

	return v.style.Render(
		lipgloss.JoinVertical(
			lipgloss.Left,
			sections...,
		),
	)
}

func (v *view) renderTitle(title string) string {
	return v.titleStyle.Render(title)
}

func (v *view) renderInput(input textinput.Model) string {
	return v.inputStyle.Render(input.View())
}

func (v *view) renderItems(slider *Slider) string {
	itemViews := utils.SliceMap(
		slider.GetItems(),
		func(item Item) string {
			return item.view(v.innerWidth())
		},
	)

	return lipgloss.JoinVertical(
		lipgloss.Left,
		itemViews...,
	)
}

func (v *view) blurredView(blurredTitle string) string {
	return v.style.Render(
		lipgloss.Place(
			v.innerWidth(), v.innerHeight(),
			lipgloss.Center, lipgloss.Center,
			v.blurredTitleStyle.Render(blurredTitle),
		),
	)
}

func (v *view) inputFocus() {
	v.inputStyle = v.inputStyle.BorderForeground(lipgloss.Color("36"))
}

func (v *view) inputBlur() {
	v.inputStyle = v.inputStyle.UnsetBorderForeground()
}
