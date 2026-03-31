package ui

import (
	"github.com/axbrunn/tempus/internal/models"
	"github.com/axbrunn/tempus/internal/store"
	"github.com/charmbracelet/bubbles/textinput"
	tea "github.com/charmbracelet/bubbletea"
)

type page int

const (
	pageMenu page = iota
	pageOverview
	pageWithdrawal
	pageAccrual
	pageRapport
)

type model struct {
	page            page
	choices         []string
	cursor          int
	store           store.Store
	inputHr         textinput.Model
	inputDesc       textinput.Model
	activField      int
	withdrawalField withdrawalField
	entryType       models.EntryType
	err             string
}

func initialModel(s store.Store) model {
	hr, desc := initEntry()
	return model{
		page: pageMenu,
		choices: []string{
			"Overzicht weergeven",
			"Uren opnemen",
			"Uren opbouwen",
			"Rapport genereren",
			"Sluiten",
		},
		store:     s,
		inputHr:   hr,
		inputDesc: desc,
	}
}

func (m model) Init() tea.Cmd {
	return nil
}

func (m model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	switch m.page {
	case pageMenu:
		return m.updateMenu(msg)
	case pageOverview:
		return m.updateOverview(msg)
	case pageWithdrawal, pageAccrual:
		return m.updateEntry(msg)
	case pageRapport:
		return m.updateRapport(msg)
	}
	return m, nil
}

func (m model) View() string {
	switch m.page {
	case pageMenu:
		return m.viewMenu()
	case pageOverview:
		return m.viewOverview()
	case pageWithdrawal:
		return m.viewEntry()
	case pageAccrual:
		return m.viewEntry()
	case pageRapport:
		return m.viewRapport()
	}
	return ""
}

func handleChoice(m model) (tea.Model, tea.Cmd) {
	switch m.cursor {
	case 0:
		m.page = pageOverview
	case 1:
		m.entryType = models.Withdrawal
		m.page = pageWithdrawal
	case 2:
		m.entryType = models.Accrual
		m.page = pageAccrual
	case 3:
		_, err := m.store.ExportCSV()
		if err != nil {
			m.err = err.Error()
		}
		m.page = pageRapport
	case 4:
		return m, tea.Quit
	}
	return m, nil
}

func Start(s store.Store) {
	p := tea.NewProgram(initialModel(s))
	p.Run()
}
