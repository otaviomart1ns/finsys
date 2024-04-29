package utils

import "time"

func Coalesce(newVal, oldVal string) string {
	if newVal != "" {
		return newVal
	}
	return oldVal
}

func CoalesceTime(newVal, oldVal time.Time) time.Time {
	if !newVal.IsZero() {
		return newVal
	}
	return oldVal
}
