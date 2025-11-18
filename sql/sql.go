package sql

type DBType string

const (
	MysqlType      DBType = "mysql"
	PostgresqlType DBType = "pgsql"
	OracleType     DBType = "oracle"
	SqliteType     DBType = "sqlite"
	MssqlType      DBType = "mssql"
)

var (
	ACTIVE_DBNAME *string
	MysqlC        MysqlConfig
)
