package main

import (
	"fmt"
	"log"

	tea "charm.land/bubbletea/v2"
)

type model struct {
	cursor int
	items  []string
}

func initialModel() model {
	return model{
		items: []string{
			"test1",
			"test2",
			"test3",
			"test4",
		},
	}
}

func (m model) Init() tea.Cmd {
	return nil
}

func (m model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	switch msg := msg.(type) {

	case tea.KeyPressMsg:
		switch msg.String() {

		case "q", "ctrl+c":
			return m, tea.Quit

		case "up", "k":
			if m.cursor > 0 {
				m.cursor--
			}

		case "down", "j":
			if m.cursor < len(m.items)-1 {
				m.cursor++
			}

		case "enter":
			fmt.Printf("You pushed: %s\n", m.items[m.cursor])
			return m, tea.Quit
		}
	}

	return m, nil
}

func (m model) View() tea.View {
	s := "\nCheck:\n\n"

	for i, item := range m.items {
		cursor := " "

		if i == m.cursor {
			cursor = ">"
		}

		s += fmt.Sprintf("%s %s\n", cursor, item)
	}

	s += "\n↑/↓ — change • Enter - choose • q — exit\n"

	return tea.NewView(s)
}

func main() {
	p := tea.NewProgram(initialModel())

	if _, err := p.Run(); err != nil {
		log.Fatal(err)
	}
}
