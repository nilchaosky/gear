package redis

import (
	"context"
	"encoding/json"
	"errors"
	"reflect"

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

func toStruct(t *tool, ctx context.Context, key string, result interface{}, fn func(t *tool, ctx context.Context, key string) (string, error)) error {
	if result == nil {
		return errors.New("result is nil")
	}
	rv := reflect.ValueOf(result)
	if rv.Kind() != reflect.Ptr {
		return errors.New("result must be a pointer to a struct/slice/map")
	}

	value, err := fn(t, ctx, key)
	if err != nil {
		return err
	}

	if value == "" {
		return errors.New("result is empty")
	}

	rv.Elem().Set(reflect.Zero(rv.Elem().Type()))
	err = json.Unmarshal([]byte(value), result)
	if err != nil {
		return err
	}
	return nil
}
