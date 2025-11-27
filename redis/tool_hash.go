package redis

import "context"

type toolsHash interface {
	HSet(ctx context.Context, key string, values ...interface{}) error
	HGet(ctx context.Context, key, field string) (string, error)
	HGetAll(ctx context.Context, key string) (map[string]string, error)
	HDel(ctx context.Context, key string, fields ...string) error
	HExists(ctx context.Context, key, field string) (bool, error)
}

// HSet 设置字段
func (t *tool) HSet(ctx context.Context, key string, values ...interface{}) error {
	return t.client.HSet(ctx, key, values).Err()
}

// HGet 获取字段
func (t *tool) HGet(ctx context.Context, key, field string) (string, error) {
	return t.client.HGet(ctx, key, field).Result()
}

// HGetAll 获取所有字段
func (t *tool) HGetAll(ctx context.Context, key string) (map[string]string, error) {
	return t.client.HGetAll(ctx, key).Result()
}

// HDel 删除一个或多个字段
func (t *tool) HDel(ctx context.Context, key string, fields ...string) error {
	return t.client.HDel(ctx, key, fields...).Err()
}

// HExists 判断字段是否存在
func (t *tool) HExists(ctx context.Context, key, field string) (bool, error) {
	return t.client.HExists(ctx, key, field).Result()
}
