package validate

import (
	"errors"
	"fmt"
	"reflect"

	"github.com/go-playground/validator/v10"
)

func ParseError(err error, obj interface{}) string {
	var res string

	objType := reflect.TypeOf(obj)
	// 如果是指针则取 Elem()
	if objType.Kind() == reflect.Ptr {
		objType = objType.Elem()
	}

	var errs validator.ValidationErrors
	if errors.As(err, &errs) {
		for _, e := range errs {
			field, _ := objType.FieldByName(e.Field())
			name := field.Name
			if lbl := field.Tag.Get("label"); lbl != "" {
				name = lbl
			}
			tag, ok := validatorTagMap[e.Tag()]
			if !ok || tag == "" {
				res = fmt.Sprintf("%s不合法,校验失败", name)
			} else {
				res = fmt.Sprintf(tag, name, e.Param())
			}
		}
	}

	return res
}
