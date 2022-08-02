package quranprogress

import (
	"github.com/gin-gonic/gin"
	"strconv"
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
	if !util.IsNumber(userId) {
		util.Response400(context, "", "user id harus angka")
		return
	}

	id, _ := strconv.Atoi(userId)

	util.Response200(context, repository.FetchQuranProgressByUserId(id), "")
}

func FetchQuranProgressByUserIdAndMethod(context *gin.Context) {
	if !auth.Auth(context) {
		return
	}

	userId := context.Param("userId")
	if !util.IsNumber(userId) {
		util.Response400(context, "", "user id harus angka")
		return
	}
	id, _ := strconv.Atoi(userId)

	method := context.Param("method")

	util.Response200(context, repository.FetchQuranProgressByUserIdAndMethod(id, method), "")
}
