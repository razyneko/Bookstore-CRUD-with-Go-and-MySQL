package config

import (
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)
// whole point of this file is to return a this var db
var (
	db *gorm.DB
)

func Connect(){
	// opening connection to MySQL
	dsn := "root:password@tcp(127.0.0.1:3306)/bookdb?charset=utf8mb4&parseTime=True&loc=Local"
    d, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})

	if err != nil{
		panic(err)
	}
	db = d
}

func GetDB() *gorm.DB{
	return db
}