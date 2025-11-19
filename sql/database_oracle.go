package sql

import (
	"fmt"
	"net"
	"net/url"

	"github.com/dzwvip/gorm-oracle"
	"gorm.io/gorm"
)

type OracleCfg struct {
	GeneralDB `yaml:",inline" mapstructure:",squash"`
}

func (m *OracleCfg) dsn() string {
	dsn := fmt.Sprintf("oracle://%s:%s@%s/%s?%s", url.PathEscape(m.Username), url.PathEscape(m.Password),
		net.JoinHostPort(m.Path, m.Port), url.PathEscape(m.Dbname), m.Config)
	return dsn
}

// gormOracle 初始化oracle数据库
func gormOracle() *gorm.DB {
	return initOracleDatabase(Oracle)
}

// GormOracleByconfig 初始化Oracle数据库用过传入配置
func GormOracleByconfig(m OracleCfg) *gorm.DB {
	return initOracleDatabase(m)
}

// initOracleDatabase 初始化Oracle数据库的辅助函数
func initOracleDatabase(m OracleCfg) *gorm.DB {
	if m.Dbname == "" {
		return nil
	}
	// 数据库配置
	if db, err := gorm.Open(oracle.Open(m.dsn()), m.deploy()); err != nil {
		panic(err)
	} else {
		sqlDB, _ := db.DB()
		sqlDB.SetMaxIdleConns(m.MaxIdleConn)
		sqlDB.SetMaxOpenConns(m.MaxOpenConn)
		return db
	}
}
