package beerproto_go

import "time"

// Seconds returns the duration as a floating point number of seconds.
func (t *TimeType) Seconds() int64 {
	switch m.Unit {
	case TimeType_SEC:
		return int64(m.Value)
	case TimeType_MIN:
		return int64(m.Value * 60)
	case TimeType_HR:
		return int64(m.Value * 60 * 60)
	case TimeType_DAY:
		return int64(m.Value * 60 * 60 * 24)
	case TimeType_WEEK:
		return int64(m.Value * 60 * 60 * 24 * 7)
	case TimeType_MONTH:
		return int64(m.Value * 60 * 60 * 24 * 30)
	case TimeType_YEAR:
		return int64(m.Value * 60 * 60 * 24 * 365)
	}

	return 0
}

// Milliseconds returns the duration as an integer millisecond count.
func (t *TimeType) Milliseconds() int64 {
	return m.Seconds() * int64(time.Millisecond)
}
