package ui

import "github.com/charmbracelet/lipgloss"

var (
	// Kleuren
	purple    = lipgloss.Color("#7D56F4")
	green     = lipgloss.Color("#04B575")
	red       = lipgloss.Color("#FF4444")
	gray      = lipgloss.Color("#626262")
	lightGray = lipgloss.Color("#9B9B9B")

	// Logo
	logoStyle = lipgloss.NewStyle().
			Foreground(purple).
			Bold(true)

	// Container om alles heen
	appStyle = lipgloss.NewStyle().
			Padding(1, 2)

	// Titel van een pagina
	titleStyle = lipgloss.NewStyle().
			Foreground(purple).
			Bold(true).
			BorderBottom(true).
			BorderStyle(lipgloss.NormalBorder()).
			BorderForeground(purple).
			MarginBottom(1)

	// Menu item geselecteerd
	selectedStyle = lipgloss.NewStyle().
			Foreground(purple).
			Bold(true)

	// Menu item normaal
	normalStyle = lipgloss.NewStyle().
			Foreground(lightGray)

	// Saldo positief
	balancePositiveStyle = lipgloss.NewStyle().
				Foreground(green).
				Bold(true)

	// Saldo negatief
	balanceNegativeStyle = lipgloss.NewStyle().
				Foreground(red).
				Bold(true)

	// Foutmelding
	errorStyle = lipgloss.NewStyle().
			Foreground(red).
			Bold(true)

	// Success melding
	successStyle = lipgloss.NewStyle().
			Foreground(green).
			Bold(true)

	// Subtekst
	subtleStyle = lipgloss.NewStyle().
			Foreground(gray)

	headerStyle = lipgloss.NewStyle().
			Foreground(purple).
			Bold(true).
			Align(lipgloss.Center)

	oddRowStyle = lipgloss.NewStyle().
			Foreground(lightGray).
			Padding(0, 1)

	evenRowStyle = lipgloss.NewStyle().
			Foreground(gray).
			Padding(0, 1)
)
