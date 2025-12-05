package snowflake

import "github.com/GUAIK-ORG/go-snowflake/snowflake"

var Builder *snowflake.Snowflake

func Register(datacenterID, workerID int) {
	sf, err := snowflake.NewSnowflake(int64(datacenterID), int64(workerID))
	if err != nil {
		panic(err)
	}
	Builder = sf
}
