package http

import (
	"errors"

	"github.com/nilchaosky/gear/logz"
	"go.uber.org/zap"
	"resty.dev/v3"
)

var Resty *resty.Client

func RegisterResty(debug bool) {
	client := resty.New()
	client.SetDebug(debug)
	Resty = client
}

func RestyRaw(do func(*resty.Request) (*resty.Response, error)) (string, error) {
	resp, err := do(Resty.R())
	if err != nil {
		logz.Print.Error("HTTP 请求错误", zap.Error(err))
		return "", err
	}
	if resp.IsError() {
		logz.Print.Error("HTTP 请求响应错误", zap.Int("status", resp.StatusCode()), zap.String("body", resp.String()))
		return "", errors.New("HTTP 请求响应错误")
	}
	return resp.String(), nil
}

func RestyDo[T interface{}](result T, e interface{}, do func(*resty.Request) (*resty.Response, error)) error {
	resp, err := do(Resty.R().SetResult(result).SetError(e))

	if err != nil {
		logz.Print.Error("HTTP 请求错误", zap.Error(err))
		return err
	}
	if resp.IsError() {
		logz.Print.Error("HTTP 请求响应错误", zap.Int("status", resp.StatusCode()), zap.String("body", resp.String()))
		return errors.New("HTTP 请求响应错误")
	}
	return nil
}
