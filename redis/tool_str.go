package redis

import (
	"context"
	"errors"
	"time"

	"github.com/nilchaosky/gear/utils"
	"github.com/redis/go-redis/v9"
)

type toolString interface {
	Cache(ctx context.Context, key string, value interface{}, exp time.Duration, fn func() error) error
	Get(ctx context.Context, key string) (string, error)
	GetToStruct(ctx context.Context, key string, value interface{}) error
	Set(ctx context.Context, key string, value interface{}) error
	SetEx(ctx context.Context, key string, value interface{}, exp time.Duration) error
	SetNX(ctx context.Context, key string, value interface{}) error
	SetNEX(ctx context.Context, key string, value interface{}, exp time.Duration) error
	SetXX(ctx context.Context, key string, value interface{}) error
	SetXEX(ctx context.Context, key string, value interface{}, exp time.Duration) error
}

func (t *tool) Cache(ctx context.Context, key string, value interface{}, exp time.Duration, fn func() error) error {
	if err := utils.ValidateNotNilStructPtr(value); err != nil {
		return err
	}
	err := t.client.Get(ctx, key).Scan(value)
	if err == nil {
		return nil // hit
	}

	if !errors.Is(err, redis.Nil) {
		return err
	}

	if err := fn(); err != nil {
		return err
	}

	return t.SetEx(ctx, key, value, exp)
}

// Get 获取数据
func (t *tool) Get(ctx context.Context, key string) (string, error) {
	return t.client.Get(ctx, key).Result()
}

// GetToStruct 获取数据并转换
func (t *tool) GetToStruct(ctx context.Context, key string, value interface{}) error {
	if err := utils.ValidateNotNilStructPtr(value); err != nil {
		return err
	}
	return t.client.Get(ctx, key).Scan(value)
}

// Set 设置 key
func (t *tool) Set(ctx context.Context, key string, value interface{}) error {
	return t.client.Set(ctx, key, value, 0).Err()
}

// SetEx 设置 key 和 过期时间
func (t *tool) SetEx(ctx context.Context, key string, value interface{}, exp time.Duration) error {
	return t.client.Set(ctx, key, value, exp).Err()
}

// SetNX 如果 key 不存在才设置
func (t *tool) SetNX(ctx context.Context, key string, value interface{}) error {
	return t.client.SetNX(ctx, key, value, 0).Err()
}

// SetNEX 如果 key 不存在才设置 和 过期时间
func (t *tool) SetNEX(ctx context.Context, key string, value interface{}, exp time.Duration) error {
	return t.client.SetNX(ctx, key, value, exp).Err()
}

// SetXX 如果 key 存在才设置
func (t *tool) SetXX(ctx context.Context, key string, value interface{}) error {
	return t.client.SetXX(ctx, key, value, 0).Err()
}

// SetXEX 	如果 key 存在才设置 和 过期时间
func (t *tool) SetXEX(ctx context.Context, key string, value interface{}, exp time.Duration) error {
	return t.client.SetXX(ctx, key, value, exp).Err()
}
