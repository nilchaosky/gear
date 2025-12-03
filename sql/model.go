package sql

import (
	"gorm.io/gorm"
	"gorm.io/plugin/optimisticlock"
)

type (
	Snowflake struct {
		ID int64 `gorm:"primarykey;column:id;autoIncrement:false" json:"id,string"` // 主键ID
	}
	AutoIncrement struct {
		ID int64 `gorm:"primarykey;column:id;" json:"id,string"` // 主键ID
	}
	OptimisticLock struct {
		Version optimisticlock.Version `gorm:"column:_version;type:int" json:"-"`
	}
	TimeModel struct {
		CreatedAt JsonTime `gorm:"column:created_at;type:datetime;not null" json:"created_at"` // 创建时间
		UpdatedAt JsonTime `gorm:"column:updated_at;type:datetime;not null" json:"updated_at"` // 更新时间
	}
	DeleteModel struct {
		DeletedAt gorm.DeletedAt `gorm:"index;column:deleted_at" json:"-"` //软删除
	}
)
