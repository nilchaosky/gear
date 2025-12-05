package variant

import (
	"database/sql/driver"
	"fmt"
	"strconv"
	"strings"
)

type SerializeInt64 int64

func (s SerializeInt64) MarshalJSON() ([]byte, error) {
	return []byte(fmt.Sprintf("\"%d\"", s)), nil
}

func (s *SerializeInt64) UnmarshalJSON(b []byte) error {
	str := strings.Trim(string(b), "\"")
	if str == "" || str == "null" {
		*s = 0
		return nil
	}
	n, err := strconv.ParseInt(str, 10, 64)
	if err != nil {
		return fmt.Errorf("invalid SerializeInt64: %w", err)
	}
	*s = SerializeInt64(n)
	return nil
}

func (s SerializeInt64) Value() (driver.Value, error) {
	return int64(s), nil
}

func (s SerializeInt64) String() string {
	return strconv.FormatInt(int64(s), 10)
}
