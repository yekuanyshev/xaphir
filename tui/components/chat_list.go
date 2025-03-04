package components

import (
	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
	"github.com/yekuanyshev/xaphir/pkg/paginator"
)

type Chat struct {
	Username    string
	LastMessage string
}

type ChatListComponent struct {
	style                        lipgloss.Style
	titleStyle                   lipgloss.Style
	itemStyle                    lipgloss.Style
	itemTitleStyle               lipgloss.Style
	itemDescriptionStyle         lipgloss.Style
	selectedItemTitleStyle       lipgloss.Style
	selectedItemDescriptionStyle lipgloss.Style
	paginatorStyle               lipgloss.Style

	items     []Chat
	paginator *paginator.ItemPaginator[Chat]
}

func NewChatList(width, height int, items []Chat) *ChatListComponent {
	return &ChatListComponent{
		style: lipgloss.NewStyle().
			Width(width).Height(height).
			PaddingLeft(1).PaddingRight(1).
			BorderStyle(lipgloss.RoundedBorder()).
			BorderForeground(lipgloss.Color("36")),

		titleStyle: lipgloss.NewStyle().
			PaddingLeft(1).PaddingRight(1).
			MarginBottom(1).
			Foreground(lipgloss.Color("#fff")).
			Background(lipgloss.Color("62")).
			Bold(true),

		itemStyle: lipgloss.NewStyle().
			Width(width).MarginBottom(1),

		itemTitleStyle: lipgloss.NewStyle().
			Foreground(lipgloss.Color("#fff")).
			Bold(true),

		itemDescriptionStyle: lipgloss.NewStyle().
			Faint(true),

		selectedItemTitleStyle: lipgloss.NewStyle().
			Foreground(lipgloss.Color("62")).
			Bold(true),

		selectedItemDescriptionStyle: lipgloss.NewStyle().
			Faint(true),

		paginatorStyle: lipgloss.NewStyle().
			Width(width).
			AlignHorizontal(lipgloss.Center),

		items:     items,
		paginator: paginator.NewItemPaginator(items, 15),
	}
}

func (clc *ChatListComponent) Init() tea.Cmd {
	return nil
}

func (clc *ChatListComponent) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	switch msg := msg.(type) {
	case tea.KeyMsg:
		switch msg.String() {
		case "ctrl+c", "q":
			return clc, tea.Quit
		case "down":
			clc.paginator.Increment()
		case "up":
			clc.paginator.Decrement()
		case "right":
			clc.paginator.SkipToNextPage()
		case "left":
			clc.paginator.SkipToPrevPage()
		}
	}

	return clc, nil
}

func (clc *ChatListComponent) View() string {
	var sections []string
	availHeight := clc.style.GetHeight()

	titleView := clc.titleStyle.Render("Chats")
	sections = append(sections, titleView)
	availHeight -= lipgloss.Height(titleView)

	itemsView := clc.itemsView()
	sections = append(sections, itemsView)
	availHeight -= lipgloss.Height(itemsView)

	// append empty space
	sections = append(sections, lipgloss.NewStyle().Height(availHeight).Render(""))

	paginatorView := clc.paginatorStyle.Render(clc.paginatorView())
	sections = append(sections, paginatorView)

	return clc.style.Render(
		lipgloss.JoinVertical(
			lipgloss.Left,
			sections...,
		),
	)
}

func (clc *ChatListComponent) paginatorView() string {
	activeDot := "○"
	inactiveDot := "•"

	view := ""
	for page := range clc.paginator.TotalPages() {
		if page == clc.paginator.CurrentPage() {
			view += activeDot
		} else {
			view += inactiveDot
		}
	}

	return view
}

func (clc *ChatListComponent) itemsView() string {
	var items []string
	for i, chatItem := range clc.paginator.ItemsOnCurrentPage() {
		isSelected := clc.paginator.Cursor() == i
		itemView := clc.itemView(chatItem, isSelected)
		items = append(items, itemView)
	}

	return clc.itemStyle.Render(lipgloss.JoinVertical(lipgloss.Left, items...))
}

func (clc *ChatListComponent) itemView(item Chat, isSelected bool) string {
	view := lipgloss.JoinVertical(
		lipgloss.Left,
		clc.itemTitleStyle.Render(item.Username),
		clc.itemDescriptionStyle.Render(item.LastMessage),
	)

	if isSelected {
		view = lipgloss.JoinVertical(
			lipgloss.Left,
			clc.selectedItemTitleStyle.Render(item.Username),
			clc.selectedItemDescriptionStyle.Render(item.LastMessage),
		)
	}

	return clc.itemStyle.Render(view)
}
