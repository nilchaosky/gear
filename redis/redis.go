package redis

import (
	"context"

	"github.com/nilchaosky/gear/logz"
	"github.com/redis/go-redis/v9"
	"go.uber.org/zap"
)

func initRedisClient(c Config) (redis.UniversalClient, error) {
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
	pong, err := client.Ping(context.Background()).Result()
	if err != nil {
		logz.Print.Error("redis connect ping failed, err:", zap.String("name", c.Name), zap.Error(err))
		return nil, err
	}

	logz.Print.Info("redis connect ping response:", zap.String("name", c.Name), zap.String("pong", pong))
	return client, nil
}

func Register(c Config) {
	client, err := initRedisClient(c)
	if err != nil {
		panic(err)
	}
	ToolKit = newTool(client)
}

func RegisterList(list []Config) {
	ToolKits = map[string]*tool{}

	for _, c := range list {
		if c.Name != "" {
			client, err := initRedisClient(c)
			if err != nil {
				panic(err)
			}
			ToolKits[c.Name] = newTool(client)
		}
	}
}
