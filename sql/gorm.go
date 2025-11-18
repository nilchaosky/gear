package sql

import "gorm.io/gorm"

type Grom struct {
}

func (g *Grom) GetDB(DBType DBType, DBName string) *gorm.DB {
	switch DBType {
	case MysqlType:
		ACTIVE_DBNAME = &DBName
		return GormMysql()
	//case "pgsql":
	//	global.ACTIVE_DBNAME = &global.CONFIG.Pgsql.Dbname
	//	return GormPgSql()
	//case "oracle":
	//	global.ACTIVE_DBNAME = &global.CONFIG.Oracle.Dbname
	//	return GormOracle()
	//case "mssql":
	//	global.ACTIVE_DBNAME = &global.CONFIG.Mssql.Dbname
	//	return GormMssql()
	//case "sqlite":
	//	global.ACTIVE_DBNAME = &global.CONFIG.Sqlite.Dbname
	//	return GormSqlite()
	default:
		ACTIVE_DBNAME = &DBName
		return GormMysql()
	}
}
