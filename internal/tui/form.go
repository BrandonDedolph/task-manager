package tui

import (
	"github.com/BrandonDedolph/task-manager/internal/service"
	"github.com/BrandonDedolph/task-manager/internal/task"
	"github.com/charmbracelet/bubbles/textinput"
	tea "github.com/charmbracelet/bubbletea"
)

type switchToListMsg struct{}

type formModel struct {
	input   textinput.Model
	service *service.TaskService
}

func NewFormModel(s *service.TaskService) formModel {
	input := textinput.New()
	input.Focus()
	input.Placeholder = "Task name..."

	return formModel{input: input, service: s}
}

func (fm formModel) Init() tea.Cmd {
	return nil
}

func (fm formModel) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	switch msg := msg.(type) {
	case tea.KeyMsg:
		switch msg.String() {
		case "enter":
			_, err := fm.service.CreateTask(
				fm.input.Value(),
				"",
				task.StatusTODO,
				0,
			)
			if err != nil {
				return fm, nil
			}
			return fm, func() tea.Msg { return switchToListMsg{} }
		case "esc":
			return fm, func() tea.Msg { return switchToListMsg{} }
		}
	}
	var cmd tea.Cmd
	fm.input, cmd = fm.input.Update(msg)
	return fm, cmd
}

func (fm formModel) View() string {
	return "Add Task\n\n" + fm.input.View()
}

