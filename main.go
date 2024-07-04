package main

import (
	"log"
	"togo/files"
	"togo/selection"
	"togo/view"

	tea "github.com/charmbracelet/bubbletea"
)

/*
model should just be a "view" into the currently selected model

	it should just display whatever model its told is selected
*/
type model struct {
	file_model      files.FileRendererModel
	selection_model selection.SelectionModel
	view            view.Window
}

func (m model) Init() tea.Cmd {
	return nil
}

func (m model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	logger := log.New(writer{}, "", 0o655)
	switch msg := msg.(type) {
	case view.ChangedWindowMsg:
		m.file_model = m.file_model.ReadFile(msg.Filepath)
		logger.Println(m.view)
		m.view = msg.Window
	}

	switch m.view {
	case view.Selection:
		var cmd tea.Cmd
		m.selection_model, cmd = m.selection_model.Update(msg)
		return m, cmd
	case view.FileRendering:
		var cmd tea.Cmd
		m.file_model, cmd = m.file_model.Update(msg)
		return m, cmd
	}
	/*
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
	*/
	return m, nil
}

func (m model) View() string {
	switch m.view {
	case view.Selection:
		return m.selection_model.View()
	case view.FileRendering:
		return m.file_model.View()
	}
	panic("")
}

func main() {
	m := model{view: view.Selection, file_model: files.FileRendererModel{}, selection_model: selection.SelectionModel{}.ReadDir("journals")}
	_, _ = tea.NewProgram(m, tea.WithAltScreen()).Run()
}
