package report

import (
	"github.com/gin-gonic/gin"
	"strconv"
	"tahfidz-backend/repository"
	"tahfidz-backend/util"
)

func FetchByUsername(context *gin.Context) {
	username := context.Param("username")

	util.Response200(context, repository.FetchUserByUsername(username, false), "")
}

func FetchByUserId(context *gin.Context) {
	id := context.Param("id")

	if !util.IsNumber(id) {
		util.Response400(context, "", "user id harus angka")
		return
	}
	idI, _ := strconv.Atoi(id)

	util.Response200(context, repository.FetchUserById(idI), "")
}

func CurrentQuranProgress(context *gin.Context) {
	userId := context.Param("userId")
	if !util.IsNumber(userId) {
		util.Response400(context, "", "user id harus angka")
		return
	}
	id, _ := strconv.Atoi(userId)

	util.Response200(context, repository.CurrentQuranProgress(id), "")
}

func CountQuranProgressMethod(context *gin.Context) {
	userId := context.Param("userId")
	if !util.IsNumber(userId) {
		util.Response400(context, "", "user id harus angka")
		return
	}
	id, _ := strconv.Atoi(userId)

	util.Response200(context, repository.CountQuranProgressMethod(id), "")
}

func FetchQuranProgressByUserId(context *gin.Context) {
	userId := context.Param("userId")
	limit := context.DefaultQuery("limit", "0")
	if !util.IsNumber(userId) {
		util.Response400(context, "", "user id harus angka")
		return
	}

	userIdI, _ := strconv.Atoi(userId)
	limitI, _ := strconv.Atoi(limit)

	util.Response200(context, repository.FetchQuranProgressByUserId(userIdI, limitI), "")
}
