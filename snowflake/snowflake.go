package snowflake

import "github.com/GUAIK-ORG/go-snowflake/snowflake"

var Builder *snowflake.Snowflake

func RegisterSnowflake(datacenterId, workerId int) *snowflake.Snowflake {
	sf, err := snowflake.NewSnowflake(int64(datacenterId), int64(workerId))
	if err != nil {
		panic(err)
	}
	Builder = sf
	return sf
}
