package ivr

import "testing"

func TestFollowage(t *testing.T) {
	t.Run("returns a users date they followed a channel", func(t *testing.T) {
		request, _ := Followage("forsen", "pajlada")
		response := "pajlada has been following forsen since 2015-03-09."

		got := request
		want := response

		if got != want {
			t.Errorf("got %q, want %q", got, want)
		}

	})
}
