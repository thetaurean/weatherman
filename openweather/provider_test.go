package openweather

import (
	"os"
	"testing"
)

func TestOWProvider(t *testing.T) {
	apiKey := os.Getenv("OPENWEATHER_KEY")
	provider := NewProvider(apiKey)
	weather, err := provider.GetWeatherByZip("21044")

	if err != nil {
		t.Log(err)
		t.Fail()
	} else {
		t.Logf("Temperature: %f\nHumidity: %f\nWindspeed: %f\nAge: %s", weather.Temp,
			weather.Humidity,
			weather.Windspeed,
			weather.Age())
	}
}
