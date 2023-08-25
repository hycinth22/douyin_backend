package model

import (
	"gorm.io/gorm"
)

type UserDetail struct {
	gorm.Model

	UserID        int64 `gorm:"foreignKey:UserID;references:User.ID"`
	Name          string
	FollowCount   int64
	FollowerCount int64

	Avatar          string
	BackgroundImage string
	Signature       string

	TotalFavorited int64
	WorkCount      int64
	FavoriteCount  int64
}