package variant

import (
	"strings"
	"time"
)

type JsonTime time.Time

func (j JsonTime) MarshalJSON() ([]byte, error) {
	formatted := time.Time(j).Format(time.DateTime)
	return []byte(`"` + formatted + `"`), nil
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
