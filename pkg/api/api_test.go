package api

import "testing"

// Idk how to test this really since almost everything in here is random
func TestApi(t *testing.T) {
	t.Run("api package test placeholder", func(t *testing.T) {
		request := 1
		response := 1

		got := request
		want := response

		if got != want {
			t.Errorf("got %q, want %q", got, want)
		}

	})
}