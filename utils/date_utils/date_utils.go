package date

import "time"

const (
	apiDateLayout = "2020-06-21T15:04:052"
)

// GetCurrentTime -
func GetCurrentTime() time.Time {
	return time.Now().UTC()
}

// GetCurrentTimeString -
func GetCurrentTimeString() {
	GetCurrentTime().Format(apiDateLayout)
}