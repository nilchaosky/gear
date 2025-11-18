package sql

type DBType string

const (
	MysqlType      DBType = "mysql"
	PostgresqlType DBType = "pgsql"
	OracleType     DBType = "oracle"
	SqliteType     DBType = "sqlite"
)

var (
	ActiveDBName       *string
	MysqlC             Mysql
	PgsqlC             Pgsql
	OracleC            Oracle
	SqliteC            Sqlite
	DisableAutoMigrate bool
)
