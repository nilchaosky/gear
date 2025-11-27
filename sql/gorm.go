package sql

import (
	"github.com/nilchaosky/gear/logz"
	"gorm.io/gorm"
)

func RegisterGorm(DBType DBType, DBName string) *gorm.DB {
	ActiveDBName = &DBName
	switch DBType {
	case mysqlType:
		DB = initMysql()
	case postgresqlType:
		DB = initPgSql()
	case oracleType:
		DB = initOracle()
	case sqliteType:
		DB = initSqlite()
	default:
		DB = initMysql()
	}
	logz.Print.Info("Database connection successful")
	return DB
}
