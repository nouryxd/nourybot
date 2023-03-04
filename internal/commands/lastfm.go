package commands

import (
	"fmt"
	"os"

	"github.com/gempir/go-twitch-irc/v3"
	"github.com/joho/godotenv"
	"github.com/lyx0/nourybot/internal/common"
	"github.com/shkh/lastfm-go/lastfm"
	"go.uber.org/zap"
)

func LastFmArtistTop(target string, message twitch.PrivateMessage, tc *twitch.Client) {
	sugar := zap.NewExample().Sugar()
	defer sugar.Sync()
	// snipLength is the length we need to "snip" off of the start
	// of `message` to only have the artists name left.
	//  `()lastfm artist top` = +20
	//  trailing space =  +1
	//      zero-based =  +1
	//                 =  22
	snipLength := 20

	artist := message.Message[snipLength:len(message.Message)]

	err := godotenv.Load()
	if err != nil {
		sugar.Error("Error loading OpenWeatherMap API key from .env file")
	}
	apiKey := os.Getenv("LAST_FM_API_KEY")
	apiSecret := os.Getenv("LAST_FM_SECRET")

	api := lastfm.New(apiKey, apiSecret)
	result, _ := api.Artist.GetTopTracks(lastfm.P{"artist": artist}) //discarding error
	for _, track := range result.Tracks {
		sugar.Infow("Top tracks: ",
			"artist:", artist,
			"track", track.Name,
		)
	}
}

func LastFmUserRecent(target, user string, tc *twitch.Client) {
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
			common.Send(target, reply, tc)
			return
		}
	}

}
