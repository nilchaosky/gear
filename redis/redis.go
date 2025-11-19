package redis

import (
	"context"

	"github.com/redis/go-redis/v9"
)

var (
	Client    redis.UniversalClient
	ClientMap map[string]redis.UniversalClient
)

func initRedisClient(c config) (redis.UniversalClient, error) {
	var client redis.UniversalClient
	// 使用集群模式
	if c.Cluster.Enable {
		client = redis.NewClusterClient(&redis.ClusterOptions{
			Addrs:    c.Cluster.AddrList,
			Password: c.Password,
		})
	} else {
		// 使用单例模式
		client = redis.NewClient(&redis.Options{
			Addr:     c.Addr,
			Password: c.Password,
			DB:       c.DB,
		})
	}
	_, err := client.Ping(context.Background()).Result()
	if err != nil {
		return nil, err
	}
	return client, nil
}

func Redis(c config) {
	client, err := initRedisClient(c)
	if err != nil {
		panic(err)
	}
	Client = client
}

func RedisList(list []config) {
	redisMap := make(map[string]redis.UniversalClient)

	for _, c := range list {
		if c.Name != "" {
			client, err := initRedisClient(c)
			if err != nil {
				panic(err)
			}
			redisMap[c.Name] = client
		}
	}

	ClientMap = redisMap
}
