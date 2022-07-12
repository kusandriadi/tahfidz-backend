package main

import (
	"github.com/sirupsen/logrus"
	"tahfidz-backend/auth"
	"tahfidz-backend/service"
	"tahfidz-backend/service/subject"
	"tahfidz-backend/service/user"

	"github.com/gin-gonic/gin"
)

func main() {
	logrus.SetFormatter(&logrus.JSONFormatter{})
	router := gin.Default()

	db := service.ConnectToDatabase()
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
	router.GET("/api/users/name/:name", user.FetchByName)
	router.GET("/api/users/username/:username", user.FetchByUsername)
	router.GET("/api/users/count", user.Count)

	router.POST("/api/users", user.Create)

	router.PUT("/api/users", user.Update)

	router.DELETE("/api/users/:id", user.Delete)
}

func subjectApi(router *gin.Engine) {
	router.GET("/api/subjects", subject.FetchSubjects)
	router.GET("/api/subjects/:id", subject.FetchSubjectById)
	router.GET("/api/subjects/name/:name", subject.FetchSubjectByName)
	router.GET("/api/subjects/type/:type", subject.FetchSubjectByType)
	router.GET("/api/subjects/count", subject.Count)

	router.POST("/api/subjects", subject.Create)

	router.PUT("/api/subjects", subject.Update)

	router.DELETE("/api/subjects/:id", subject.Delete)
}
