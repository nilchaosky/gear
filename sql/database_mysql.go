package sql

import (
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

type MysqlCfg struct {
	GeneralDB `yaml:",inline" mapstructure:",squash"`
}

func (m *MysqlCfg) dsn() string {
	return m.Username + ":" + m.Password + "@tcp(" + m.Path + ":" + m.Port + ")/" + m.Dbname + "?" + m.Config
}

// initMysql 初始化Mysql数据库
// Author [piexlmax](https://github.com/piexlmax)
// Author [SliverHorn](https://github.com/SliverHorn)
// Author [ByteZhou-2018](https://github.com/ByteZhou-2018)
func initMysql() *gorm.DB {
	return initMysqlDatabase(Mysql)
}

// InitMysqlByConfig 通过传入配置初始化Mysql数据库
func InitMysqlByConfig(m MysqlCfg) *gorm.DB {
	return initMysqlDatabase(m)
}

// initMysqlDatabase 初始化Mysql数据库的辅助函数
func initMysqlDatabase(m MysqlCfg) *gorm.DB {
	if m.Dbname == "" {
		return nil
	}

	mysqlconfig := mysql.Config{
		DSN:                       m.dsn(), // DSN data source name
		DefaultStringSize:         191,     // string 类型字段的默认长度
		SkipInitializeWithVersion: false,   // 根据版本自动配置
	}
	// 数据库配置
	if db, err := gorm.Open(mysql.New(mysqlconfig), m.deploy()); err != nil {
		panic(err)
	} else {
		db.InstanceSet("gorm:table_options", "ENGINE="+m.Engine)
		sqlDB, _ := db.DB()
		sqlDB.SetMaxIdleConns(m.MaxIdleConn)
		sqlDB.SetMaxOpenConns(m.MaxOpenConn)
		return db
	}
}
