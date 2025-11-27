package sql

import (
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

type PgsqlCfg struct {
	GeneralDB `yaml:",inline" mapstructure:",squash"`
}

// dsn 基于配置文件获取 dsn
// Author [SliverHorn](https://github.com/SliverHorn)
func (p *PgsqlCfg) dsn() string {
	return "host=" + p.Path + " user=" + p.Username + " password=" + p.Password + " dbname=" + p.Dbname + " port=" + p.Port + " " + p.Config
}

// linkDsn 根据 dbname 生成 dsn
// Author [SliverHorn](https://github.com/SliverHorn)
func (p *PgsqlCfg) linkDsn(dbname string) string {
	return "host=" + p.Path + " user=" + p.Username + " password=" + p.Password + " dbname=" + dbname + " port=" + p.Port + " " + p.Config
}

// initPgSql 初始化 Postgresql 数据库
// Author [piexlmax](https://github.com/piexlmax)
// Author [SliverHorn](https://github.com/SliverHorn)
func initPgSql() *gorm.DB {
	return initPgSqlDatabase(Pgsql)
}

// InitPgSqlByConfig 初始化 Postgresql 数据库 通过指定参数
func InitPgSqlByConfig(p PgsqlCfg) *gorm.DB {
	return initPgSqlDatabase(p)
}

// initPgSqlDatabase 初始化 Postgresql 数据库的辅助函数
func initPgSqlDatabase(p PgsqlCfg) *gorm.DB {
	if p.Dbname == "" {
		return nil
	}
	pgsqlConfig := postgres.Config{
		DSN:                  p.dsn(), // DSN data source name
		PreferSimpleProtocol: false,
	}
	// 数据库配置
	if db, err := gorm.Open(postgres.New(pgsqlConfig), p.deploy()); err != nil {
		panic(err)
	} else {
		sqlDB, _ := db.DB()
		sqlDB.SetMaxIdleConns(p.MaxIdleConn)
		sqlDB.SetMaxOpenConns(p.MaxOpenConn)
		return db
	}
}
