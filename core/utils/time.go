package utils

import "time"

// StringToTime parse string data to time
func StringToTime(strTime string) time.Time {
	layout := "2006-01-02 15:04:05 -0700 -07"
	t, _ := time.Parse(layout, strTime)
	return t
}
