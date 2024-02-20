package common

import "time"

var (
	uptime time.Time
)

func StartTime() {
	uptime = time.Now()
}

func GetUptime() time.Time {
	return uptime
}
