package handlers

import "testing"

// Idk how to test this really we don't return anything
func TestHandlers(t *testing.T) {
	t.Run("handlers package test placeholder", func(t *testing.T) {
		request := 1
		response := 1

		got := request
		want := response

		if got != want {
			t.Errorf("got %q, want %q", got, want)
		}

	})
}
