package cacheweather

import (
	"testing"
	"time"

	"github.com/patrickmn/go-cache"
	"github.com/thetaurean/weatherman/mockweather"
)

func TestCachedMockProvider(t *testing.T) {
	provider := mockweather.NewProvider("apiKey")
	cache := cache.New(time.Hour, time.Hour)
	cachedProvider := NewProvider(provider, cache)

	w1, _ := cachedProvider.GetWeatherByZip("21044")
	t.Logf("Weather 1\nTemperature: %f\nHumidity: %f\nWindspeed: %f\nAge: %s", w1.Temp,
		w1.Humidity,
		w1.Windspeed,
		w1.Age())

	w2, _ := cachedProvider.GetWeatherByZip("21044")
	t.Logf("Weather 2\nTemperature: %f\nHumidity: %f\nWindspeed: %f\nAge: %s", w2.Temp,
		w2.Humidity,
		w2.Windspeed,
		w2.Age())
	if w1 != w2 {
		t.Fail()
	}

}
