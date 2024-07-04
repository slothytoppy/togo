package files

import (
	"fmt"
	"github.com/charmbracelet/bubbles/textarea"
	tea "github.com/charmbracelet/bubbletea"
	"os"
	"strings"
)

type FileRendererModel struct {
	file_name string
	buffer    []string
	ta        textarea.Model
}

func (file_renderer FileRendererModel) ReadFile(file string) FileRendererModel {
	file_renderer.file_name = file
	buffer, err := os.ReadFile(file)
	if err != nil {
		panic(fmt.Sprintln("skill issues", err.Error()))
	}
	file_renderer.buffer = strings.Split(string(buffer), "\n")
	file_renderer.ta = textarea.New()
	file_renderer.ta.Focus()
	return file_renderer
}

func (f FileRendererModel) Init() tea.Cmd {
	return nil
}

func (f FileRendererModel) Update(msg tea.Msg) (FileRendererModel, tea.Cmd) {
	var (
		tiCmd tea.Cmd
	)

	f.ta, tiCmd = f.ta.Update(msg)

	switch key := msg.(type) {
	case tea.KeyMsg:
		switch key.String() {
		case "ctrl+q":
			return f, tea.Quit
		}
	}
	return f, tea.Batch(tiCmd)
}

func (f FileRendererModel) View() string {
	var s string
	for i := range f.buffer {
		s += fmt.Sprintln(f.buffer[i])
	}
	return s
}
