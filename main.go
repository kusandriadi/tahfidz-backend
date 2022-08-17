package main

import (
	"github.com/sirupsen/logrus"
	"tahfidz-backend/auth"
	"tahfidz-backend/service"
	"tahfidz-backend/service/quranprogress"
	"tahfidz-backend/service/subject"
	"tahfidz-backend/service/subjectprogress"
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
	router.Use(CORSMiddleware())

	authApi(router)
	userApi(router)
	subjectApi(router)
	quranProgressApi(router)
	subjectProgressApi(router)

	router.Run(":8088")
}

func CORSMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {

		c.Header("Access-Control-Allow-Origin", "*")
		c.Header("Access-Control-Allow-Credentials", "true")
		c.Header("Access-Control-Allow-Headers", "Content-Type, Content-Length, Accept-Encoding, X-CSRF-Token, Authorization, accept, origin, Cache-Control, X-Requested-With")
		c.Header("Access-Control-Allow-Methods", "POST,HEAD,PATCH, OPTIONS, GET, PUT")

		if c.Request.Method == "OPTIONS" {
			c.AbortWithStatus(204)
			return
		}

		c.Next()
	}
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

func quranProgressApi(router *gin.Engine) {
	router.GET("/api/quran-progress", quranprogress.FetchQuranProgress)
	router.GET("/api/quran-progress/user/:userId", quranprogress.FetchQuranProgressByUserId)
	router.GET("/api/quran-progress/user/:userId/method/:method", quranprogress.FetchQuranProgressByUserIdAndMethod)
	router.GET("/api/quran-progress/user/:userId/method/count", quranprogress.CountQuranProgressMethod)
	router.GET("/api/quran-progress/user/:userId/progress", quranprogress.CurrentQuranProgress)
	router.GET("/api/quran-progress/method/:method", quranprogress.FetchQuranProgressByMethod)

	router.POST("/api/quran-progress", quranprogress.Update)

	router.DELETE("/api/quran-progress/:id", quranprogress.Delete)
}

func subjectProgressApi(router *gin.Engine) {
	router.GET("/api/subject-progress", subjectprogress.FetchSubjectProgress)
	router.GET("/api/subject-progress/user/:userId", subjectprogress.FetchSubjectProgressByUserId)
	router.GET("/api/subject-progress/user/:userId/subject/:subjectId", subjectprogress.FetchSubjectProgressByUserIdAndSubjectId)
	router.GET("/api/subject-progress/subject/:subjectId", subjectprogress.FetchSubjectProgressBySubjectId)

	router.POST("/api/subject-progress", subjectprogress.Update)

	router.DELETE("/api/subject-progress/:id", subjectprogress.Delete)
}
