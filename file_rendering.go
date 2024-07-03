package main

import (
	"fmt"
	"os"
	"strings"

	"github.com/charmbracelet/bubbles/textarea"
	tea "github.com/charmbracelet/bubbletea"
)

type file_renderer struct {
	file_name string
	buffer    []string
	ta        textarea.Model
}

func (f file_renderer) Init() tea.Cmd {
	buffer, err := os.ReadFile(f.file_name)
	if err != nil {
		return tea.Quit
	}
	f.buffer = strings.Split(string(buffer), "\n")
	return nil
}

func (f file_renderer) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	var (
		tiCmd tea.Cmd
	)

	f.ta, tiCmd = f.ta.Update(msg)

	switch key := msg.(type) {
	case tea.KeyMsg:
		switch key.String() {
		case "ctrl+q", "esc":
			return f, tea.Quit
		}
	}
	return f, tea.Batch(tiCmd)
}

func (f file_renderer) View() string {
	var s string
	for i := range f.buffer {
		s += fmt.Sprintln(f.buffer[i])
	}
	return s
}
