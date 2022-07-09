package handler

import (
	"log"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func ConnectToDatabse() {
	dsn := "root:password@tcp(103.189.234.120:3306)/tunasmulia"
	_, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatal("DB Connection Error " + err.Error())
	}

	log.Println("DB Connection succeeded")
}
