package beerproto_go

import (
	"time"

	beerprotov1 "github.com/beerproto/beerproto_go/v1"
)

// Seconds returns the duration as a floating point number of seconds.
func Seconds(t *beerprotov1.TimeType) int64 {
	switch t.Unit {
	case beerprotov1.TimeUnit_TIME_UNIT_SEC:
		return t.Value
	case beerprotov1.TimeUnit_TIME_UNIT_MIN:
		return t.Value * 60
	case beerprotov1.TimeUnit_TIME_UNIT_HR:
		return t.Value * 60 * 60
	case beerprotov1.TimeUnit_TIME_UNIT_DAY:
		return t.Value * 60 * 60 * 24
	case beerprotov1.TimeUnit_TIME_UNIT_WEEK:
		return t.Value * 60 * 60 * 24 * 7
	}

	return 0
}

// Milliseconds returns the duration as an integer millisecond count.
func Milliseconds(t *beerprotov1.TimeType) int64 {
	return Seconds(t) * int64(time.Millisecond)
}
