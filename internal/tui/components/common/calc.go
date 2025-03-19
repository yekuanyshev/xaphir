package common

import "github.com/charmbracelet/lipgloss"

func CalculateAvailableHeight(height int, views ...string) int {
	availableHeight := height
	for _, view := range views {
		availableHeight -= lipgloss.Height(view)
	}
	return availableHeight
}

func CalculateAvailableWidth(width int, views ...string) int {
	availableWidth := width
	for _, view := range views {
		availableWidth -= lipgloss.Width(view)
	}
	return availableWidth
}

func FillWithEmptySpace(height int) string {
	return lipgloss.NewStyle().Height(height).Render("")
}
