package main

import (
	"fmt"
	tea "github.com/charmbracelet/bubbletea"
)

type selection_model struct {
	selection uint
	choices   []string
	mode      bool
}

func (m selection_model) Init() tea.Cmd {
	if len(m.choices) <= 0 {
		panic("len of fis is <=0")
	}
	return nil
}

func (m selection_model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	switch msg := msg.(type) {
	case tea.KeyMsg:
		switch msg.String() {
		case "ctrl+c", "q":
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
			/*
				buffer, _ := os.ReadFile("journals/" + m.choices[m.selection])
					str := strings.Split(string(buffer), "\n")
						for i := range str {
							m.buffer = append(m.buffer, str[i])
						}
			*/
			m.mode = file_rendering
		}
	}

	return m, nil
}

func (m selection_model) View() string {
	s := ""
	s += "what journal would you like to edit?\n"
	cursor := " "
	for i, choice := range m.choices {
		if m.selection == uint(i) {
			cursor = ">"
		}
		s += fmt.Sprintf("%s %s\n", cursor, choice)
	}
	return s
}
