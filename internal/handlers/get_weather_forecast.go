package handlers

import (
	"errors"
	"weatherbot/api/restapi/operations"
	"weatherbot/internal/usecase"

	"github.com/go-openapi/runtime/middleware"
)



type WeatherForecastHandler struct {
	usecase usecase.UseCase
}

func NewWeatherForecastHandler(usecase usecase.UseCase) *WeatherForecastHandler {
	return &WeatherForecastHandler{
		usecase: usecase,
	}
}

func(w *WeatherForecastHandler) Handle(params operations.GetWeatherParams) middleware.Responder {
	city, units := params.City, *params.Units

	if city == "" || units == "" {
		return operations.NewGetWeatherBadRequest().WithPayload(&operations.GetWeatherBadRequestBody{
			Error: errors.New("invalid request").Error(),
			Message: "Неправильные данные запроса",
		})
	}

	resp, err := w.usecase.GetForeCast(city, units)
	if err != nil {
		return operations.NewGetWeatherInternalServerError().WithPayload(&operations.GetWeatherInternalServerErrorBody{
			Error: err.Error(),
			Message: "Внутрення ошибка сервера",
		})
	}

	return operations.NewGetWeatherOK().WithPayload(&operations.GetWeatherOKBody{
		City: resp.Location.Name,
		FeelsLike: resp.Current.FeelsLikeC,
		Humidity: float64(resp.Current.Humidity),
		Temperature: resp.Current.TempC,
		WindSpeed: resp.Current.WindKph,
	})
}