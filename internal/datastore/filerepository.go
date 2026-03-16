package datastore

import (
	"encoding/json"
	"errors"
	"github.com/BrandonDedolph/task-manager/internal/task"
	"os"
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
