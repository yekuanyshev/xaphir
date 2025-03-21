package chatlist

import (
	"github.com/charmbracelet/bubbles/textinput"
	"github.com/charmbracelet/lipgloss"
	"github.com/yekuanyshev/xaphir/internal/tui/components/common"
	"github.com/yekuanyshev/xaphir/pkg/paginator"
	"github.com/yekuanyshev/xaphir/pkg/utils"
)

type view struct {
	width  int
	height int

	style            lipgloss.Style
	titleStyle       lipgloss.Style
	filterInputStyle lipgloss.Style
	paginatorStyle   lipgloss.Style
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

	filterInputStyle := lipgloss.NewStyle().
		PaddingLeft(1).PaddingRight(1).
		BorderStyle(lipgloss.RoundedBorder()).
		BorderForeground(lipgloss.Color("36"))

	paginatorStyle := lipgloss.NewStyle().
		AlignHorizontal(lipgloss.Center)

	return &view{
		style:            style,
		titleStyle:       titleStyle,
		filterInputStyle: filterInputStyle,
		paginatorStyle:   paginatorStyle,
	}
}

func (v *view) setWidth(width int) {
	v.width = width
	v.style = v.style.Width(width)
	v.filterInputStyle = v.filterInputStyle.Width(
		v.innerWidth() - v.filterInputStyle.GetHorizontalFrameSize(),
	)
	v.paginatorStyle = v.paginatorStyle.Width(v.innerWidth())
}

func (v *view) setHeight(height int) {
	v.height = height
	v.style = v.style.Height(height)
}

func (v *view) focus() {
	v.style = v.style.Faint(false)
	v.titleStyle = v.titleStyle.Faint(false)
}

func (v *view) blur() {
	v.style = v.style.Faint(true)
	v.titleStyle = v.titleStyle.Faint(true)
}

func (v *view) innerWidth() int {
	return v.width - v.style.GetHorizontalFrameSize()
}

func (v *view) innerHeight() int {
	return v.height - v.style.GetVerticalFrameSize()
}

func (v *view) render(
	title string,
	filter *filter,
	paginator *paginator.ItemPaginator[Item],
) string {
	titleview := v.renderTitle(title)
	headerview := titleview

	if filter.enabled {
		filterInputview := v.renderFilterInput(filter.input)
		headerview = filterInputview
	}

	itemsview := v.renderItems(paginator)
	paginatorview := v.renderPaginator(paginator)
	availableHeight := common.CalculateAvailableHeight(
		v.innerHeight(), headerview, itemsview, paginatorview,
	)
	emptySpace := common.VerticalGap(availableHeight)

	sections := []string{
		headerview,
		itemsview,
		emptySpace,
		paginatorview,
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

func (v *view) renderFilterInput(filterInput textinput.Model) string {
	return v.filterInputStyle.Render(filterInput.View())
}

func (v *view) renderItems(paginator *paginator.ItemPaginator[Item]) string {
	itemviews := utils.SliceMap(
		paginator.ItemsOnCurrentPage(),
		v.renderItem,
	)

	return lipgloss.JoinVertical(
		lipgloss.Left,
		itemviews...,
	)
}

func (v *view) renderItem(item Item) string {
	return item.view(v.innerWidth())
}

func (v *view) renderPaginator(paginator *paginator.ItemPaginator[Item]) string {
	return v.paginatorStyle.Render(paginator.String())
}
