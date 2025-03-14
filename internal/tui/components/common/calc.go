package common

import "github.com/charmbracelet/lipgloss"

func CalculateAvailableHeight(height int, views ...string) int {
	availableHeight := height
	for _, view := range views {
		availableHeight -= lipgloss.Height(view)
	}
	return availableHeight
}

func FillWithEmptySpace(height int) string {
	return lipgloss.NewStyle().Height(height).Render("")
}
