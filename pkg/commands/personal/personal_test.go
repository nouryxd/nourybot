package personal

import "testing"

func TestCommands(t *testing.T) {
	t.Run("personal package test placeholder", func(t *testing.T) {
		request := 1
		response := 1

		got := request
		want := response

		if got != want {
			t.Errorf("got %q, want %q", got, want)
		}

	})
}