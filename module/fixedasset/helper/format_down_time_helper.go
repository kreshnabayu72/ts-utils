package helper

import (
	"fmt"
	"time"
)

func FormatDownTime(seconds float64) string {
	sec := int(seconds)
	d := time.Duration(sec) * time.Second

	days := d / (24 * time.Hour)
	hours := (d % (24 * time.Hour)) / time.Hour
	minutes := (d % time.Hour) / time.Minute
	secondsLeft := (d % time.Minute) / time.Second

	if days > 0 {
		return fmt.Sprintf("%dd %dh", days, hours)
	} else if hours > 0 {
		return fmt.Sprintf("%dh %dm", hours, minutes)
	} else {
		return fmt.Sprintf("%dm %ds", minutes, secondsLeft)
	}
}
