package task

type Status string

const (
	StatusTODO       Status = "todo"
	StatusInProgress Status = "inprogress"
	StatusComplete   Status = "complete"
	StatusUndefined  Status = ""
)

type Task struct {
	ID          string
	Name        string
	Description string
	Status      Status
	Priority    int
}

func (t Task) Validate() bool {
	switch t.Status {
	case StatusTODO, StatusInProgress, StatusComplete, StatusUndefined:
		return true
	}
	return false
}
