package humanize

import (
	"time"

	"github.com/dustin/go-humanize"
)

// Time returns a more human readable
// time as a string for a given time.Time
func Time(t time.Time) string {
	return humanize.Time(t)
}
