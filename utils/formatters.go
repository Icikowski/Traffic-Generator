package utils

import (
	"fmt"
	"time"
)

// FormatDuration returns human-readable representation of given time.Duration
func FormatDuration(duration time.Duration) string {
	rounded := duration.Round(time.Second)
	minutes := rounded / time.Minute
	seconds := (rounded - (minutes * time.Minute)) / time.Second
	return fmt.Sprintf("%dm%02ds", minutes, seconds)
}
