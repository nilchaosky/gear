package validate

import (
	"errors"
	"fmt"
	"reflect"
	"strconv"
	"strings"

	"github.com/go-playground/validator/v10"
)

func FieldParseError(err error, req interface{}) string {
	if err == nil {
		return ""
	}

	var errs validator.ValidationErrors
	ok := errors.As(err, &errs)
	if !ok {
		return err.Error()
	}

	vt := reflect.TypeOf(req)
	if vt.Kind() == reflect.Ptr {
		vt = vt.Elem()
	}

	message := make([]string, 0)

	for _, e := range errs {

		// 解析字段链路（含数组下标）
		name := buildLabelNameByNamespace(e.Namespace(), vt)

		// 取错误模板
		tagTpl, ok := validatorTagMap[e.Tag()]
		var msg string

		// 模板不存在
		if !ok || tagTpl == "" {
			msg = fmt.Sprintf("%s不合法,校验失败", name)
		} else {
			if e.Param() == "" {
				msg = fmt.Sprintf(tagTpl, name)
			} else {
				msg = fmt.Sprintf(tagTpl, name, e.Param())
			}
		}

		message = append(message, msg)
	}
	return strings.Join(message, "; ")
}

func buildLabelNameByNamespace(namespace string, root reflect.Type) string {
	parts := strings.Split(namespace, ".")

	// 跳过最外层 struct 名称，如 Outer.List[0].Title → 从 List 开始
	parts = parts[1:]

	currentType := root
	var res strings.Builder

	for i, part := range parts {

		// 处理数组字段: List[0]
		if strings.Contains(part, "[") {
			fieldName := part[:strings.Index(part, "[")]
			indexStr := part[strings.Index(part, "[")+1 : strings.Index(part, "]")]
			index, _ := strconv.Atoi(indexStr)

			field, _ := currentType.FieldByName(fieldName)
			label := field.Tag.Get("label")
			if label == "" {
				label = fieldName
			}

			// 标题列表中的第1项
			res.WriteString(fmt.Sprintf("%s中的第%d项", label, index+1))

			// 更新到切片的元素类型
			currentType = field.Type.Elem()
		} else {
			// 普通字段
			field, _ := currentType.FieldByName(part)
			label := field.Tag.Get("label")
			if label == "" {
				label = part
			}

			res.WriteString(label)
		}

		if i != len(parts)-1 {
			res.WriteString("的")
		}
	}

	return res.String()
}
