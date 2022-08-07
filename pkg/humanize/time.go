package humanize

import (
	"time"

	"github.com/dustin/go-humanize"
)

func Time(t time.Time) string {
	return humanize.Time(t)
}
