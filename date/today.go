package date

import (
	"time"
)

func ParseNanoToTime(n int64) time.Time {
	tm := time.Unix(n/1000, 0)
	return tm
}
