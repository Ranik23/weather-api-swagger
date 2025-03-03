package usecase

import (
	"encoding/json"
	"fmt"
	"net/http"
	"weatherbot/config"
	"weatherbot/internal/models"
)


type UseCase interface {
	GetForeCast(city string, units string) (*models.WeatherResponse, error)
}

type UseCaseImpl struct {
	client 	*http.Client
	config 	*config.Config
}

func NewUseCaseImpl( config *config.Config) *UseCaseImpl {
	return &UseCaseImpl{
		client: http.DefaultClient,
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