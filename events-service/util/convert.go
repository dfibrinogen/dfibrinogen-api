package util

import "time"

func ConvertStringToTime(origin string) time.Time {
	t, err := time.Parse(time.RFC3339, origin)
	if err != nil {
		return time.Now()
	}
	return t
}
