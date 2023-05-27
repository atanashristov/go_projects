package util

import "time"

func DurationInSeconds(seconds int64) time.Duration {
	return time.Duration(seconds) * time.Second
}
