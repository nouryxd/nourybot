package aiden

import (
	"testing"
)

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
