package tui

import (
	"fmt"
	"github.com/BrandonDedolph/task-manager/internal/service"
	"github.com/BrandonDedolph/task-manager/internal/task"
	tea "github.com/charmbracelet/bubbletea"
	"strings"
)

type model struct {
	tasks   []task.Task
	cursor  int
	service *service.TaskService
}

func NewModel(s *service.TaskService) (model, error) {
	tasks, err := s.ListTasks()
	if err != nil {
		return model{}, err
	}
	return model{tasks: tasks, service: s}, nil
}

func (m model) Init() tea.Cmd {
	return nil
}

func (m model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	switch msg := msg.(type) {
	case tea.KeyMsg:
		switch msg.String() {
		case "q", "ctrl+c":
			return m, tea.Quit
		case "up", "k":
			if m.cursor > 0 {
				m.cursor--
			}
		case "down", "j":
			if m.cursor < len(m.tasks)-1 {
				m.cursor++
			}
		case "e":
			return m, func() tea.Msg {
				return switchToEditMsg{task: &m.tasks[m.cursor]}
			}
		case "d":
			return m, func() tea.Msg {
				return switchToConfirmMsg{task: &m.tasks[m.cursor]}
			}
		case "a":
			return m, func() tea.Msg {
				return switchToFormMsg{}
			}
		case "space", "enter":
			m.service.CompleteTask(m.tasks[m.cursor].ID)
			tasks, err := m.service.ListTasks()
			if err == nil {
				m.tasks = tasks
			}
		}
	}
	return m, nil
}

func (m model) View() string {
	header := "Task Manager\n\n"

	tasks := m.tasks

	if tasks == nil {
		return ""
	}

	var builder strings.Builder

	for index, t := range tasks {
		prepend := " "
		if index == m.cursor {
			prepend = ">"
		}
		fmt.Fprintf(&builder, "%s%s [%s]\n", prepend, t.Name, t.Status)
	}

	footer := "q: quit"

	return header + builder.String() + footer
}
