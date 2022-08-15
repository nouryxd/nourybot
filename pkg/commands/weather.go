package commands

import (
	"fmt"
	"os"

	owm "github.com/briandowns/openweathermap"
	"github.com/gempir/go-twitch-irc/v3"
	"github.com/joho/godotenv"
	"github.com/lyx0/nourybot/pkg/common"
	"go.uber.org/zap"
)

// Weather queries the OpenWeatherMap Api for the given location and sends the
// current weather response to the target twitch chat.
func Weather(target, location string, tc *twitch.Client) {
	sugar := zap.NewExample().Sugar()
	defer sugar.Sync()

	err := godotenv.Load()
	if err != nil {
		sugar.Error("Error loading OpenWeatherMap API key from .env file")
	}
	owmKey := os.Getenv("OWM_KEY")

	w, err := owm.NewCurrent("C", "en", owmKey)
	if err != nil {
		sugar.Error(err)
	}
	w.CurrentByName(location)

	// Longitude and Latitude are returned as 0 when the supplied location couldn't be
	// assigned to a OpenWeatherMap location.
	if w.GeoPos.Longitude == 0 && w.GeoPos.Latitude == 0 {
		reply := "Location not found FeelsBadMan"
		common.Send(target, reply, tc)
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
		common.Send(target, reply, tc)
	}
}
