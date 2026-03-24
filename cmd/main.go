package main

import (
	"github.com/BrandonDedolph/task-manager/internal/datastore"
	"github.com/BrandonDedolph/task-manager/internal/service"
	"github.com/BrandonDedolph/task-manager/internal/tui"
	tea "github.com/charmbracelet/bubbletea"
)

func main() {
	repo := datastore.NewFileRepository("tasks.json")
	svc := service.NewTaskService(repo)
	app, err := tui.NewViewModel(svc)
	if err != nil {
		return
	}

	tea.NewProgram(app).Run()
}
