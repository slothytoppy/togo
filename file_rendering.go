package main

import (
	"fmt"
	"github.com/charmbracelet/bubbles/textarea"
	tea "github.com/charmbracelet/bubbletea"
	"os"
	"strings"
)

func load_file(file string) []string {
	if s, err := os.ReadFile(file); err.Error() != "" {
		return strings.Split(string(s), "\n")
	}
	return []string{}
}

type file_model struct {
	ta        textarea.Model
	buffer    []string
	file_name string
}

func (m file_model) Init() tea.Cmd {
	return textarea.Blink
}

func (m file_model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
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
		return m, nil
	}
	return m, tea.Batch(tiCmd)
}

func (m file_model) View() string {
	var s string
	file := "journals/" + m.file_name
	s += file + "\n"
	s += m.ta.View()
	/*
		for i := range m.buffer {
			s += string(m.buffer[i] + "\n")
		}
	*/
	s += fmt.Sprintln(m.buffer)
	return s
}
