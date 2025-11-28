package redis

import "context"

type toolSet interface {
	SAdd(ctx context.Context, key string, members ...interface{}) error
	SRem(ctx context.Context, key string, members ...interface{}) error
	SRandMember(ctx context.Context, key string) (string, error)
	SRandMemberToStruct(ctx context.Context, key string, value interface{}) error
	SIsMember(ctx context.Context, key string, member interface{}) (bool, error)
}

// SAdd 集合添加成员
func (t *tool) SAdd(ctx context.Context, key string, members ...interface{}) error {
	return t.client.SAdd(ctx, key, members).Err()

}

// SRem 集合移除一个或多个成员，如果成员不存在，则忽略
func (t *tool) SRem(ctx context.Context, key string, members ...interface{}) error {
	return t.client.SRem(ctx, key, members).Err()
}

// SRandMember 随机取一个元素
func (t *tool) SRandMember(ctx context.Context, key string) (string, error) {
	return t.client.SRandMember(ctx, key).Result()
}

// SRandMemberToStruct 随机取一个元素并转换
func (t *tool) SRandMemberToStruct(ctx context.Context, key string, value interface{}) error {
	return t.client.SRandMember(ctx, key).Scan(value)
}

// SIsMember 判断成员是否存在
func (t *tool) SIsMember(ctx context.Context, key string, member interface{}) (bool, error) {
	return t.client.SIsMember(ctx, key, member).Result()
}
