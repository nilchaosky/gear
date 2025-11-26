package http

import (
	"errors"
	"log"

	"github.com/nilchaosky/gear/logz"
	"go.uber.org/zap"
	"resty.dev/v3"
)

var Client *resty.Client

func NewClient() *resty.Client {
	client := resty.New()
	client.SetDebug(true)
	Client = client
	return client
}

func RestyDo[T any](
	do func(*resty.Request) (*resty.Response, error),
	result *T, e any,
) (*T, error) {
	resp, err := do(Client.R().SetResult(result).SetError(e))

	if err != nil {
		log.Printf(": %v", err)
		logz.Print.Error("HTTP 请求错误", zap.Error(err))
		return nil, err
	}
	if resp.IsError() {
		logz.Print.Error("HTTP 请求响应错误", zap.Int("status", resp.StatusCode()), zap.String("body", resp.String()))
		return nil, errors.New("HTTP 请求响应错误")
	}
	return result, nil
}
