package task

type Status string

const (
	StatusTODO       Status = "todo"
	StatusInProgress Status = "inprogress"
	StatusComplete   Status = "complete"
	StatusUndefined  Status = ""
)

type Task struct {
	ID          string `json:"id"`
	Name        string `json:"name"`
	Description string `json:"description"`
	Status      Status `json:"status"`
	Priority    int    `json:"priority"`
}

func (t Task) Validate() bool {
	switch t.Status {
	case StatusTODO, StatusInProgress, StatusComplete, StatusUndefined:
		return true
	}
	return false
}
