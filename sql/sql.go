package sql

import (
	"time"

	"gorm.io/gorm"
	"gorm.io/plugin/optimisticlock"
)

type DBType string

const (
	mysqlType      DBType = "mysql"
	postgresqlType DBType = "pgsql"
	oracleType     DBType = "oracle"
	sqliteType     DBType = "sqlite"
)

var (
	DB           *gorm.DB
	ActiveDBName *string
	Mysql        MysqlCfg
	Pgsql        PgsqlCfg
	Oracle       OracleCfg
	Sqlite       SqliteCfg
)

type (
	IdModel struct {
		ID int64 `gorm:"primarykey;column:id;autoIncrement:false" json:"id"` // 主键ID
	}
	AutoIncrementIdModel struct {
		ID int64 `gorm:"primarykey;column:id unsigned" json:"id"` // 主键ID
	}
	VersionModel struct {
		Version optimisticlock.Version `gorm:"column:_version;type:int unsigned" json:"-"`
	}
	TimeModel struct {
		CreatedAt time.Time `gorm:"column:created_at;not null" json:"created_at"` // 创建时间
		UpdatedAt time.Time `gorm:"column:updated_at;not null" json:"updated_at"` // 更新时间
	}
	DeleteModel struct {
		DeletedAt gorm.DeletedAt `gorm:"index;column:deleted_at" json:"-"` //软删除
	}
)
