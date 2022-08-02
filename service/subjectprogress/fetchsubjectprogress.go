package subjectprogress

import (
	"github.com/gin-gonic/gin"
	"tahfidz-backend/auth"
	"tahfidz-backend/repository"
	"tahfidz-backend/util"
)

func FetchSubjectProgress(context *gin.Context) {
	if !auth.Auth(context) {
		return
	}

	util.Response200(context, repository.FetchSubjectProgress(), "")
}

func FetchSubjectProgressByUserIdAndSubjectId(context *gin.Context) {
	if !auth.Auth(context) {
		return
	}

	userId := context.Param("userId")
	subjectId := context.Param("subjectId")

	util.Response200(context, repository.FetchSubjectProgressByUserIdAndSubjectId(userId, subjectId), "")
}

func FetchSubjectProgressBySubjectId(context *gin.Context) {
	if !auth.Auth(context) {
		return
	}

	subjectId := context.Param("subjectId")

	util.Response200(context, repository.FetchSubjectProgressBySubjectId(subjectId), "")
}

func FetchSubjectProgressByUserId(context *gin.Context) {
	if !auth.Auth(context) {
		return
	}

	userId := context.Param("userId")

	util.Response200(context, repository.FetchSubjectProgressByUserId(userId), "")
}
