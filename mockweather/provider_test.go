package mockweather

import (
	"testing"
)

func TestMockProvider(t *testing.T) {
	provider := NewProvider("apiKey")
	weather, _ := provider.GetWeatherByZip("21044")

	t.Logf("Temperature: %f\nHumidity: %f\nWindspeed: %f\nAge: %s", weather.Temp,
		weather.Humidity,
		weather.Windspeed,
		weather.Age())
}
