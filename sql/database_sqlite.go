package sql

import (
	"path/filepath"

	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

type Sqlite struct {
	GeneralDB `yaml:",inline" mapstructure:",squash"`
}

func (s *Sqlite) dsn() string {
	return filepath.Join(s.Path, s.Dbname+".db")
}

// gormSqlite 初始化Sqlite数据库
func gormSqlite() *gorm.DB {
	return initSqliteDatabase(SqliteC)
}

// GormSqliteByConfig 初始化Sqlite数据库用过传入配置
func GormSqliteByConfig(s Sqlite) *gorm.DB {
	return initSqliteDatabase(s)
}

// initSqliteDatabase 初始化Sqlite数据库辅助函数
func initSqliteDatabase(s Sqlite) *gorm.DB {
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
