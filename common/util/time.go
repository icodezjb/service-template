package util

import "time"

func CalRequestTime(t int64) int64 {
	return (time.Now().UnixNano() - t) / 1000000
}
