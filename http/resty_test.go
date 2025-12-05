package http

import (
	"testing"

	"resty.dev/v3"
)

func Test_Resty(t *testing.T) {
	RegisterResty(true)
	body, _ := RestyRaw(func(request *resty.Request) (*resty.Response, error) {
		return request.Get("https://www.baidu.com")
	})
	t.Log(body)
}
