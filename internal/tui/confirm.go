package tui

import (
	"fmt"

	"github.com/BrandonDedolph/task-manager/internal/service"
	"github.com/BrandonDedolph/task-manager/internal/task"
	tea "github.com/charmbracelet/bubbletea"
)

type confirmModel struct {
	task    *task.Task
	service *service.TaskService
}

func NewConfirmModel(s *service.TaskService, t *task.Task) confirmModel {
	return confirmModel{task: t, service: s}
}

func (cm confirmModel) Init() tea.Cmd {
	return nil
}

func (cm confirmModel) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	switch msg := msg.(type) {
	case tea.KeyMsg:
		switch msg.String() {
		case "y", "enter":
			err := cm.service.DeleteTask(cm.task.ID)
			if err != nil {
				return cm, nil
			} else {
				return cm, func() tea.Msg { return switchToListMsg{} }
			}
		case "n", "esc":
			return cm, func() tea.Msg { return switchToListMsg{} }
		}
	}
	return cm, nil
}

func (cm confirmModel) View() string {
	return fmt.Sprintf("Are you sure you want to delete '%s'? (y/n)", cm.task.Name)
}
