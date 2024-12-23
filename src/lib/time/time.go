package time

import (
	"fmt"
	"time"
)

var defaultLoc, _ = time.LoadLocation("Asia/Taipei")

func ParseTime(layout, dateStr string, loc *time.Location) time.Time {
	if loc == nil {
		loc = defaultLoc
	}

	parsedTime, err := time.ParseInLocation(layout, dateStr, loc)
	if err != nil {
		fmt.Printf("found an error: %v\n", err)
		return time.Time{}
	}

	return parsedTime
}
