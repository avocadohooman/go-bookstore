package config

import (
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var (
	db * gorm.DB
)

func ConnectDb() {
	dsn := "host=localhost user=postgres password=TheAnswerIsPosgres42! dbname=go-db port=1111 sslmode=disable TimeZone=Asia/Shanghai"
	d, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		panic(err)
	}
	db = d
}

func GetDb() *gorm.DB {
	return db
}