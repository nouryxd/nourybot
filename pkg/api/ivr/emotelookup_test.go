package ivr

import "testing"

func TestEmoteLookup(t *testing.T) {
	t.Run("returns the channel a twitch emote is from and its tier", func(t *testing.T) {
		// Tier 1
		request, _ := EmoteLookup("forsenE")
		response := "forsenE is a Tier 1 emote to channel forsen."

		got := request
		want := response

		if got != want {
			t.Errorf("got %q, want %q", got, want)
		}

		// Tier 2
		request, _ = EmoteLookup("nanZ")
		response = "nanZ is a Tier 2 emote to channel nani."

		got = request
		want = response

		if got != want {
			t.Errorf("got %q, want %q", got, want)
		}

		// Tier 3
		request, _ = EmoteLookup("pajaCORAL")
		response = "pajaCORAL is a Tier 3 emote to channel pajlada."

		got = request
		want = response

		if got != want {
			t.Errorf("got %q, want %q", got, want)
		}

	})
}
