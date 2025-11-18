package sql

import (
	"fmt"
	"net"
	"net/url"

	"github.com/oracle-samples/gorm-oracle/oracle"
	"gorm.io/gorm"
)

type Oracle struct {
	GeneralDB `yaml:",inline" mapstructure:",squash"`
}

func (m *Oracle) Dsn() string {
	dsn := fmt.Sprintf("oracle://%s:%s@%s/%s?%s", url.PathEscape(m.Username), url.PathEscape(m.Password),
		net.JoinHostPort(m.Path, m.Port), url.PathEscape(m.Dbname), m.Config)
	return dsn
}

// GormOracle 初始化oracle数据库
func GormOracle() *gorm.DB {
	return initOracleDatabase(OracleC)
}

// GormOracleByConfig 初始化Oracle数据库用过传入配置
func GormOracleByConfig(m Oracle) *gorm.DB {
	return initOracleDatabase(m)
}

// initOracleDatabase 初始化Oracle数据库的辅助函数
func initOracleDatabase(m Oracle) *gorm.DB {
	if m.Dbname == "" {
		return nil
	}
	// 数据库配置
	if db, err := gorm.Open(oracle.Open(m.Dsn()), m.Deploy()); err != nil {
		panic(err)
	} else {
		sqlDB, _ := db.DB()
		sqlDB.SetMaxIdleConns(m.MaxIdleConn)
		sqlDB.SetMaxOpenConns(m.MaxOpenConn)
		return db
	}
}
