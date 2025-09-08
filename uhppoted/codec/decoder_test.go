package codec

import (
	"fmt"
	"time"
)

func string2datetime(v string) time.Time {
	if d, err := time.ParseInLocation("2006-01-02 15:04:05", v, time.Local); err != nil {
		panic(fmt.Sprintf("invalid datetime (%v)", v))
	} else {
		return d
	}
}

func string2time(v string) time.Time {
	if d, err := time.ParseInLocation("15:04:05", v, time.Local); err != nil {
		panic(fmt.Sprintf("invalid time (%v)", v))
	} else {
		return d
	}
}

func string2HHmm(v string) time.Time {
	if d, err := time.ParseInLocation("15:04", v, time.Local); err != nil {
		panic(fmt.Sprintf("invalid time (%v)", v))
	} else {
		return d
	}
}
