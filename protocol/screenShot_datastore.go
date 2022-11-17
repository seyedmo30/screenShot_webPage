package protocol

import (
	"context"
	"screenShot-service/entity"
)

type ScreenShotDataStore interface {
	CheckUrlExist(ctx context.Context, Url string) (bool, error)
	GetScreenShotByUrl(ctx context.Context, url string) (entity.ScreenShot, error)
	GetScreenShotById(ctx context.Context, id uint64) (entity.ScreenShot, error)
	CreateScreenShot(ctx context.Context, url string) (entity.ScreenShot, error)
	// QueueScreenShot(ctx context.Context, url string) error
}
