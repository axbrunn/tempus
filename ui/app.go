package ui

import (
	"fmt"
	"os"

	"github.com/axbrunn/tempus/internal/models"
	"github.com/axbrunn/tempus/internal/store"
	"github.com/charmbracelet/bubbles/textinput"
	tea "github.com/charmbracelet/bubbletea"
)

type page int
type storesLoadedMsg []store.Store

const (
	pageFilePicker page = iota
	pageMenu
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
	stores          []store.Store
	inputName       textinput.Model
	creatingNew     bool
	inputHr         textinput.Model
	inputDesc       textinput.Model
	activField      int
	withdrawalField withdrawalField
	entryType       models.EntryType
	err             string
}

func loadStoresCmd() tea.Msg {
	stores, _ := store.ListFiles()
	return storesLoadedMsg(stores)
}

func initialModel(stores []store.Store) model {
	hr, desc := initEntry()
	name := initFilePicker()
	return model{
		page: pageFilePicker,
		choices: []string{
			"Overzicht weergeven",
			"Uren opnemen",
			"Uren opbouwen",
			"Rapport genereren",
			"Bestand kiezen",
			"Sluiten",
		},
		stores:    stores,
		inputName: name,
		inputHr:   hr,
		inputDesc: desc,
	}
}

func (m model) Init() tea.Cmd {
	return nil
}

func (m model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	switch msg := msg.(type) {
	case storesLoadedMsg:
		m.stores = []store.Store(msg)
		return m, nil
	}
	switch m.page {
	case pageFilePicker:
		return m.updateFilePicker(msg)
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
	case pageFilePicker:
		return m.viewFilePicker()
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
		m.cursor = 0
		m.page = pageFilePicker
	case 5:
		return m, tea.Quit
	}
	return m, nil
}

func Start() {
	stores, err := store.ListFiles()
	if err != nil {
		fmt.Println("Failed to load files:", err)
		os.Exit(1)
	}

	p := tea.NewProgram(initialModel(stores))
	p.Run()
}
