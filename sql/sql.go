package sql

import "gorm.io/gorm"

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
