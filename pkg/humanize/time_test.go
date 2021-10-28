package humanize

import (
	"fmt"
	"testing"
	"time"
)

func TestTime(t *testing.T) {
	t.Run("tests the humanized time", func(t *testing.T) {

		now := time.Now()

		fmt.Println("now:", now)

		count := 10
		// then is the current time - 10 minutes.
		then := now.Add(time.Duration(-count) * time.Minute)

		request := Time(then)
		response := "10 minutes ago"

		got := request
		want := response

		if got != want {
			t.Errorf("got %v, want %v", got, want)
		}
	})
}
