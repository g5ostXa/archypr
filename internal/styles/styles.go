package styles

import (
	"charm.land/lipgloss/v2"
)

// Main colors
var (
	white = lipgloss.Color("#FFFFFF")
	blue  = lipgloss.Color("#2436c5")
	green = lipgloss.Color("#1fc876")
	//purple = lipgloss.Color("#8010d0")
)

// Master header style
var (
	MasterTileStyle = lipgloss.NewStyle().
			Bold(true).
			Align(lipgloss.Center).
			Foreground((white)).
			BorderStyle(lipgloss.RoundedBorder()).
			BorderForeground((blue)).
			Width(24)

	MasterVersionStyle = lipgloss.NewStyle().
				Bold(true).
				Align(lipgloss.Center).
				Foreground((blue)).
				Width(15).
				PaddingLeft(1)

	CommonSeparatorStyle = lipgloss.NewStyle().
				Bold(true).
				Foreground((green)).
				PaddingLeft(1)

	CommonPromptStyle = lipgloss.NewStyle().
				Bold(true).
				Align(lipgloss.Center).
				Foreground((white)).
				MarginTop(1).
				PaddingLeft(1)
)
