package main

import (
	"fmt"
	//"github.com/charmbracelet/bubbles/textarea"
	"log"
	"os"

	"github.com/charmbracelet/bubbles/textarea"
	tea "github.com/charmbracelet/bubbletea"
)

type model struct {
	cursor  uint
	choices []string
	mode    uint
	logger  *log.Logger
	f       file_renderer
}

const (
	selection = iota
	file_rendering
)

func (m model) Init() tea.Cmd {
	return nil
}

func (m model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	if m.mode == file_rendering {
		m.logger.Println("chosen")
		return m.f.Update(msg)
	}
	if m.mode == selection {
		switch key := msg.(type) {
		case tea.KeyMsg:
			switch key.String() {
			case "ctrl+q", "esc":
				return m, tea.Quit
			case "up", "w", "k":
				if m.cursor > 0 {
					m.cursor -= 1
				}
			case "down", "s", "j":
				if m.cursor < uint(len(m.choices)-1) {
					m.cursor += 1
				}
			case "enter", " ":
				m.mode = file_rendering
				m.f.ta = textarea.New()
				return m, m.f.ta.Focus()
			}
		}
	}
	return m, nil
}

func (m model) View() string {
	s := "what journal would you like to edit\n"
	for i, choice := range m.choices {
		cursor := " "
		if m.cursor == uint(i) {
			cursor = ">"
		}
		s += fmt.Sprintf("%s %s\n", cursor, choice)
	}
	return s
}

func main() {
	m := model{}
	b, _ := os.ReadDir("journals")
	for _, choices := range b {
		m.choices = append(m.choices, choices.Name())
	}
	l := log.New(writer{}, "", 0)
	m.logger = l
	//m.file_model.ta = textarea.New()
	_, _ = tea.NewProgram(m, tea.WithAltScreen()).Run()
}
