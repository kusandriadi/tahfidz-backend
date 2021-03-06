package subject

import (
	"github.com/gin-gonic/gin"
	"tahfidz-backend/auth"
	"tahfidz-backend/repository"
	"tahfidz-backend/util"
)

func FetchSubjects(context *gin.Context) {
	if !auth.Auth(context) {
		return
	}

	util.Response200(context, repository.FetchSubjects(), "")
}

func FetchSubjectByType(context *gin.Context) {
	auth.Auth(context)
	role := context.Param("subjectType")

	util.Response200(context, repository.FetchSubjectByType(role), "")
}

func FetchSubjectByName(context *gin.Context) {
	auth.Auth(context)
	name := context.Param("name")

	util.Response200(context, repository.FetchSubjectByName(name), "")
}

func FetchSubjectById(context *gin.Context) {
	auth.Auth(context)
	id := context.Param("id")
	if !util.IsNumber(id) {
		util.Response400(context, "", "user id harus angka")
		return
	}

	util.Response200(context, repository.FetchSubjectById(id), "")
}
