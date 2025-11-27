package redis

import (
	"context"
	"time"
)

type toolGeneric interface {
	Del(ctx context.Context, key ...string) error
	Exists(ctx context.Context, key string) error
	Expire(ctx context.Context, key string, exp time.Duration) error
	TTL(ctx context.Context, key string) (time.Duration, error)
	Persist(ctx context.Context, key string) error
	Keys(ctx context.Context, pattern string) ([]string, error)
	FlushDB(ctx context.Context) error
	FlushAll(ctx context.Context) error
}

// Del 删除 key
func (t *tool) Del(ctx context.Context, key ...string) error {
	return t.client.Del(ctx, key...).Err()
}

// Exists 判断 key 是否存在
func (t *tool) Exists(ctx context.Context, key string) error {
	return t.client.Exists(ctx, key).Err()
}

// Expire 设置过期时间
func (t *tool) Expire(ctx context.Context, key string, exp time.Duration) error {
	return t.client.Expire(ctx, key, exp).Err()
}

// TTL 查看剩余过期时间
func (t *tool) TTL(ctx context.Context, key string) (time.Duration, error) {
	return t.client.TTL(ctx, key).Result()
}

// Persist 移除 key 的过期时间 持久化key
func (t *tool) Persist(ctx context.Context, key string) error {
	return t.client.Persist(ctx, key).Err()
}

// Keys 获取匹配 key
func (t *tool) Keys(ctx context.Context, pattern string) ([]string, error) {
	return t.client.Keys(ctx, pattern).Result()
}

// FlushDB 清空当前 DB
func (t *tool) FlushDB(ctx context.Context) error {
	return t.client.FlushDB(ctx).Err()
}

// FlushAll 清空全部 DB
func (t *tool) FlushAll(ctx context.Context) error {
	return t.client.FlushAll(ctx).Err()
}
