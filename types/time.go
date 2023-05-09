package types

import "time"

type Time time.Time

func (t Time) String() string {
	return time.Time(t).In(time.Local).Format("2006-01-02 15:04:05")
}

func (t Time) MarshalJSON() ([]byte, error) {
	return []byte(`"` + time.Time(t).In(time.Local).Format("2006-01-02 15:04:05") + `"`), nil
}
