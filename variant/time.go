package variant

import (
	"database/sql/driver"
	"strings"
	"time"
)

type JsonTime time.Time

func (j JsonTime) MarshalJSON() ([]byte, error) {
	formatted := time.Time(j).Format(time.DateTime)
	return []byte(`"` + formatted + `"`), nil
}

func (j JsonTime) Value() (driver.Value, error) {
	t := time.Time(j)
	if t.IsZero() {
		return nil, nil
	}
	return t, nil // 直接返回 time.Time，驱动会处理
}

func (j *JsonTime) UnmarshalJSON(b []byte) error {
	str := strings.Trim(string(b), "\"")

	if str == "" {
		*j = JsonTime(time.Time{})
		return nil
	}

	t, err := time.ParseInLocation(time.DateTime, str, time.Local)
	if err != nil {
		return err
	}

	*j = JsonTime(t)
	return nil
}
