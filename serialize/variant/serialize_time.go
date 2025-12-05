package variant

import (
	"database/sql/driver"
	"strings"
	"time"
)

type SerializeTime time.Time

func (s SerializeTime) MarshalJSON() ([]byte, error) {
	formatted := time.Time(s).Format(time.DateTime)
	return []byte(`"` + formatted + `"`), nil
}

func (s *SerializeTime) UnmarshalJSON(b []byte) error {
	str := strings.Trim(string(b), "\"")

	if str == "" {
		*s = SerializeTime(time.Time{})
		return nil
	}

	t, err := time.ParseInLocation(time.DateTime, str, time.Local)
	if err != nil {
		return err
	}

	*s = SerializeTime(t)
	return nil
}

func (s SerializeTime) Value() (driver.Value, error) {
	t := time.Time(s)
	if t.IsZero() {
		return nil, nil
	}
	return t, nil // 直接返回 time.Time，驱动会处理
}
