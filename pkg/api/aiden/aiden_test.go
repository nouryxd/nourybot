package aiden

import (
	"testing"
)

// TestAidenBotRatelimits calls aidens bot ratelimits api and checks
// the results against each other with my bot and another verified bot.
func TestAidenBotRatelimits(t *testing.T) {
	t.Run("returns a twitch accounts ratelimits", func(t *testing.T) {
		requestPajbot, _ := ApiCall("api/v1/twitch/botStatus/pajbot?includeLimits=1")
		requestNourybot, _ := ApiCall("api/v1/twitch/botStatus/nourybot?includeLimits=1")

		got := requestPajbot
		want := requestNourybot

		if got != want {
			t.Errorf("got %q, want %q", got, want)
		}
	})
}
