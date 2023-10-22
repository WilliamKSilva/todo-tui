package main

import (
	"fmt"
	"os"

	tea "github.com/charmbracelet/bubbletea"
)

type model struct {
	todos    []string
	cursor   int
	selected map[int]struct{} // which to-do items are s
}

func initialModel() model {
	return model{
		todos: []string{"Continue building typing game", "iaushduiashdi"},
	}
}

func (m model) Init() tea.Cmd {
	return nil
}

func (m model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	switch msg := msg.(type) {
	case tea.KeyMsg:
		switch msg.String() {
		case "ctrl+c", "q":
			return m, tea.Quit
		case "up":
			if m.cursor > 0 {
				m.cursor = m.cursor - 1
			}
		case "down":
			if m.cursor < len(m.todos)-1 {
				m.cursor = m.cursor + 1
			}
		}
	}

	return m, nil
}

func (m model) View() string {
	s := "What important do you have to do?\n"

	for i, todo := range m.todos {
		cursor := " "

		if m.cursor == i {
			cursor = ">"
		}

		checked := " "
		if _, ok := m.selected[i]; ok {
			checked = "x"
		}

		s += fmt.Sprintf("%s [%s] %s\n", cursor, checked, todo)
	}

	s += "\nPress q to quit.\n"

	return s
}

func main() {
	program := tea.NewProgram(initialModel())

	_, err := program.Run()

	if err != nil {
		fmt.Printf("Error running the TODO app")
		os.Exit(1)
	}
}
