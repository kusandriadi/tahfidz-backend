package subjectprogress

import (
	"github.com/gin-gonic/gin"
	"strconv"
	"tahfidz-backend/auth"
	"tahfidz-backend/model/enum"
	"tahfidz-backend/repository"
	"tahfidz-backend/util"
)

func FetchSubjectProgress(context *gin.Context) {
	if !auth.Auth(context, enum.UserRoleEnum().EMPTY) {
		return
	}

	util.Response200(context, repository.FetchSubjectProgress(), "")
}

func FetchSubjectProgressByUserIdAndSubjectId(context *gin.Context) {
	if !auth.Auth(context, enum.UserRoleEnum().EMPTY) {
		return
	}

	userId := context.Param("userId")
	if !util.IsNumber(userId) {
		util.Response400(context, "", "user id harus angka")
		return
	}
	uId, _ := strconv.Atoi(userId)

	subjectId := context.Param("subjectId")
	if !util.IsNumber(subjectId) {
		util.Response400(context, "", "subject id harus angka")
		return
	}

	sId, _ := strconv.Atoi(subjectId)

	util.Response200(context, repository.FetchSubjectProgressByUserIdAndSubjectId(uId, sId), "")
}

func FetchSubjectProgressBySubjectId(context *gin.Context) {
	if !auth.Auth(context, enum.UserRoleEnum().EMPTY) {
		return
	}

	subjectId := context.Param("subjectId")
	if !util.IsNumber(subjectId) {
		util.Response400(context, "", "subject id harus angka")
		return
	}

	id, _ := strconv.Atoi(subjectId)

	util.Response200(context, repository.FetchSubjectProgressBySubjectId(id), "")
}

func FetchSubjectProgressByUserId(context *gin.Context) {
	if !auth.Auth(context, enum.UserRoleEnum().EMPTY) {
		return
	}

	userId := context.Param("userId")
	if !util.IsNumber(userId) {
		util.Response400(context, "", "user id harus angka")
		return
	}

	id, _ := strconv.Atoi(userId)

	util.Response200(context, repository.FetchSubjectProgressByUserId(id), "")
}
