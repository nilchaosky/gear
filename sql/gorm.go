package sql

import (
	"gorm.io/gorm"
)

func Gorm(DBType DBType, DBName string) *gorm.DB {
	ActiveDBName = &DBName
	switch DBType {
	case mysqlType:
		return gormMysql()
	case postgresqlType:
		return gormPgSql()
	case oracleType:
		return gormOracle()
	case sqliteType:
		return gormSqlite()
	default:
		return gormMysql()
	}
}

func bizModel(db *gorm.DB) error {
	err := db.AutoMigrate()
	if err != nil {
		return err
	}
	return nil
}
