package ui

import (
	"strconv"
	"strings"

	"github.com/axbrunn/tempus/internal/models"
	"github.com/charmbracelet/bubbles/textinput"
	tea "github.com/charmbracelet/bubbletea"
)

type withdrawalField int

const (
	fieldHr withdrawalField = iota
	fieldDesc
)

func initEntry() (textinput.Model, textinput.Model) {
	hr := textinput.New()
	hr.Placeholder = "bijv. 8"
	hr.Focus()

	desc := textinput.New()
	desc.Placeholder = "bijv. week 14"

	return hr, desc
}

func (m model) updateEntry(msg tea.Msg) (tea.Model, tea.Cmd) {
	switch msg := msg.(type) {
	case tea.KeyMsg:
		switch msg.String() {
		case "tab", "down":
			if m.activField == 0 {
				m.activField = 1
				m.inputHr.Blur()
				m.inputDesc.Focus()
			}
		case "shift+tab", "up":
			if m.activField == 1 {
				m.activField = 0
				m.inputDesc.Blur()
				m.inputHr.Focus()
			}
		case "enter":
			if m.activField == 1 {
				return m.saveEntry()
			}
			m.activField = 1
			m.inputHr.Blur()
			m.inputDesc.Focus()
		case "esc":
			m.page = pageMenu
			return m, nil
		}
	}

	var cmd tea.Cmd
	if m.activField == 0 {
		m.inputHr, cmd = m.inputHr.Update(msg)
	} else {
		m.inputDesc, cmd = m.inputDesc.Update(msg)
	}
	return m, cmd
}

func (m model) saveEntry() (tea.Model, tea.Cmd) {
	hours, err := strconv.ParseFloat(m.inputHr.Value(), 64)
	if err != nil || hours <= 0 {
		m.err = "Vul een geldig aantal uren in"
		return m, nil
	}

	entry := models.NewEntry(hours, m.inputDesc.Value(), m.entryType)
	m.store.Entries = append(m.store.Entries, entry)
	m.store.Save()

	// reset
	m.inputHr.SetValue("")
	m.inputDesc.SetValue("")
	m.activField = 0
	m.inputHr.Focus()
	m.inputDesc.Blur()
	m.err = ""
	m.page = pageMenu
	return m, nil
}

func (m model) viewEntry() string {
	var s strings.Builder

	titel := "Uren opbouwen"
	if m.entryType == models.Withdrawal {
		titel = "Uren opnemen"
	}

	s.WriteString(titleStyle.Render(titel) + "\n")

	s.WriteString("Uren:         " + m.inputHr.View() + "\n")
	s.WriteString("Omschrijving: " + m.inputDesc.View() + "\n")

	if m.err != "" {
		s.WriteString("\n" + errorStyle.Render(m.err) + "\n")
	}

	s.WriteString("\n" + subtleStyle.Render("tab om te wisselen • enter om op te slaan • esc om te annuleren"))

	return appStyle.Render(logoStyle.Render(logo) + "\n\n" + s.String())
}
