package handlers

import (
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

	resp, err := w.usecase.GetForeCast(city, units)
	if err != nil {
		return operations.NewGetWeatherInternalServerError()
	}

	return operations.NewGetWeatherOK().WithPayload(&operations.GetWeatherOKBody{
		City: resp.Location.Name,
		FeelsLike: resp.Current.FeelsLikeC,
		Humidity: float64(resp.Current.Humidity),
		Temperature: resp.Current.TempC,
		WindSpeed: resp.Current.WindKph,
	})
}