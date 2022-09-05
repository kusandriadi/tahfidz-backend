package subject

import (
	"github.com/gin-gonic/gin"
	"strconv"
	"tahfidz-backend/auth"
	"tahfidz-backend/repository"
	"tahfidz-backend/util"
)

func FetchSubjects(context *gin.Context) {
	if !auth.Auth(context, nil) {
		return
	}

	util.Response200(context, repository.FetchSubjects(), "")
}

func FetchSubjectByType(context *gin.Context) {
	if !auth.Auth(context, nil) {
		return
	}

	subjectType := context.Param("type")

	util.Response200(context, repository.FetchSubjectByType(subjectType), "")
}

func FetchSubjectByName(context *gin.Context) {
	if !auth.Auth(context, nil) {
		return
	}

	name := context.Param("name")

	util.Response200(context, repository.FetchSubjectByName(name), "")
}

func FetchSubjectById(context *gin.Context) {
	if !auth.Auth(context, nil) {
		return
	}

	id := context.Param("id")
	if !util.IsNumber(id) {
		util.Response400(context, "", "user id harus angka")
		return
	}
	idI, _ := strconv.Atoi(id)

	util.Response200(context, repository.FetchSubjectById(idI), "")
}
