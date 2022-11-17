package entity

type status int

const (
	StatusCreated status = iota
	StatusProcesing
	StatusDone
	StatusInRetry
	StatusFailed

	MaxRetryCount = 3
)

type Task struct {
	Id         string                `json:"id"`
	Url        string                `json:"url"`
	Status     status                `json:"status"`
	RetryCount int                   `json:"retry_count"`
	Errors     [MaxRetryCount]string `json:"errors"`
}

func NewTask(url string) *Task {
	id := generateID()
	return &Task{
		Id:     id,
		Url:    url,
		Status: StatusCreated,
	}
}

func generateID() string {
	panic("unimplemented")
}
