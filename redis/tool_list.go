package redis

import "context"

type toolList interface {
	LPush(ctx context.Context, key string, values ...interface{}) error
	RPush(ctx context.Context, key string, values ...interface{}) error
	LPop(ctx context.Context, key string) (string, error)
	LPopToStruct(ctx context.Context, key string, result interface{}) error
	RPop(ctx context.Context, key string) (string, error)
	RPopToStruct(ctx context.Context, key string, result interface{}) error
	LLen(ctx context.Context, key string) (int64, error)
	LRange(ctx context.Context, key string, start, stop int64) ([]string, error)
}

// LPush 列表左推入
func (t *tool) LPush(ctx context.Context, key string, values ...interface{}) error {
	return t.client.LPush(ctx, key, values).Err()
}

// RPush 列表右推入
func (t *tool) RPush(ctx context.Context, key string, values ...interface{}) error {
	return t.client.RPush(ctx, key, values).Err()
}

// LPop 列表左推出
func (t *tool) LPop(ctx context.Context, key string) (string, error) {
	return t.client.LPop(ctx, key).Result()
}

// LPopToStruct 列表左推出并转换
func (t *tool) LPopToStruct(ctx context.Context, key string, result interface{}) error {
	return toStruct(t, ctx, key, result, func(t *tool, ctx context.Context, key string) (string, error) {
		return t.client.LPop(ctx, key).Result()
	})
}

// RPop 列表右推出
func (t *tool) RPop(ctx context.Context, key string) (string, error) {
	return t.client.RPop(ctx, key).Result()
}

// RPopToStruct 列表右推出并转换
func (t *tool) RPopToStruct(ctx context.Context, key string, result interface{}) error {
	return toStruct(t, ctx, key, result, func(t *tool, ctx context.Context, key string) (string, error) {
		return t.client.RPop(ctx, key).Result()
	})
}

// LLen 获取列表长度
func (t *tool) LLen(ctx context.Context, key string) (int64, error) {
	return t.client.LLen(ctx, key).Result()
}

// LRange 获取列表范围
func (t *tool) LRange(ctx context.Context, key string, start, stop int64) ([]string, error) {
	return t.client.LRange(ctx, key, start, stop).Result()
}
