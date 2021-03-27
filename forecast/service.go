package forecast

import (
	"fmt"

	"github.com/thetaurean/weatherman/weather"
)

type WeatherProvider interface {
	GetWeatherByZip(zip string) (weather.Weather, error)
}

type service struct {
	weatherProvider WeatherProvider
}

func NewService(p WeatherProvider) *service {
	return &service{
		weatherProvider: p,
	}
}

func (s *service) LocalForecast(zip string) (weather.Weather, error) {
	w, err := s.weatherProvider.GetWeatherByZip(zip)
	if err != nil {
		return weather.Weather{}, fmt.Errorf("LocalForecast: %w", err)
	}

	return w, nil
}
