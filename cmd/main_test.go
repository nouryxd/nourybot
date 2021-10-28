package main

import "testing"

func TestMain(t *testing.T) {
	t.Run("main package test placeholder", func(t *testing.T) {
		request := 1
		response := 1

		got := request
		want := response

		if got != want {
			t.Errorf("got %q, want %q", got, want)
		}

	})
}
