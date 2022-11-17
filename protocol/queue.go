package protocol

import (
	"context"
	"screenShot-service/entity"
)

type RetryQueue interface {
	Push(ctx context.Context, task *entity.Task) error
	Pop(ctx context.Context) (*entity.Task, error)
}
