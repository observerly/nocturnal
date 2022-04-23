package utils

import "time"

const RFC3339 = "2006-01-02T15:04:05.000Z"

func ParseDatetimeRFC3339(d string) (time.Time, error) {
	return time.Parse(RFC3339, d)
}
