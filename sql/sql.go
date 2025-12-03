package sql

import (
	"time"

	"gorm.io/gorm"
)

type DBType string

const (
	mysqlType      DBType = "mysql"
	postgresqlType DBType = "pgsql"
	oracleType     DBType = "oracle"
	sqliteType     DBType = "sqlite"
)

var (
	DB           *gorm.DB
	ActiveDBName *string
	Mysql        MysqlCfg
	Pgsql        PgsqlCfg
	Oracle       OracleCfg
	Sqlite       SqliteCfg
)

type JsonTime time.Time

func (j JsonTime) MarshalJSON() ([]byte, error) {
	formatted := time.Time(j).Format(time.DateTime)
	return []byte(`"` + formatted + `"`), nil
}
