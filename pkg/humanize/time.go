package humanize

import (
	"time"

	humanize "github.com/dustin/go-humanize"
)

// Time returns the humanized time for a supplied time value.
func Time(t time.Time) string {
	return humanize.Time(t)
}
