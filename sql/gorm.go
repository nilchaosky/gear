package sql

import (
	"github.com/nilchaosky/gear/logz"
	"gorm.io/gorm"
)

func Gorm(DBType DBType, DBName string) *gorm.DB {
	ActiveDBName = &DBName
	switch DBType {
	case mysqlType:
		DB = gormMysql()
	case postgresqlType:
		DB = gormPgSql()
	case oracleType:
		DB = gormOracle()
	case sqliteType:
		DB = gormSqlite()
	default:
		DB = gormMysql()
	}
	logz.Print.Info("Database connection successful")
	return DB
}
