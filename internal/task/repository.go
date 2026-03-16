package task

type Repository interface {
	Create(task Task) error
	Update(task Task) error
	FindAll() ([]Task, error)
	FindById(id string) (Task, error)
	Delete(id string) error
}
