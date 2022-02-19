package ivr

import "testing"

func TestUserid(t *testing.T) {
	t.Run("returns a link to a users profile picture", func(t *testing.T) {
		request := Userid("whereismymnd")
		response := "31437432"

		got := request
		want := response

		if got != want {
			t.Errorf("got %q, want %q", got, want)
		}

	})
}
