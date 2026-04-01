package ui

import (
	"strings"

	"github.com/axbrunn/tempus/internal/store"
	"github.com/charmbracelet/bubbles/textinput"
	tea "github.com/charmbracelet/bubbletea"
)

func initFilePicker() textinput.Model {
	name := textinput.New()
	name.Placeholder = "bijv. werk-2026"
	name.CharLimit = 40

	return name
}

func (m model) viewFilePicker() string {
	var s strings.Builder

	s.WriteString(titleStyle.Render("Kies een bestand") + "\n")

	if m.creatingNew {
		s.WriteString("Naam nieuw bestand:\n")
		s.WriteString(m.inputName.View() + "\n")
		if m.err != "" {
			s.WriteString(errorStyle.Render(m.err) + "\n")
		}
		s.WriteString(subtleStyle.Render("enter om aan te maken • esc om te annuleren"))
	} else {
		if len(m.stores) == 0 {
			s.WriteString(subtleStyle.Render("Geen bestanden gevonden.") + "\n")
		} else {
			for i, f := range m.stores {
				cursor := "  "
				style := normalStyle
				if m.cursor == i {
					cursor = "▸ "
					style = selectedStyle
				}
				s.WriteString(cursor + style.Render(f.Path) + "\n")
			}
			s.WriteString("\n")
		}
		s.WriteString(subtleStyle.Render("↑/↓ navigeren • enter selecteren • n nieuw bestand • q afsluiten"))
	}

	return appStyle.Render(logoStyle.Render(logo) + "\n" + s.String())
}

func (m model) updateFilePicker(msg tea.Msg) (tea.Model, tea.Cmd) {
	switch msg := msg.(type) {
	case tea.KeyMsg:
		if m.creatingNew {
			switch msg.String() {
			case "esc":
				m.creatingNew = false
				m.inputName.SetValue("")
				m.err = ""
			case "enter":
				name := m.inputName.Value()
				if name == "" {
					m.err = "Naam mag niet leeg zijn"
					return m, nil
				}
				s, err := store.CreateFile(name)
				if err != nil {
					m.err = err.Error()
					return m, nil
				}
				m.store = s
				m.creatingNew = false
				m.inputName.SetValue("")
				m.page = pageMenu
				return m, loadStoresCmd
			default:
				var cmd tea.Cmd
				m.inputName, cmd = m.inputName.Update(msg)
				return m, cmd
			}
			return m, nil
		}

		switch msg.String() {
		case "up", "k":
			if m.cursor > 0 {
				m.cursor--
			}
		case "down", "j":
			if m.cursor < len(m.stores)-1 {
				m.cursor++
			}
		case "enter":
			if len(m.stores) > 0 {
				m.store = m.stores[m.cursor]
				m.cursor = 0
				m.page = pageMenu
			}
		case "n":
			m.creatingNew = true
			m.inputName.Focus()
		case "q", "ctrl+c":
			return m, tea.Quit
		}
	}
	return m, nil
}
