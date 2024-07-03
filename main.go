package main

import (
	"log"
	"os"

	"github.com/charmbracelet/bubbles/textarea"
	tea "github.com/charmbracelet/bubbletea"
)

type model struct {
	selection_model selection_model
	file_model      file_model
	mode            bool
	// change mode to be a bool since i only have two views
}

const (
	selection      = false
	file_rendering = true
)

func (m model) Init() tea.Cmd {
	return textarea.Blink
}

func (m model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	switch key := msg.(type) {
	case tea.KeyMsg:
		if m.mode == file_rendering {
			m, cmd := m.file_model.Update(key)
			return m, cmd
		}
		if m.mode == selection {
			m.file_model.file_name = m.selection_model.choices[m.selection_model.selection]
			m.file_model.ta = textarea.New()
			m.file_model.ta.Focus()
			m, cmd := m.file_model.Update(key)
			return m, cmd
		}
	}
	return m, nil
}

func (m model) View() string {
	if m.mode == file_rendering {
		return m.file_model.View()
	} else if m.mode == selection {
		return m.selection_model.View()
	}
	return ""
}

func main() {
	m := model{}
	m.mode = selection
	b, _ := os.ReadDir("journals")
	for _, choices := range b {
		m.selection_model.choices = append(m.selection_model.choices, choices.Name())
	}
	l := log.New(writer{}, "", 0)
	l.Println("hello")
	_, _ = tea.NewProgram(m, tea.WithAltScreen()).Run()
}
