package handlers

import (
	"weatherbot/api/restapi/operations"
	"weatherbot/internal/usecase"

	"github.com/go-openapi/runtime/middleware"
)



type WeatherForecast30DaysHandler struct {
	usecase usecase.UseCase
}


func NewWeatherForecast30DaysHandler(usecase usecase.UseCase) *WeatherForecast30DaysHandler {
	return &WeatherForecast30DaysHandler{
		usecase: usecase,
	}
}

func (w *WeatherForecast30DaysHandler) Handle(params operations.GetWeatherForecast30daysParams) middleware.Responder {
	return operations.NewGetWeatherForecast30daysOK().WithPayload(nil)
}