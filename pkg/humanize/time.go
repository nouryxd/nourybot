package humanize

import (
	"time"

	"github.com/dustin/go-humanize"
)

// HumanizeTime returns a more human readable
// time as a string for a given time.Time
func HumanizeTime(t time.Time) string {
	return humanize.Time(t)
}
