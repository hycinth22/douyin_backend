package main

import (
	"douyin_relation_service/consts"
	"douyin_relation_service/model"
	"douyin_relation_service/query"
	"time"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var db *gorm.DB

func init() {
	var err error
	db, err = gorm.Open(mysql.Open(consts.DSN), &gorm.Config{})
	if err != nil {
		panic(err)
	}
	sqlDB, err := db.DB()
	if err != nil {
		panic(err)
	}
	// SetMaxIdleConns sets the maximum number of connections in the idle connection pool.
	sqlDB.SetMaxIdleConns(10)
	// SetMaxOpenConns sets the maximum number of open connections to the database.
	sqlDB.SetMaxOpenConns(100)
	// SetConnMaxLifetime sets the maximum amount of time a connection may be reused.
	sqlDB.SetConnMaxLifetime(time.Hour)
	// Automatically update datable tables' structure
	db.AutoMigrate(&model.Message{})
	db.AutoMigrate(&model.Relation{})
	db.AutoMigrate(&model.UserDetail{})
	// Setup gorm generated code
	query.SetDefault(db)
}
