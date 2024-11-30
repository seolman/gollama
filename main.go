package main

import (
  "fmt"
  "log"
  "os"

  tea "github.com/charmbracelet/bubbletea"
)

type Model struct {
  choices []string
  cursor int
  selected map[int]bool
}

func initModel() Model {
  return Model{
    choices: []string{"Option 1", "Option 2", "Option 3"},
    selected: make(map[int]bool),
  }
}

func (m Model) Init() tea.Cmd {
  return nil
}


func (m Model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
  switch msg := msg.(type) {
  case tea.KeyMsg:
    switch msg.String() {
    case "ctrl+c", "q":
      return m, tea.Quit
    case "up", "k":
      if m.cursor > 0 {
        m.cursor--
      }
    case "down", "j":
      if m.cursor < len(m.choices) - 1 {
        m.cursor++
      }
    case " ":
      _, ok := m.selected[m.cursor]
      m.selected[m.cursor] = !ok
    }
  }
  return m, nil
}

func (m Model) View() string {
  s := "Enter Question\n\n"

  for i, choice := range m.choices {
    cursor := "  "
    if m.cursor == i {
      cursor = "> "
    }
    checked := "  "
    if _, ok := m.selected[i]; ok {
      checked = "x "
    }

    s += fmt.Sprintf("%s [%s] %s\n", cursor, checked, choice)
  }

  s += "\nPress q to quit.\n"
  return s
}


func main() {
  p := tea.NewProgram(initModel())
  if _, err := p.Run(); err != nil {
    log.Fatalf("Error: %v", err)
    os.Exit(1)
  }
}
