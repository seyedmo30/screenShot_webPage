package main

import (
	"screenShot-service/service"
	"screenShot-service/worker"

	"github.com/go-delve/delve/service"
)

// type CustomValidator struct {
// 	validator *validator.Validate
// }

// func (cv *CustomValidator) Validate(i interface{}) error {
// 	if err := cv.validator.Struct(i); err != nil {
// 		// Optionally, you could return the error to give each route more control over the status code
// 		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
// 	}
// 	return nil
// }

func main() {

	intractor := service.NewController(nil, nil)
	workerInteractor := worker.New()
	s := intractor.Tasker()
	workerInteractor.SetTasker(s)

}
