package openweather

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"time"

	"go.uber.org/ratelimit"

	"github.com/thetaurean/weatherman/weather"
)

const (
	endpoint               = "https://api.openweathermap.org/data/2.5"
	pathFormatWeatherByZip = "/weather?zip=%s&appid=%s&units=imperial"
)

type provider struct {
	apiKey  string
	limiter *ratelimit.Limiter
}

func NewProvider(apiKey string) *provider {
	l := ratelimit.New(1)
	return &provider{
		apiKey:  apiKey,
		limiter: &l,
	}
}

func (p *provider) GetWeatherByZip(zip string) (weather.Weather, error) {
	(*p.limiter).Take()

	// compose the request url
	path := fmt.Sprintf(pathFormatWeatherByZip, zip, p.apiKey)
	u := endpoint + path

	res, err := http.Get(u)
	if err != nil {
		return weather.Weather{}, fmt.Errorf("openweather.GetWeatherByZip failed http GET: %s", err)
	}

	defer func() {
		_ = res.Body.Close()
	}()

	// read the response
	bodyRaw, err := ioutil.ReadAll(res.Body)
	if err != nil {
		return weather.Weather{}, fmt.Errorf("openweather.GetWeatherByZip failed reading body: %s", err)
	}

	var weatherRes openWeatherResponse
	if err = json.Unmarshal(bodyRaw, &weatherRes); err != nil {
		return weather.Weather{}, fmt.Errorf("openweather.GetWeatherByZip failed encoding body: %s", err)
	}

	if res.StatusCode != http.StatusOK {
		return weather.Weather{}, fmt.Errorf("openweather.GetWeatherByZip got error from OpenWeather: %s", weatherRes.Message)
	}

	if len(weatherRes.Message) > 0 {
		return weather.Weather{}, fmt.Errorf("openweather.GetWeatherByZip got message from OpenWeather: %s", weatherRes.Message)
	}

	// return the external response converted into an entity
	return weatherRes.ToWeather(), nil
}

type openWeatherResponse struct {
	Message string
	Main    struct {
		Temp     float32 `json:"temp"`
		Humidity float32 `json:"humidity"`
	}
	Wind struct {
		Speed float32 `json:"speed"`
	}
	Created int64 `json:"dt"`
}

func (r openWeatherResponse) ToWeather() weather.Weather {
	return weather.Weather{
		Temp:      r.Main.Temp,
		Humidity:  r.Main.Humidity,
		Windspeed: r.Wind.Speed,
		Created:   time.Unix(r.Created, 0),
	}
}
