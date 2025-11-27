package redis

import (
	"context"

	"github.com/redis/go-redis/v9"
)

type toolSortedSet interface {
	ZAdd(ctx context.Context, key string, score float64, member interface{}) error
	ZRem(ctx context.Context, key string, member ...interface{}) error
}

// ZAdd 有序集合添加成员
func (t *tool) ZAdd(ctx context.Context, key string, score float64, member interface{}) error {
	return t.client.ZAdd(ctx, key, redis.Z{
		Score:  score,
		Member: member,
	}).Err()
}

// ZRem 有序集合删除一个或多个成员
func (t *tool) ZRem(ctx context.Context, key string, member ...interface{}) error {
	return t.client.ZRem(ctx, key, member).Err()
}
