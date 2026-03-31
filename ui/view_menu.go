package ui

import (
	"strings"

	tea "github.com/charmbracelet/bubbletea"
)

func (m model) updateMenu(msg tea.Msg) (tea.Model, tea.Cmd) {
	switch msg := msg.(type) {
	case tea.KeyMsg:
		switch msg.String() {
		case "up", "k":
			if m.cursor > 0 {
				m.cursor--
			}
		case "down", "j":
			if m.cursor < len(m.choices)-1 {
				m.cursor++
			}
		case "enter":
			switch m.page {
			case pageMenu:
				return handleChoice(m)
			default:
				m.page = pageMenu
			}

		case "q", "ctrl+c":
			return m, tea.Quit
		}

	}

	return m, nil
}

func (m model) viewMenu() string {
	var s strings.Builder

	s.WriteString(logoStyle.Render(logo) + "\n")
	s.WriteString(titleStyle.Render("Hoofdmenu") + "\n")

	for i, choice := range m.choices {
		cursor := "  "
		style := normalStyle
		if m.cursor == i {
			cursor = "▸ "
			style = selectedStyle
		}
		s.WriteString(cursor + style.Render(choice) + "\n")
	}

	s.WriteString("\n" + subtleStyle.Render("↑/↓ navigeren • enter selecteren"))

	return appStyle.Render(s.String())
}
