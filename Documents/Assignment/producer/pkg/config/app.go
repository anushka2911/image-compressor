package config

import (
	_ "github.com/go-sql-driver/mysql"
	"github.com/jinzhu/gorm"
	"github.com/sirupsen/logrus"
)

var db *gorm.DB

func Connect() {
	d, err := gorm.Open("mysql", "root:root@tcp(localhost:3306)/zocket?charset=utf8&parseTime=True&loc=Local")
	//here mysql is the database we want to connect to
	//root is the username
	//root is the password
	//localhost:3306 is the port
	//zocket is the database name
	//charset=utf8 is the character set
	//parseTime=True is to parse the time
	//loc=Local is the location

	if err != nil {
		logrus.Error("Could not connect to the database")
		panic(err)
	}

	db = d
}

// function of this file is to return the db connection
func GetDB() *gorm.DB {
	return db
}
