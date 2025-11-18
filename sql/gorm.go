package sql

import (
	"gorm.io/gorm"
)

func Gorm(DBType DBType, DBName string) *gorm.DB {
	ActiveDBName = &DBName
	switch DBType {
	case MysqlType:
		return GormMysql()
	case PostgresqlType:
		return GormPgSql()
	case OracleType:
		return GormOracle()
	case SqliteType:
		return GormSqlite()
	default:
		return GormMysql()
	}
}

func bizModel(db *gorm.DB) error {
	err := db.AutoMigrate()
	if err != nil {
		return err
	}
	return nil
}
