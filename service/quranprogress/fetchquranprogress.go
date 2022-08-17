package quranprogress

import (
	"github.com/gin-gonic/gin"
	"strconv"
	"tahfidz-backend/auth"
	"tahfidz-backend/model/enum"
	"tahfidz-backend/repository"
	"tahfidz-backend/util"
)

func FetchQuranProgress(context *gin.Context) {
	if !auth.Auth(context, enum.UserRoleEnum().EMPTY) {
		return
	}

	util.Response200(context, repository.FetchQuranProgress(), "")
}

func FetchQuranProgressByUserId(context *gin.Context) {
	if !auth.Auth(context, enum.UserRoleEnum().EMPTY) {
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
	if !auth.Auth(context, enum.UserRoleEnum().EMPTY) {
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

func FetchQuranProgressByMethod(context *gin.Context) {
	if !auth.Auth(context, enum.UserRoleEnum().EMPTY) {
		return
	}

	method := context.Param("method")

	util.Response200(context, repository.FetchQuranProgressByMethod(method), "")
}

func CountQuranProgressMethod(context *gin.Context) {
	if !auth.Auth(context, enum.UserRoleEnum().EMPTY) {
		return
	}

	userId := context.Param("userId")
	if !util.IsNumber(userId) {
		util.Response400(context, "", "user id harus angka")
		return
	}
	id, _ := strconv.Atoi(userId)

	util.Response200(context, repository.CountQuranProgressMethod(id), "")
}

func CurrentQuranProgress(context *gin.Context) {
	if !auth.Auth(context, enum.UserRoleEnum().EMPTY) {
		return
	}

	userId := context.Param("userId")
	if !util.IsNumber(userId) {
		util.Response400(context, "", "user id harus angka")
		return
	}
	id, _ := strconv.Atoi(userId)

	util.Response200(context, repository.CurrentQuranProgress(id), "")
}
