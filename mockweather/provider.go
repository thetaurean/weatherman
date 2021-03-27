package mockweather

import (
	"github.com/thetaurean/weatherman/weather"
	"time"
)

type provider struct {
	apiKey string
}

func NewProvider(apiKey string) *provider {
	return &provider{
		apiKey: apiKey,
	}
}

type mockWeatherResponse struct {
	Message string
	Main    struct {
		Temp     float32 `json:"temp"`
		Humidity float32 `json:"humidity"`
	}
	Wind struct {
		Speed float32 `json:"speed"`
	}
	Created time.Time
}

func (r mockWeatherResponse) ToWeather() weather.Weather {
	return weather.Weather{
		Temp:      r.Main.Temp,
		Humidity:  r.Main.Humidity,
		Windspeed: r.Wind.Speed,
		Created:   r.Created,
	}
}

func (p *provider) GetWeatherByZip(zip string) (weather.Weather, error) {
	weatherResult := mockWeatherResponse{
		Message: "This is a mock weather response",
		Main: struct {
			Temp     float32 `json:"temp"`
			Humidity float32 `json:"humidity"`
		}{
			Temp:     72.0,
			Humidity: 46.0,
		},
		Wind: struct {
			Speed float32 `json:"speed"`
		}{
			Speed: 5.0,
		},
		Created: time.Now(),
	}

	return weatherResult.ToWeather(), nil
}
