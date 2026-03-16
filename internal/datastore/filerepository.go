package datastore

import (
	"encoding/json"
	"errors"
	"fmt"
	"os"
	"slices"

	"github.com/BrandonDedolph/task-manager/internal/task"
)

type FileRepository struct {
	filePath string
}

func NewFileRepository(filePath string) *FileRepository {
	return &FileRepository{filePath: filePath}
}

func (r *FileRepository) FindAll() ([]task.Task, error) {
	fileData, err := os.ReadFile(r.filePath)
	if errors.Is(err, os.ErrNotExist) {
		return []task.Task{}, nil
	} else if err != nil {
		return nil, err
	}

	var tasks []task.Task

	err = json.Unmarshal(fileData, &tasks)

	if err != nil {
		return nil, err
	}

	return tasks, nil

}

func (r *FileRepository) Write(t []task.Task) error {
	bytes, err := json.Marshal(t)
	if err != nil {
		return err
	}

	err = os.WriteFile(r.filePath, bytes, 0644)
	if err != nil {
		return err
	}

	return nil
}

func (r *FileRepository) Create(t task.Task) error {
	retrievedTasks, err := r.FindAll()
	if err != nil {
		return err
	}

	allTasks := append(retrievedTasks, t)

	return r.Write(allTasks)

}

func FindTaskById(t_slice []task.Task, id string) *task.Task {
	taskIndex := slices.IndexFunc(t_slice, func(t task.Task) bool {
		return t.ID == id
	})

	if taskIndex != -1 {
		return &t_slice[taskIndex]
	}

	return nil
}

func (r *FileRepository) Update(t task.Task) error {
	retrievedTasks, err := r.FindAll()
	if err != nil {
		return err
	}

	existingTask := FindTaskById(retrievedTasks, t.ID)

	if existingTask != nil {
		*existingTask = t
	} else {
		return fmt.Errorf("Task with id %s didn't exist, update failed", t.ID)
	}

	return r.Write(retrievedTasks)
}

func (r *FileRepository) FindById(id string) (*task.Task, error) {
	retrievedTasks, err := r.FindAll()
	if err != nil {
		return nil, err
	}

	t := FindTaskById(retrievedTasks, id)

	if t != nil {
		return t, nil
	}

	return nil, nil

}

func (r *FileRepository) Delete(id string) error {
	retrievedTasks, err := r.FindAll()
	if err != nil {
		return err
	}

	tasksAfterDeletion := slices.DeleteFunc(retrievedTasks, func(t task.Task) bool {
		return t.ID == id
	})

	return r.Write(tasksAfterDeletion)
}
