package main

import (
	tea "github.com/charmbracelet/bubbletea"
)

type model struct {
	todos    []string
	cursor   int
	selected map[int]struct{} // which to-do items are s
}

func initialModel() model {
	return model{
		todos: []string{"Continue building typing game"},
	}
}

func (m model) Init() tea.Cmd {
	return nil
}

func main() {
}
