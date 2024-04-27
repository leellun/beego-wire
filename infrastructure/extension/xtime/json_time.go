package xtime

import (
	"database/sql/driver"
	"fmt"
	"time"
)

const (
	TimeLayout = "2006-01-02 15:04:05"
)

// JsonTime format xjson time field by myself
type JsonTime struct {
	time.Time
}

var TimeTemplate1 = "2006-01-02 15:04:05"

func TimeInt(t1 string) int64 {
	stamp, _ := time.ParseInLocation(TimeTemplate1, t1, time.Local)
	return stamp.Unix()
}
func Zero() JsonTime {
	var t JsonTime
	return t
}
func Now() JsonTime {
	var t JsonTime
	t.Time = time.Now()
	return t
}
func NowTime() JsonTime {
	var t JsonTime
	t.Time = time.Now()
	return t
}
func Date(year int, month time.Month, day int, hour int, min int, sec int, nsec int, loc *time.Location) JsonTime {
	var tt JsonTime
	tt.Time = time.Date(year, month, day, hour, min, sec, nsec, loc)
	return tt
}
func (t JsonTime) Before(u JsonTime) bool {
	return t.Time.Before(u.Time)
}
func (t JsonTime) IsZero() bool {
	return t.Time.IsZero()
}
func (t JsonTime) Add(d time.Duration) JsonTime {
	var tt JsonTime
	tt.Time = t.Time.Add(d)
	return tt
}
func (t JsonTime) AddDate(years int, months int, days int) JsonTime {
	var tt JsonTime
	tt.Time = t.Time.AddDate(years, months, days)
	return tt
}
func (t *JsonTime) UnmarshalJSON(data []byte) (err error) {
	if len(data) == 2 {
		*t = JsonTime{Time: time.Time{}}
		return
	}
	strValue := string(data)
	if strValue == "null" {
		*t = JsonTime{Time: time.Time{}}
		return
	}
	now, err := time.Parse(`"`+TimeLayout+`"`, strValue)
	*t = JsonTime{Time: now}
	return
}

// MarshalJSON on JsonTime format Time field with Y-m-d H:i:s
func (t JsonTime) MarshalJSON() ([]byte, error) {
	if t.Time.IsZero() {
		return []byte("null"), nil
	}
	formatted := fmt.Sprintf("\"%s\"", t.Format(TimeLayout))
	return []byte(formatted), nil
}

// Value insert timestamp into mysql need this function.
func (t JsonTime) Value() (driver.Value, error) {
	var zeroTime time.Time
	if t.Time.UnixNano() == zeroTime.UnixNano() {
		return nil, nil
	}
	return t.Time, nil
}

// Scan value of time.Time
func (t *JsonTime) Scan(v interface{}) error {
	value, ok := v.(time.Time)
	if ok {
		*t = JsonTime{Time: value}
		return nil
	}
	return fmt.Errorf("can not convert %v to timestamp", v)
}
