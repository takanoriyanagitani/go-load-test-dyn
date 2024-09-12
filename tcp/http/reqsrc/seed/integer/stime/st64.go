package stime

import (
	"context"
	"time"

	iseed "github.com/takanoriyanagitani/go-load-test-dyn/v2/tcp/http/reqsrc/seed/integer"
)

type TimeSource func(context.Context) (time.Time, error)

func (t TimeSource) ToIntegerSource64(t2i TimeToInteger64i) iseed.SeedSource64i {
	return func(ctx context.Context) (int64, error) {
		tm, e := t(ctx)
		if nil != e {
			return 0, e
		}

		var i int64 = t2i(tm)

		return i, nil
	}
}

var TimeSourceDefault TimeSource = func(_ context.Context) (time.Time, error) {
	return time.Now(), nil
}

type TimeToInteger64i func(time.Time) int64

var TimeToUnixtimeMicros64i TimeToInteger64i = func(t time.Time) int64 {
	return t.UnixMicro()
}

var SeedSourceUnixtimeMicros64i iseed.SeedSource64i = TimeSourceDefault.
	ToIntegerSource64(TimeToUnixtimeMicros64i)
