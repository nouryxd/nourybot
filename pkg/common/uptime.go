package common

import "time"

var (
	uptime time.Time
)

// StartTime adds the current time to the uptime variable.
func StartTime() {
	uptime = time.Now()
}

// GetUptime returns the time from the uptime variable.
func GetUptime() time.Time {
	return uptime
}
