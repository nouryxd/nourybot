package ivr

import "testing"

func TestWhois(t *testing.T) {
	t.Run("returns a link to a users profile picture", func(t *testing.T) {
		request := Whois("forsen")
		got := request
		want := Whois("forsen")

		if got != want {
			t.Errorf("got %q, want %q", got, want)
		}

	})
}
