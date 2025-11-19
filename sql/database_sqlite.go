package sql

import (
	"path/filepath"

	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

type SqliteCfg struct {
	GeneralDB `yaml:",inline" mapstructure:",squash"`
}

func (s *SqliteCfg) dsn() string {
	return filepath.Join(s.Path, s.Dbname+".db")
}

// gormSqlite 初始化Sqlite数据库
func gormSqlite() *gorm.DB {
	return initSqliteDatabase(Sqlite)
}

// GormSqliteByconfig 初始化Sqlite数据库用过传入配置
func GormSqliteByconfig(s SqliteCfg) *gorm.DB {
	return initSqliteDatabase(s)
}

// initSqliteDatabase 初始化Sqlite数据库辅助函数
func initSqliteDatabase(s SqliteCfg) *gorm.DB {
	if s.Dbname == "" {
		return nil
	}

	// 数据库配置
	if db, err := gorm.Open(sqlite.Open(s.dsn()), s.deploy()); err != nil {
		panic(err)
	} else {
		sqlDB, _ := db.DB()
		sqlDB.SetMaxIdleConns(s.MaxIdleConn)
		sqlDB.SetMaxOpenConns(s.MaxOpenConn)
		return db
	}
}
