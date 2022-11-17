package worker

import (
	"context"
	"screenShot-service/entity"
)

type worker struct {
	tasker       chan *entity.Task
	workersCount int
}

func New() *worker {
	return &worker{workersCount: 10}
}

func (w *worker) SetTasker(t chan *entity.Task) {
	w.tasker = t
}

func (w *worker) RunForEverInshaallah(ctx context.Context) error {

	for i := 0; i < w.workersCount; i++ {
		go func() {
			for {
				select {
				case task := <-w.tasker:

				}
			}
		}()
	}

}
