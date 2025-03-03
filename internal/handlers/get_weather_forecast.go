package handlers

import (
	"errors"
	"time"
	"weatherbot/api/restapi/operations"
	"weatherbot/internal/usecase"

	"github.com/go-openapi/runtime/middleware"
	"github.com/prometheus/client_golang/prometheus"
)



type WeatherForecastHandler struct {
	usecase 	usecase.UseCase
	duration 	*prometheus.HistogramVec
	counter 	*prometheus.CounterVec
}

func NewWeatherForecastHandler(usecase usecase.UseCase, duration *prometheus.HistogramVec,
								counter *prometheus.CounterVec) *WeatherForecastHandler {
	return &WeatherForecastHandler{
		usecase: usecase,
		duration: duration,
		counter: counter,
	}
}

func(w *WeatherForecastHandler) Handle(params operations.GetWeatherParams) middleware.Responder {


	start := time.Now()

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


	w.duration.WithLabelValues(params.HTTPRequest.Method, params.HTTPRequest.URL.Path).Observe(time.Since(start).Seconds())
	w.counter.WithLabelValues(params.HTTPRequest.Method, params.HTTPRequest.URL.Path).Inc()


	return operations.NewGetWeatherOK().WithPayload(&operations.GetWeatherOKBody{
		City: resp.Location.Name,
		FeelsLike: resp.Current.FeelsLikeC,
		Humidity: float64(resp.Current.Humidity),
		Temperature: resp.Current.TempC,
		WindSpeed: resp.Current.WindKph,
	})
}