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

func (ts *TaskService) CompleteTask(
	id string,
) (bool, error) {
	t, err := ts.repo.FindById(id)

	if err != nil {
		return false, err
	}

	if t == nil {
		return false, nil
	}

	t.Status = task.StatusComplete
	err = ts.repo.Update(*t)

	if err != nil {
		return false, err
	}

	return true, nil

}

func (ts *TaskService) DeleteTask(
	id string,
) error {
	err := ts.repo.Delete(id)

	if err != nil {
		return err
	}

	return nil
}

func (ts *TaskService) ListTasks() ([]task.Task, error) {
	tasks, err := ts.repo.FindAll()
	if err != nil {
		return nil, err
	}

	return tasks, nil

}

func (ts *TaskService) UpdateTask(t task.Task) error {
	return ts.repo.Update(t)

}
