package utils

import (
	"testing"
)

// I don't know how to test the other ones since they are random
func TestCommandsUsed(t *testing.T) {
	t.Run("tests the commands used counter", func(t *testing.T) {

		request := mockCommandsUsed(127)
		response := 127

		got := request
		want := response

		if got != want {
			t.Errorf("got %v, want %v", got, want)
		}

		// 127 + 53
		// request = mockCommandsUsed(53)
		// response = 180

		// got = request
		// want = response

		// if got != want {
		// 	t.Errorf("got %v, want %v", got, want)
		// }
	})
}

func mockCommandsUsed(n int) int {
	for i := 0; i < n; i++ {
		CommandUsed()
	}
	return tempCommands
}
