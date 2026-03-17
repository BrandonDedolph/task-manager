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
	editing *task.Task
	service *service.TaskService
}

func NewFormModel(s *service.TaskService) formModel {
	input := textinput.New()
	input.Focus()
	input.Placeholder = "Task name..."

	return formModel{input: input, service: s}
}

func NewEditFormModel(s *service.TaskService, t *task.Task) formModel {
	input := textinput.New()
	input.Focus()
	input.SetValue(t.Name)

	return formModel{input: input, editing: t, service: s}
}

func (fm formModel) Init() tea.Cmd {
	return nil
}

func (fm formModel) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	switch msg := msg.(type) {
	case tea.KeyMsg:
		switch msg.String() {
		case "enter":
			if fm.editing == nil {
				_, err := fm.service.CreateTask(
					fm.input.Value(),
					"",
					task.StatusTODO,
					0,
				)
				if err != nil {
					return fm, nil
				}
			} else {
				fm.editing.Name = fm.input.Value()
				err := fm.service.UpdateTask(*fm.editing)
				if err != nil {
					return fm, nil
				}
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
	title := "Add Task"
	if fm.editing != nil {
		title = "Edit Task"
	}
	return title + "\n\n" + fm.input.View()
}

