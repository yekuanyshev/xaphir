package chatlist

import (
	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
	"github.com/yekuanyshev/xaphir/pkg/paginator"
	"github.com/yekuanyshev/xaphir/pkg/utils"
	"github.com/yekuanyshev/xaphir/tui/components/dialog"
)

type Component struct {
	width     int
	height    int
	items     []ChatItem
	paginator *Paginator[ChatItem]

	dialog *dialog.Component

	style      lipgloss.Style
	titleStyle lipgloss.Style
}

func NewComponent(
	chats []Chat,
	dialog *dialog.Component,
) *Component {
	items := utils.SliceMap(chats, func(chat Chat) ChatItem {
		return NewChatItem(chat)
	})
	paginatorLimit := 15

	return &Component{
		dialog: dialog,

		style: lipgloss.NewStyle().
			PaddingLeft(1).PaddingRight(1).
			BorderStyle(lipgloss.RoundedBorder()).
			BorderForeground(lipgloss.Color("36")),

		titleStyle: lipgloss.NewStyle().
			PaddingLeft(1).PaddingRight(1).
			MarginBottom(1).
			Foreground(lipgloss.Color("#fff")).
			Background(lipgloss.Color("62")).
			Bold(true),

		items:     items,
		paginator: NewPaginator(paginator.NewItemPaginator(items, paginatorLimit)),
	}
}

func (cl *Component) Init() tea.Cmd {
	return nil
}

func (cl *Component) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	switch msg := msg.(type) {
	case tea.KeyMsg:
		switch msg.String() {
		case "down":
			cl.paginator.Increment()
		case "up":
			cl.paginator.Decrement()
		case "right":
			cl.paginator.SkipToNextPage()
		case "left":
			cl.paginator.SkipToPrevPage()
		}
	}

	cl.dialog.SetTitle(cl.paginator.CurrentItem().Username)

	return cl, nil
}

func (cl *Component) View() string {
	cl.style = cl.style.Width(cl.width).Height(cl.height)
	cl.paginator.SetWidth(cl.width - cl.style.GetHorizontalFrameSize())

	var sections []string
	availHeight := cl.style.GetHeight()

	titleView := cl.titleStyle.Render("Chats")
	sections = append(sections, titleView)
	availHeight -= lipgloss.Height(titleView)

	itemsView := cl.itemsView()
	sections = append(sections, itemsView)
	availHeight -= lipgloss.Height(itemsView)

	// append empty space
	sections = append(sections, lipgloss.NewStyle().Height(availHeight).Render(""))

	paginatorView := cl.paginator.View()
	sections = append(sections, paginatorView)

	return cl.style.Render(
		lipgloss.JoinVertical(
			lipgloss.Left,
			sections...,
		),
	)
}

func (cl *Component) SetWidth(w int) {
	cl.width = w
}

func (cl *Component) SetHeight(h int) {
	cl.height = h
}

func (cl *Component) itemsView() string {
	var items []string
	for i, chatItem := range cl.paginator.ItemsOnCurrentPage() {
		isSelected := cl.paginator.Cursor() == i
		itemView := chatItem.View(isSelected)
		items = append(items, itemView)
	}

	return lipgloss.JoinVertical(lipgloss.Left, items...)
}
