package main

import (
	"github.com/charmbracelet/bubbles/textarea"
	tea "github.com/charmbracelet/bubbletea"
)

type selection_model struct {
	ta        textarea.Model
	selection uint
	choices   []string
	mode      int
}

func (m selection_model) Init() tea.Cmd {
	return textarea.Blink
}

func (m selection_model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	switch msg := msg.(type) {
	case tea.KeyMsg:
		switch msg.String() {
		case "q":
			return m, tea.Quit
		case "up", "w", "k":
			if m.selection > 0 {
				m.selection -= 1
			}
		case "down", "s", "j":
			if m.selection < uint(len(m.choices)-1) {
				m.selection += 1
			}
		case "enter", " ":
			m.ta.Reset()
			/*
				buffer, _ := os.ReadFile("journals/" + m.choices[m.selection])
					str := strings.Split(string(buffer), "\n")
						for i := range str {
							m.buffer = append(m.buffer, str[i])
						}
			*/
			m.mode = file_rendering
			m.ta.SetHeight(47)
		}
	}

	return m, nil
}

func (m selection_model) View() string {
	return ""
}
