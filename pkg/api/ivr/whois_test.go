package ivr

import "testing"

func TestWhois(t *testing.T) {
	t.Run("returns a link to a users profile picture", func(t *testing.T) {
		request := Whois("Nouryxd")
		response := "User: Nouryxd, ID: 31437432, Created on: 2012-06-18, Color: #00F2FB, Affiliate: false, Partner: false, Staff: false, Admin: false, Bot: false, Bio: It's good to be king. Wait, maybe. I think maybe I'm just like a little bizarre little person who walks back and forth. Whatever, you know, but..."

		got := request
		want := response

		if got != want {
			t.Errorf("got %q, want %q", got, want)
		}

	})
}
