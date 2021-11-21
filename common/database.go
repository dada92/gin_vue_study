package common

import (
	"fmt"
	"gin_vue_study/model"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var db *gorm.DB

func init() {
	const (
		user    = "root"
		passwd  = "123"
		host    = "localhost"
		post    = "3306"
		dbname  = "ginvue"
		charset = "utf8mb4"
	)

	dsn := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=%s&parseTime=True&loc=Local",
		user,
		passwd,
		host,
		post,
		dbname,
		charset)
	database, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		panic("failed to connect database, err: " + err.Error())
	}
	database.AutoMigrate(&model.User{})
	db = database
	return
}

func GetDB() *gorm.DB {
	return db
}
