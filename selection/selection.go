package selection

import (
	"fmt"
	tea "github.com/charmbracelet/bubbletea"
	"os"
	"togo/view"
)

type SelectionModel struct {
	choices  []string
	selected int
}

func (selection SelectionModel) Init() tea.Cmd {
	return nil
}

func (selection SelectionModel) ReadDir(dir string) SelectionModel {
	if fis, err := os.ReadDir(dir); err == nil {
		for _, fi := range fis {
			selection.choices = append(selection.choices, dir+"/"+fi.Name())
		}
	} else {
		fmt.Println(dir, err)
		panic("")
	}
	return selection
}

func (selection SelectionModel) Update(msg tea.Msg) (SelectionModel, tea.Cmd) {
	switch key := msg.(type) {
	case tea.KeyMsg:
		switch key.String() {
		case "ctrl+q":
			return selection, tea.Quit
		case "esc":
			return selection, func() tea.Msg {
				return view.ChangeWindow(selection.choices[selection.selected], view.Window(view.Selection))
			}
		case "q":
		case "up", "w", "k":
			if selection.selected > 0 {
				selection.selected -= 1
			}
		case "down", "s", "j":
			if selection.selected < len(selection.choices)-1 {
				selection.selected += 1
			}
		case "enter", " ":
			return selection, func() tea.Msg {
				return view.ChangeWindow(selection.choices[selection.selected], view.Window(view.FileRendering))
			}
		}
	}
	return selection, nil
}

func (selection SelectionModel) View() string {
	s := "what journal would you like to edit?\n"
	for i, choice := range selection.choices {
		cursor := " "
		if selection.selected == i {
			cursor = ">"
		}
		s += cursor + choice + "\n"
	}
	return s
}
