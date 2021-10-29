package ivr

import "testing"

func TestWhois(t *testing.T) {
	t.Run("returns a link to a users profile picture", func(t *testing.T) {
		request := Whois("nouryqt")
		response := "User: Nouryqt, ID: 31437432, Created on: 2012-06-18, Color: #00F2FB, Affiliate: false, Partner: false, Staff: false, Admin: false, Bot: false, Bio: me :)"

		got := request
		want := response

		if got != want {
			t.Errorf("got %q, want %q", got, want)
		}

	})
}
