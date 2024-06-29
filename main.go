package main

import (
	"fmt"
	// "github.com/charmbracelet/bubbles/viewport"
	tea "github.com/charmbracelet/bubbletea"
	//"github.com/charmbracelet/glamour"
	"os"
)

type model struct {
	cursor  uint
	choices []string
	buffer  []byte
	mode    int
}

const (
	selection = iota
	file_rendering
)

func (m model) Init() tea.Cmd {
	return nil
}

func (m model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	switch msg := msg.(type) {
	case tea.KeyMsg:
		switch msg.String() {
		case "ctrl+c", "q":
			return m, tea.Quit
		}
		if m.mode == file_rendering {
		}
		if m.mode == selection {
			switch msg.String() {
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
			}
		}
	}
	return m, nil
}

func (m model) View() string {
	var s string
	if m.mode == selection {
		s = "which journal would you like to edit?\n"
	}
	for i, choice := range m.choices {
		if m.mode == file_rendering {
			file := "journals/" + m.choices[m.cursor]
			s += file + "\n"
			if len(m.buffer) <= 0 {
				m.buffer, _ = os.ReadFile(file)
			}
			s += string(m.buffer)
			break
		} else if m.mode == selection {
			cursor := " "
			if m.cursor == uint(i) {
				cursor = ">"
			}
			s += fmt.Sprintf("%s %s\n", cursor, choice)
		}
	}
	return s
}

func main() {
	m := model{}
	m.mode = selection
	storage_name := "journals"
	if _, err := os.Stat(storage_name); os.IsNotExist(err) {
		os.Mkdir(storage_name, 0o755)
		fmt.Println("hello")
	}
	fis, err := os.ReadDir(storage_name)
	if err != nil {
		fmt.Println(err)
		return
	}
	for _, fi := range fis {
		m.choices = append(m.choices, fi.Name())
	}
	/*	var journal_names []string
		var file_info = make([][]byte, len(fis))
	*/
	_, _ = tea.NewProgram(m, tea.WithAltScreen()).Run()
}
