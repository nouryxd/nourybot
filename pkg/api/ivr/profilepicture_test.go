package ivr

import "testing"

func TestProfilePicture(t *testing.T) {
	t.Run("returns a link to a users profile picture", func(t *testing.T) {
		request, _ := ProfilePicture("forsen")
		response := "https://static-cdn.jtvnw.net/jtv_user_pictures/forsen-profile_image-48b43e1e4f54b5c8-600x600.png"

		got := request
		want := response

		if got != want {
			t.Errorf("got %q, want %q", got, want)
		}

	})
}
