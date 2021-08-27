package utils

import (
	"fmt"
	"time"
)

// FormatDuration returns human-readable representation of given time.Duration
func FormatDuration(duration time.Duration) string {
	duration = duration.Round(time.Second)
	minutes := duration / time.Minute
	duration -= minutes
	seconds := duration / time.Second
	return fmt.Sprintf("%dm%02ds", minutes, seconds)
}
