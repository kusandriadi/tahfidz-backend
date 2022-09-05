package quranprogress

import (
	"github.com/gin-gonic/gin"
	"strconv"
	"tahfidz-backend/auth"
	"tahfidz-backend/model/enum"
	"tahfidz-backend/repository"
	"tahfidz-backend/util"
	"time"
)

func FetchQuranProgress(context *gin.Context) {
	if !auth.Auth(context, nil) {
		return
	}

	util.Response200(context, repository.FetchQuranProgress(), "")
}

func FetchQuranProgressByUserId(context *gin.Context) {
	if !auth.Auth(context, nil) {
		return
	}

	userId := context.Param("userId")
	limit := context.DefaultQuery("limit", "0")
	if !util.IsNumber(userId) {
		util.Response400(context, "", "user id harus angka")
		return
	}

	id, _ := strconv.Atoi(userId)
	limitI, _ := strconv.Atoi(limit)

	util.Response200(context, repository.FetchQuranProgressByUserId(id, limitI), "")
}

func FetchQuranProgressByUserIdAndMethod(context *gin.Context) {
	if !auth.Auth(context, nil) {
		return
	}

	userId := context.Param("userId")
	if !util.IsNumber(userId) {
		util.Response400(context, "", "user id harus angka")
		return
	}
	id, _ := strconv.Atoi(userId)

	method := context.Param("method")

	util.Response200(context, repository.FetchQuranProgressByUserIdAndMethodAndCreatedDate(id, method, time.Now()), "")
}

func FetchQuranProgressByMethod(context *gin.Context) {
	if !auth.Auth(context, nil) {
		return
	}

	method := context.Param("method")

	util.Response200(context, repository.FetchQuranProgressByMethod(method), "")
}

func CountQuranProgressMethod(context *gin.Context) {
	if !auth.Auth(context, nil) {
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
	if !auth.Auth(context, nil) {
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

func GetAllQuranProgress(context *gin.Context) {
	if !auth.Auth(context, []string{enum.UserRoleEnum().ADMIN, enum.UserRoleEnum().TEACHER}) {
		return
	}

	util.Response200(context, repository.GetAllQuranProgress(), "")
}

func GetAllQuranProgressByName(context *gin.Context) {
	if !auth.Auth(context, []string{enum.UserRoleEnum().ADMIN, enum.UserRoleEnum().TEACHER}) {
		return
	}

	name := context.Param("name")

	util.Response200(context, repository.GetAllQuranProgressByName(name), "")
}
