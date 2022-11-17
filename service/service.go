package service

import (
	"context"
	"errors"
	"screenShot-service/dto"
	"screenShot-service/entity"
	"screenShot-service/protocol"
)

type service struct {
	store  protocol.ScreenShotDataStore
	queue  protocol.RetryQueue
	tasker chan *entity.Task
}

var _ protocol.ScreenShotInteractor = &service{}

func NewController(store protocol.ScreenShotDataStore, queue protocol.RetryQueue) *service {
	return &service{
		store:  store,
		queue:  queue,
		tasker: make(chan *entity.Task, 20),
	}
}

func (s *service) TakeScreenShot(ctx context.Context, req dto.TakeScreenShotRequest) (dto.TakeScreenShotResponse, error) {
	//validate

	exist, err := s.store.CheckUrlExist(ctx, req.Url)

	if err != nil {
		return dto.TakeScreenShotResponse{}, errors.New("CheckUrlExist 500")
	}
	if exist {

		return dto.TakeScreenShotResponse{}, errors.New("url exist")
	}

	task := entity.NewTask(req.Url)
	go func() {
		s.tasker <- task
	}()
	return dto.TakeScreenShotResponse{ID: task.Id, Message: "sabr kon"}, nil
}

func (s service) GetScreenShotDetailById(ctx context.Context, req dto.GetScreenShotDetailByIdRequest) (dto.GetScreenShotDetailResponse, error) {
	//validate
	screenShot, err := s.store.GetScreenShotById(ctx, req.Id)
	if err != nil {
		return dto.GetScreenShotDetailResponse{}, errors.New("GetScreenShotById 500")
	}

	return dto.GetScreenShotDetailResponse{ScreenShot: screenShot}, nil
}

// Tasker implements protocol.ScreenShotInteractor
func (s *service) Tasker() chan *entity.Task {
	return s.tasker
}
