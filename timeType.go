package beerprotov1

import (
	"fmt"
	"time"
)

// Seconds returns the duration as a floating point number of seconds.
func (t *TimeType) Seconds() int64 {
	switch t.Unit {
	case TimeUnit_TIME_UNIT_SEC:
		return t.Value
	case TimeUnit_TIME_UNIT_MIN:
		return t.Value * 60
	case TimeUnit_TIME_UNIT_HR:
		return t.Value * 60 * 60
	case TimeUnit_TIME_UNIT_DAY:
		return t.Value * 60 * 60 * 24
	case TimeUnit_TIME_UNIT_WEEK:
		return t.Value * 60 * 60 * 24 * 7
	}

	return 0
}

// Milliseconds returns the duration as an integer millisecond count.
func (t *TimeType) Milliseconds() int64 {
	return t.Seconds() * int64(time.Millisecond)
}

func (x *TimeType) DisplayName() string {
	if x == nil {
		return ""
	}
	return fmt.Sprintf("%v %v", x.Value, x.Unit.DisplayName())
}
