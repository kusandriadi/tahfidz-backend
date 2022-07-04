package main

import (
	"github.com/gin-gonic/gin"
	"tahfidz-backend/auth"
)

func main() {
	router := gin.Default()

	router.GET("/up", func(context *gin.Context) {
		context.String(200, "Tahfidz backend is online")
	})

	router.POST("/api/auth", auth.Login)

	router.Run()
}
