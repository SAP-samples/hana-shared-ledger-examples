package times

import (
	"fmt"
	t "time"
)

// This layout is compatible to HANA datetime
var timeLayout = "2006-01-02 15:04:05.0000000"

func ToString(time t.Time) string {
	return time.Format(timeLayout)
}

func ParseString(time string) (t.Time, error) {
	parsedTime, err := t.ParseInLocation(timeLayout, time, t.UTC)
	if err != nil {
		return t.Time{}, fmt.Errorf("not a valid time for layout %v: %w", timeLayout, err)
	}

	return parsedTime, nil
}

func TimeIsInIntervall(time t.Time, deviation t.Duration) bool {
	now := t.Now().UTC()

	if time.Before(now.Add(-deviation)) {
		return false
	}

	if time.After(now.Add(deviation)) {
		return false
	}

	return true
}
