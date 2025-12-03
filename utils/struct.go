package utils

import (
	"errors"
	"reflect"
)

func ValidateNotNilStructPtr(value interface{}) error {
	if value == nil {
		return errors.New("参数不能为空")
	}
	rv := reflect.ValueOf(value)
	if rv.Kind() != reflect.Ptr || rv.IsNil() {
		return errors.New("参数必须为非空指针")
	}
	return nil
}
