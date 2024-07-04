package view

import (
	tea "github.com/charmbracelet/bubbletea"
)

type Window int

type ChangedWindowMsg struct {
	Filepath string
	Window   Window
}

const (
	Selection = iota
	FileRendering
)

func ChangeWindow(file string, window Window) tea.Msg {
	return ChangedWindowMsg{file, window}
}
