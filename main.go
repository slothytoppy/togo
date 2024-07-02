package main

import (
	"fmt"
	"github.com/charmbracelet/bubbles/textarea"
	tea "github.com/charmbracelet/bubbletea"
	"os"
)

type model struct {
	ta        textarea.Model
	selection uint
	choices   []string
	buffer    []string
	cursor    uint
	mode      int
	// change mode to be a bool since i only have two views
}

const (
	selection = iota
	file_rendering
)

func (m model) Init() tea.Cmd {
	return textarea.Blink
}

func update_file_rendering(key tea.KeyMsg, m model) (tea.Model, tea.Cmd) {
	switch key.String() {
	case tea.KeyEnter.String():
		//m.buffer = append(m.buffer, m.ta.Value())
	case "esc":
		m.buffer = nil
		m.ta.Reset()
		m.mode = selection
	}
	return m, nil
}

func update_selection(key tea.KeyMsg, m model) (tea.Model, tea.Cmd) {
	switch key.String() {
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
	return m, nil
}

func (m model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	var (
		tiCmd tea.Cmd
	)

	m.ta, tiCmd = m.ta.Update(msg)

	switch msg := msg.(type) {
	case tea.KeyMsg:
		switch msg.String() {
		case "ctrl+c":
			return m, tea.Quit
		}
		if m.mode == file_rendering {
			m, cmd := update_file_rendering(msg, m)
			return m, cmd
		}
		if m.mode == selection {
			m, cmd := update_selection(msg, m)
			return m, cmd
		}
	}
	return m, tea.Batch(tiCmd)
}

func (m model) View() string {
	var s string
	if m.mode == selection {
		s = "which journal would you like to edit?\n"
	}
	for i, choice := range m.choices {
		if m.mode == file_rendering {
			file := "journals/" + m.choices[m.selection]
			s += file + "\n"
			s += m.ta.View()
			/*
				for i := range m.buffer {
					s += string(m.buffer[i] + "\n")
				}
			*/
			return s
		} else if m.mode == selection {
			cursor := " "
			if m.selection == uint(i) {
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
	m.ta = textarea.New()
	m.ta.Focus() // enables typing, removing the focus could probably be used for different textareas

	_, _ = tea.NewProgram(m, tea.WithAltScreen()).Run()
}
