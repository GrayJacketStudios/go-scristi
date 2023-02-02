package config

import (
	"gorm.io/driver/mysql"
	_ "gorm.io/driver/mysql"
	"gorm.io/gorm"
	"os"
)

var (
	db *gorm.DB
)

func Connect() {
	if len(os.Getenv("BD_STRING")) == 0 {
		panic("BD_STRING its not defined.")
	}
	d, err := gorm.Open(mysql.Open(os.Getenv("BD_STRING")))
	if err != nil {
		panic(err)
	}
	db = d
}

func GetDB() *gorm.DB {
	return db
}
