package protocol

import (
	"context"
	"screenShot-service/dto"
	"screenShot-service/entity"
)

type ScreenShotInteractor interface {
	TakeScreenShot(ctx context.Context, req dto.TakeScreenShotRequest) (dto.TakeScreenShotResponse, error)
	Tasker() chan *entity.Task
}
