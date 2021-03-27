# Weatherman

Weatherman is a Go module for retrieving current weather data.

## Installation

Use the go import system to import the package

```go
import (
	"github.com/thetaurean/weatherman/cacheweather"
	"github.com/thetaurean/weatherman/openweather"
)
```

## Usage

#### Be sure to set environment variable OPENWEATHER_KEY to your api key.

```go
apiKey := os.Getenv("OPENWEATHER_KEY")
provider := openweather.NewProvider(apiKey)
cachedProvider := cacheweather.NewProvider(provider)

weather,err := cachedProvider.GetWeatherByZip("21044")
```

#### See `main.go` for an extended example on usage:

## Architecture

Weatherman uses a provider design pattern to enable modularity with respect to the implementation
used for retrieving the current weather. As such, the various pieces are described below:

Data structures / interfaces
- `weather` this package contains the target data structure and properties for retrieving weather data
- `forecast` this is a generalized interface for interacting with weather providers (see below)

Weather providers
- `mockweather` this is a mock implementation returning static weather data for testing purposes
- `openweather` this is a real implementation for retrieving weather data from openweathermaps API
- `cacheweather` this is a nesting provider which caches results from a contained weather provider


## Considerations / TODO

- I am not using the forecast service
  - forecast should also probably include a helper or two for instantiation
- I am instantiating cache / rate limiter inside weather providers for simplicity
  - previously this was done outside providers (i.e. main)  and passed byref to the weather providers
- Better error checking / handling
- Better documentation

