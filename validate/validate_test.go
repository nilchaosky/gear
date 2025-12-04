package validate

import (
	"testing"

	"github.com/go-playground/validator/v10"
)

type Interior struct {
	Title string `json:"title" label:"标题" validate:"required,min=2,max=20"`
}

type Outer struct {
	List []Interior `json:"list" label:"标题列表" validate:"required,min=1,dive"`
}

var validate = validator.New()

func Test(t *testing.T) {
	req := Outer{
		List: []Interior{
			{Title: ""},     // 不合法：required
			{Title: "a"},    // 不合法：min=2
			{Title: "测试OK"}, // 合法
		},
	}
	err := validate.Struct(&req)
	if err == nil {
		t.Fatal("expected validation error, got nil")
	}

	msg := FieldParseError(err, req)
	if msg == "" {
		t.Fatal("expected error message, got empty string")
	}

	t.Log(msg)
}
