package lastfm

import (
	"fmt"
	"os"

	"github.com/joho/godotenv"
	"github.com/shkh/lastfm-go/lastfm"
	"go.uber.org/zap"
)

func LastFmUserRecent(target, user string) string {
	sugar := zap.NewExample().Sugar()
	defer sugar.Sync()

	err := godotenv.Load()
	if err != nil {
		sugar.Error("Error loading LASTFM API keys from .env file")
	}

	apiKey := os.Getenv("LAST_FM_API_KEY")
	apiSecret := os.Getenv("LAST_FM_SECRET")

	api := lastfm.New(apiKey, apiSecret)
	result, _ := api.User.GetRecentTracks(lastfm.P{"user": user}) //discarding error

	var reply string
	for i, track := range result.Tracks {
		// The 0th result is the most recent one since it goes from most recent
		// to least recent.
		if i == 0 {
			sugar.Infow("Most recent: ",
				"user:", user,
				"track", track.Name,
				"artist", track.Artist.Name,
			)

			reply = fmt.Sprintf("Most recently played track for user %v: %v - %v", user, track.Artist.Name, track.Name)
		}
	}

	return reply
}
