package sql

type DBType string

const (
	mysqlType      DBType = "mysql"
	postgresqlType DBType = "pgsql"
	oracleType     DBType = "oracle"
	sqliteType     DBType = "sqlite"
)

var (
	ActiveDBName *string
	MysqlC       Mysql
	PgsqlC       Pgsql
	OracleC      Oracle
	SqliteC      Sqlite
)
