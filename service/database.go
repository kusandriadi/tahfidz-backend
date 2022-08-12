package service

import (
	"log"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func ConnectToDatabase() *gorm.DB {
	dsn := "root:!1Password@tcp(localhost:3306)/tunasmulia?loc=Asia%2FJakarta&charset=utf8&parseTime=True"
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatal("DB Connection Error " + err.Error())
	}

	return db
}
