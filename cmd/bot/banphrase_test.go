package bot

import "testing"

func TestBanphrase(t *testing.T) {
	t.Run("tests the banphrase api", func(t *testing.T) {
		// Banned message
		request, _ := CheckMessage("https://gyazo.com/asda")
		response := true

		got := request
		want := response

		if got != want {
			t.Errorf("got %v, want %v", got, want)
		}

		// Okay message
		request, _ = CheckMessage("LUL")
		response = false

		got = request
		want = response

		if got != want {
			t.Errorf("got %v, want %v", got, want)
		}

	})
}
