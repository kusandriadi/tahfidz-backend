package main

import (
	"tahfidz-backend/auth"
	"tahfidz-backend/handler"
	"tahfidz-backend/handler/user"

	"github.com/gin-gonic/gin"
)

func main() {
	router := gin.Default()

	db := handler.ConnectToDatabase()
	router.Use(func(c *gin.Context) {
		c.Set("db", db)
	})

	authApi(router)
	userApi(router)
	subjectApi(router)

	router.Run(":8088")
}

func authApi(router *gin.Engine) {
	router.GET("/up", func(context *gin.Context) {
		context.String(200, "Tahfidz backend is online")
	})

	router.POST("/api/auth", auth.Login)
}

func userApi(router *gin.Engine) {
	router.GET("/api/users", user.FetchAll)
	router.GET("/api/users/:id", user.FetchById)
	router.GET("/api/users/role/:role", user.FetchByRole)
	router.GET("/api/users/count", user.Count)

	router.POST("/api/users", user.Create)

	router.PUT("/api/users", user.Update)

	router.DELETE("/api/users/:id", user.Delete)
}

func subjectApi(router *gin.Engine) {
}
