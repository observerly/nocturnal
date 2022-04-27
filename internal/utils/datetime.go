package utils

import "time"

func ParseDatetimeRFC3339(d string) (time.Time, error) {
	return time.Parse(time.RFC3339, d)
}
