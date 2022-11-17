package dto

import "screenShot-service/entity"

type TakeScreenShotRequest struct {
	Url string `json:"url" validate:"required"`
}

type TakeScreenShotResponse struct {
	ID      string `json:"id"`
	Message string `json:"message"`
}

type GetScreenShotDetailByIdRequest struct {
	Id uint64 `json:"id" validate:"required"`
}

type GetScreenShotDetailByUrlRequest struct {
	Url string `json:"url" validate:"required"`
}

type GetScreenShotDetailResponse struct {
	ScreenShot entity.ScreenShot `json:"screenshot"`
}
