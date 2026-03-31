package ui

import (
	"strings"

	tea "github.com/charmbracelet/bubbletea"
)

func (m model) viewRapport() string {
	var s strings.Builder

	if m.err != "" {
		s.WriteString(errorStyle.Render("✗ Fout: " + m.err))
	} else {
		s.WriteString(successStyle.Render("✓ CSV geëxporteerd naar Downloads!"))
	}

	s.WriteString("\n\n" + subtleStyle.Render("Esc om terug te gaan"))

	return appStyle.Render(logoStyle.Render(logo) + "\n\n" + s.String())
}

func (m model) updateRapport(msg tea.Msg) (tea.Model, tea.Cmd) {
	switch msg := msg.(type) {
	case tea.KeyMsg:
		switch msg.String() {
		case "esc":
			m.page = pageMenu
		}
	}
	return m, nil
}
