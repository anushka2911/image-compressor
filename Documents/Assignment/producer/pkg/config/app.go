package config

import (
	_ "github.com/go-sql-driver/mysql"
	"github.com/jinzhu/gorm"
	"github.com/sirupsen/logrus"
)

var db *gorm.DB

func Connect() {
	d, err := gorm.Open("mysql", "root:root@tcp(localhost:3306)/Zocket?charset=utf8&parseTime=True&loc=Local")

	if err != nil {
		logrus.Error("Could not connect to the database")
		panic(err)
	}

	db = d
}

func GetDB() *gorm.DB {
	return db
}
