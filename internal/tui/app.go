package tui

import (
	"github.com/BrandonDedolph/task-manager/internal/service"
	"github.com/BrandonDedolph/task-manager/internal/task"
	tea "github.com/charmbracelet/bubbletea"
)

type ViewModel struct {
	activeView tea.Model
	service    *service.TaskService
}

type switchToEditMsg struct {
	task *task.Task
}

type switchToFormMsg struct{}

type switchToConfirmMsg struct {
	task *task.Task
}

func NewViewModel(v tea.Model, s *service.TaskService) *ViewModel {
	return &ViewModel{activeView: v, service: s}
}

func (vm ViewModel) Init() tea.Cmd {
	return nil
}

func (vm ViewModel) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	switch msg := msg.(type) {
	case switchToListMsg:
		listModel, err := NewModel(vm.service)
		if err != nil {
			return vm, tea.Quit
		}
		vm.activeView = listModel
		return vm, nil

	case switchToEditMsg:
		vm.activeView = NewEditFormModel(vm.service, msg.task)
		return vm, nil

	case switchToFormMsg:
		vm.activeView = NewFormModel(vm.service)
		return vm, nil

	case switchToConfirmMsg:
		vm.activeView = NewConfirmModel(vm.service, msg.task)
		return vm, nil

	}
	updatedView, cmd := vm.activeView.Update(msg)
	vm.activeView = updatedView
	return vm, cmd
}

func (vm ViewModel) View() string {
	return vm.activeView.View()
}
