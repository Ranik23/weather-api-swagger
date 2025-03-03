package handlers

import (
	operations "weatherbot/api/restapi/operations"
	"weatherbot/internal/usecase"

	"github.com/go-openapi/runtime/middleware"
)



type WeatherForecast7DaysHandler struct {
	usecase usecase.UseCase
}

func NewWeatherForecast7DaysHandler(usecase usecase.UseCase) *WeatherForecast7DaysHandler {
	return &WeatherForecast7DaysHandler{
		usecase: usecase,
	}
}


func (w *WeatherForecast7DaysHandler) Handle(params operations.GetWeatherForecast7daysParams) middleware.Responder {
	return operations.NewGetWeatherForecast7daysOK().WithPayload(nil)
}