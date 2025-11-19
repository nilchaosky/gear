package sql

import (
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
	return DB
}

func bizModel() error {
	err := DB.AutoMigrate()
	if err != nil {
		return err
	}
	return nil
}
