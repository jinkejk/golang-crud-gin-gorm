package config

import (
	"fmt"
	"golang-crud-gin/helper"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

const (
	host     = "10.239.215.11"
	port     = 3306
	user     = "root"
	password = "123456"
	dbName   = "liops"
)

func DatabaseConnection() *gorm.DB {
	sqlInfo := fmt.Sprintf("%s:%s@tcp(%s:%d)/%s?charset=utf8mb4&parseTime=True&loc=Local", user, password, host, port, dbName)
	db, err := gorm.Open(mysql.Open(sqlInfo), &gorm.Config{})
	helper.ErrorPanic(err)

	// 连接池相关
	sqlDB, err := db.DB()
	if sqlDB != nil {
		sqlDB.SetMaxIdleConns(10)
		sqlDB.SetMaxOpenConns(100)
	}

	return db
}
