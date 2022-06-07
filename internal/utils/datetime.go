package utils

import "time"

func ParseDatetimeRFC3339(d string) (time.Time, error) {
	return time.Parse(time.RFC3339, d)
}

func FormatDatetimeRFC3339(d *time.Time) *string {
	if d != nil {
		var f = d.Format(time.RFC3339)
		return &f
	}

	return nil
}
