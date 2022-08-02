package quranprogress

import (
	"github.com/gin-gonic/gin"
	"tahfidz-backend/auth"
	"tahfidz-backend/repository"
	"tahfidz-backend/util"
)

func FetchQuranProgress(context *gin.Context) {
	if !auth.Auth(context) {
		return
	}

	util.Response200(context, repository.FetchQuranProgress(), "")
}

func FetchQuranProgressByUserId(context *gin.Context) {
	if !auth.Auth(context) {
		return
	}

	userId := context.Param("userId")

	util.Response200(context, repository.FetchQuranProgressByUserId(userId), "")
}

func FetchQuranProgressByUserIdAndMethod(context *gin.Context) {
	if !auth.Auth(context) {
		return
	}

	userId := context.Param("userId")
	method := context.Param("method")

	util.Response200(context, repository.FetchQuranProgressByUserIdAndMethod(userId, method), "")
}
