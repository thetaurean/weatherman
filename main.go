package main

import (
	"log"
	"os"

	"github.com/thetaurean/weatherman/cacheweather"
	"github.com/thetaurean/weatherman/openweather"
)

func main() {
	apiKey := os.Getenv("OPENWEATHER_KEY")
	provider := openweather.NewProvider(apiKey)
	cachedProvider := cacheweather.NewProvider(provider)

	w1, err := cachedProvider.GetWeatherByZip("21044")
	if err == nil {
		log.Printf("Temperature: %f\nHumidity: %f\nWindspeed: %f\nAge: %s", w1.Temp,
			w1.Humidity,
			w1.Windspeed,
			w1.Age())
	} else {
		log.Fatalf("Error occurred: %v", err.Error())
	}
}
