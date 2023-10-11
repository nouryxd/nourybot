package owm

import (
	"errors"
	"fmt"
	"os"

	owm "github.com/briandowns/openweathermap"
	"github.com/joho/godotenv"
)

var (
	ErrInternalServerError     = errors.New("internal server error")
	ErrWeatherLocationNotFound = errors.New("location not found")
)

// Weather queries the OpenWeatherMap Api for the given location and sends the
// current weather response to the target twitch chat.
func Weather(location string) (string, error) {
	err := godotenv.Load()
	if err != nil {
		return "", ErrInternalServerError
	}
	owmKey := os.Getenv("OWM_KEY")

	w, err := owm.NewCurrent("C", "en", owmKey)
	if err != nil {
		return "", ErrInternalServerError
	}

	if err := w.CurrentByName(location); err != nil {
		return "", ErrInternalServerError
	}

	// Longitude and Latitude are returned as 0 when the supplied location couldn't be
	// assigned to a OpenWeatherMap location.
	if w.GeoPos.Longitude == 0 && w.GeoPos.Latitude == 0 {
		return "", ErrWeatherLocationNotFound
	} else {
		// Weather for Vilnius, LT: Feels like: 29.67°C. Currently 29.49°C with a high of 29.84°C and a low of 29.49°C, humidity: 45%, wind: 6.17m/s.
		reply := fmt.Sprintf("Weather for %s, %s: Feels like: %v°C. Currently %v°C with a high of %v°C and a low of %v°C, humidity: %v%%, wind: %vm/s.",
			w.Name,
			w.Sys.Country,
			w.Main.FeelsLike,
			w.Main.Temp,
			w.Main.TempMax,
			w.Main.TempMin,
			w.Main.Humidity,
			w.Wind.Speed,
		)
		return reply, nil
	}
}
