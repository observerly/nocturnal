package utils

import (
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

func TestParseDatetimeRFC3339(t *testing.T) {
	d := "2006-01-02T15:04:05.000Z"

	datetime, _ := ParseDatetimeRFC3339(d)

	// Assert we get the correctly parse datetime:
	assert.Equal(t, datetime.Year(), 2006)
	assert.Equal(t, datetime.Month(), time.Month(1))
	assert.Equal(t, datetime.Day(), 2)
	assert.Equal(t, datetime.Hour(), 15)
	assert.Equal(t, datetime.Minute(), 4)
	assert.Equal(t, datetime.Second(), 5)
}

func TestParseDatetimeRFC3339Alt(t *testing.T) {
	d := "2021-05-14T00:00:00.000Z"

	datetime, _ := ParseDatetimeRFC3339(d)

	// Assert we get the correctly parse datetime:
	assert.Equal(t, datetime.Year(), 2021)
	assert.Equal(t, datetime.Month(), time.Month(5))
	assert.Equal(t, datetime.Day(), 14)
	assert.Equal(t, datetime.Hour(), 0)
	assert.Equal(t, datetime.Minute(), 0)
	assert.Equal(t, datetime.Second(), 0)
}

func TestFormatDatetimeRFC3339(t *testing.T) {
	d := time.Date(2006, 1, 2, 15, 4, 5, 0, time.UTC)

	datetime := FormatDatetimeRFC3339(&d)

	// Assert we get the correctly parse datetime:
	assert.Equal(t, *datetime, "2006-01-02T15:04:05Z")
}

func TestFormatDatetimeRFC3339Nil(t *testing.T) {
	var d *time.Time

	datetime := FormatDatetimeRFC3339(d)

	// Assert we get the correctly parse datetime:

	assert.Equal(t, datetime, (*string)(nil))
}
