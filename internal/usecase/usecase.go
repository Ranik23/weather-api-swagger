package usecase

import (
	"encoding/json"
	"fmt"
	"net/http"
	"weatherbot/config"
	"weatherbot/internal/storage"
	"weatherbot/internal/models"
)


type UseCase interface {
	GetForeCast(city string, units string) (*models.WeatherResponse, error)
	GetForeCast7Days(city string, units string) ([]models.WeatherResponse7Days, error)
	GetForeCast30Days(city string, units string) ([]models.WeatherResponse30Days, error)
}

type UseCaseImpl struct {
	client 	*http.Client
	strg 	storage.Storage
	config 	*config.Config
}

func NewUseCaseImpl(strg storage.Storage, config *config.Config) *UseCaseImpl {
	return &UseCaseImpl{
		client: http.DefaultClient,
		strg: strg,
		config: config,
	}
}

func (u *UseCaseImpl) GetForeCast(city string, units string) (*models.WeatherResponse, error) {
	url := fmt.Sprintf("http://api.weatherapi.com/v1/current.json?key=%s&q=%s&aqi=no", u.config.API_KEY, city)

	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return nil, err
	}

	resp, err := u.client.Do(req)
	if err != nil {
		return nil, err
	}

	defer resp.Body.Close()

	var weatherResponse models.WeatherResponse
	if err := json.NewDecoder(resp.Body).Decode(&weatherResponse); err != nil {
		return nil, err
	}

	return &weatherResponse, nil
}

func (u *UseCaseImpl) GetForeCast7Days(city string, units string) ([]models.WeatherResponse7Days, error) {
	url := fmt.Sprintf("http://api.weatherapi.com/v1/forecast.json?key=%s&q=%s&days=7&aqi=no", u.config.API_KEY, city)

	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return nil, err
	}

	resp, err := u.client.Do(req)
	if err != nil {
		return nil, err
	}

	defer resp.Body.Close()

	var weatherResponse []models.WeatherResponse7Days
	if err := json.NewDecoder(resp.Body).Decode(&weatherResponse); err != nil {
		return nil, err
	}

	// Возвращаем прогноз на 7 дней
	return weatherResponse, nil
}

func (u *UseCaseImpl) GetForeCast30Days(city string, units string) ([]models.WeatherResponse30Days, error) {
	url := fmt.Sprintf("http://api.weatherapi.com/v1/forecast.json?key=%s&q=%s&days=30&aqi=no", u.config.API_KEY, city)

	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return nil, err
	}

	resp, err := u.client.Do(req)
	if err != nil {
		return nil, err
	}

	defer resp.Body.Close()

	var weatherResponse []models.WeatherResponse30Days
	if err := json.NewDecoder(resp.Body).Decode(&weatherResponse); err != nil {
		return nil, err
	}

	// Возвращаем прогноз на 30 дней
	return weatherResponse, nil
}
