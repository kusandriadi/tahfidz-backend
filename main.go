package main

import (
	"tahfidz-backend/auth"
	"tahfidz-backend/handler"

	"github.com/gin-gonic/gin"
)

func main() {
	router := gin.Default()

	handler.ConnectToDatabse()

	router.GET("/up", func(context *gin.Context) {
		context.String(200, "Tahfidz backend is online")
	})

	router.POST("/api/auth", auth.Login)

	router.Run(":8088")
}
