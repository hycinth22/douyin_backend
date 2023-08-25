package model

import (
	"gorm.io/gorm"
	"time"
)

type Relation struct {
	ID         int64 `gorm:"primarykey"`
	CreatedAt  time.Time
	UpdatedAt  time.Time
	DeletedAt  gorm.DeletedAt `gorm:"index"`
	FromUserID int64
	ToUserID   int64
	IsFollow   bool
	PrimaryKey string `gorm:"primaryKey"` //联合主键
}
