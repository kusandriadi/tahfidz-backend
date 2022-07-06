package handler

import (
	"fmt"
	"log"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func ConnectToDatabse() {
	dsn := "root:@tcp(27.112.79.164:3306)/tahfidz?charset=utf8mb4&parseTime=True&loc=Local"
	_, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatal("DB Connection Error")
	}

	fmt.Println("DB Connection succeeded")
}
