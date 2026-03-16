package service

import (
	"github.com/BrandonDedolph/task-manager/internal/task"
	"github.com/google/uuid"
)

type TaskService struct {
	repo task.Repository
}

func NewTaskService(tr task.Repository) *TaskService {
	return &TaskService{repo: tr}
}

func (ts *TaskService) CreateTask(
	name string,
	description string,
	status task.Status,
	priority int,
) (*task.Task, error) {
	t := task.Task{
		ID:          uuid.New().String(),
		Name:        name,
		Description: description,
		Status:      status,
		Priority:    priority}

	err := ts.repo.Create(t)
	if err != nil {
		return nil, err
	}
	return &t, nil
}
