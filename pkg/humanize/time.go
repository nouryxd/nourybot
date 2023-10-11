package humanize

import (
	"time"

	humanize "github.com/dustin/go-humanize"
)

func Time(t time.Time) string {
	return humanize.Time(t)
}
