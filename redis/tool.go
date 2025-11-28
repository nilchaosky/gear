package redis

import (
	"github.com/redis/go-redis/v9"
)

var (
	ToolKit  *tool
	ToolKits map[string]*tool
)

type tool struct {
	client redis.UniversalClient

	toolGeneric
	toolString
	toolList
	toolSet
	toolSortedSet
	toolsHash
}

func newTool(client redis.UniversalClient) *tool {
	return &tool{
		client: client,
	}
}
