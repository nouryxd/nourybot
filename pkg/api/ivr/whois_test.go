package ivr

import "testing"

func TestWhois(t *testing.T) {
	t.Run("returns a link to a users profile picture", func(t *testing.T) {
		request := Whois("forsen")
		response := "User: forsen, ID: 22484632, Created on: 2011-05-19, Color: #FF0000, Affiliate: false, Partner: true, Staff: false, Admin: false, Bot: false, Bio: Approach with caution! No roleplaying or tryharding allowed."
		got := request
		want := response

		if got != want {
			t.Errorf("got %q, want %q", got, want)
		}

	})
}
