package ui

import (
	"fmt"
	"strings"

	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
	"github.com/charmbracelet/lipgloss/table"
)

func (m model) updateOverview(msg tea.Msg) (tea.Model, tea.Cmd) {
	switch msg := msg.(type) {
	case tea.KeyMsg:
		switch msg.String() {
		case "esc":
			m.page = pageMenu
		}
	}
	return m, nil
}

func (m model) viewOverview() string {
	var s strings.Builder

	balance := m.store.CalculateBalance()

	// Tabel rijen opbouwen
	t := table.New().
		Border(lipgloss.NormalBorder()).
		BorderStyle(lipgloss.NewStyle().Foreground(purple)).
		Headers("Datum", "Omschrijving", "Type", "Uren").
		StyleFunc(func(row, col int) lipgloss.Style {
			switch {
			case row == table.HeaderRow:
				return headerStyle
			case row%2 == 0:
				return evenRowStyle
			default:
				return oddRowStyle
			}
		})

	for _, e := range m.store.Entries {
		t.Row(
			e.Date.Format("02-01-2006"),
			e.Description,
			string(e.Type),
			fmt.Sprintf("%.2f", e.Hours),
		)
	}

	s.WriteString(titleStyle.Render("Overzicht") + "\n")
	s.WriteString(t.Render() + "\n\n")

	// Saldo
	balanceLabel := "Saldo: "
	if balance >= 0 {
		s.WriteString(balanceLabel + balancePositiveStyle.Render(fmt.Sprintf("%.2f uur", balance)) + "\n")
	} else {
		s.WriteString(balanceLabel + balanceNegativeStyle.Render(fmt.Sprintf("%.2f uur", balance)) + "\n")
	}

	s.WriteString("\n" + subtleStyle.Render("Esc om terug te gaan"))

	return appStyle.Render(logoStyle.Render(logo) + "\n" + s.String())
}
