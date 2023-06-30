package types

import (
	"bytes"
	"time"
)

type Time time.Time

func (t Time) String() string {
	return time.Time(t).In(time.Local).Format("2006-01-02 15:04:05")
}

func (t Time) MarshalJSON() ([]byte, error) {
	return []byte(`"` + time.Time(t).In(time.Local).Format("2006-01-02 15:04:05") + `"`), nil
}

func (t *Time) UnmarshalJSON(b []byte) (err error) {
	var tm time.Time
	if tm, err = time.ParseInLocation("2006-01-02 15:04:05", string(bytes.Trim(b, "\"")), time.Local); err != nil {
		return
	}
	*t = Time(tm)
	return nil
}
