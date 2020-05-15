package beerproto

import "time"

// Seconds returns the duration as a floating point number of seconds.
func (t *TimeType) Seconds() int64 {
	switch m.Unit {
	case TimeType_SEC:
		return m.Value
	case TimeType_MIN:
		return int64(m.Value * 60)
	case TimeType_HR:
		return int64(m.Value * 60 * 60)
	case TimeType_DAY:
		return int64(m.Value * 60 * 60 * 24)
	case TimeType_WEEK:
		return int64(m.Value * 60 * 60 * 24 * 7)
	}

	return 0
}

// Milliseconds returns the duration as an integer millisecond count.
func (t *TimeType) Milliseconds() int64 {
	return m.Seconds() * int64(time.Millisecond)
}
