package util

import (
	"database/sql/driver"
	"fmt"
	"time"
)

type Datetime string

func (d *Datetime) Scan(src interface{}) error {
	if src == nil {
		*d = TimeToDatetime(time.Now())
		return nil
	}
	switch src := src.(type) {
	case string:
		*d = Datetime(src)
		return nil
	case time.Time:
		*d = TimeToDatetime(src)
		return nil
	}

	return fmt.Errorf("cannot scan %T", src)
}

func (d Datetime) Value() (driver.Value, error) {
	tz, _ := time.LoadLocation("Asia/Seoul")
	return time.ParseInLocation("2006-01-02 15:04:05", string(d), tz)
}

func NowDatetime() Datetime {
	return TimeToDatetime(time.Now())
}

// TimeToDatetime converts `time` to `yyyy-MM-dd hh:mm:ss` formatted string
func TimeToDatetime(time time.Time) Datetime {
	return Datetime(time.Format("2006-01-02 15:04:05"))
}
