package cacheweather

import (
	"github.com/thetaurean/weatherman/forecast"
	"github.com/thetaurean/weatherman/weather"
	"time"

	"github.com/patrickmn/go-cache"
)

type provider struct {
	provider forecast.WeatherProvider
	cache    *cache.Cache
}

func NewProvider(wp forecast.WeatherProvider) *provider {
	return &provider{
		provider: wp,
		cache:    cache.New(time.Hour, time.Hour),
	}
}

func (p *provider) GetWeatherByZip(zip string) (weather.Weather, error) {
	cached, found := p.cache.Get(zip)

	if found {
		return cached.(weather.Weather), nil
	}

	result, err := p.provider.GetWeatherByZip(zip)
	p.cache.Set(zip, result, cache.DefaultExpiration)

	return result, err
}
