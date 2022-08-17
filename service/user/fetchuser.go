package user

import (
	"github.com/gin-gonic/gin"
	"strconv"
	"tahfidz-backend/auth"
	"tahfidz-backend/model/enum"
	"tahfidz-backend/repository"
	"tahfidz-backend/util"
)

func FetchAll(context *gin.Context) {
	if !auth.Auth(context, enum.UserRoleEnum().EMPTY) {
		return
	}

	util.Response200(context, repository.FetchUsers(), "")
}

func FetchByRole(context *gin.Context) {
	if !auth.Auth(context, enum.UserRoleEnum().EMPTY) {
		return
	}

	role := context.Param("role")
	if !util.Role(role) {
		util.Response400(context, "Role "+role+" tidak ditemukan.", "")
		return
	}

	util.Response200(context, repository.FetchUserByRole(role), "")
}

func FetchByName(context *gin.Context) {
	if !auth.Auth(context, enum.UserRoleEnum().EMPTY) {
		return
	}

	name := context.Param("name")

	util.Response200(context, repository.FetchUserByName(name), "")
}

func FetchByUsername(context *gin.Context) {
	if !auth.Auth(context, enum.UserRoleEnum().EMPTY) {
		return
	}

	username := context.Param("username")

	util.Response200(context, repository.FetchUserByUsername(username), "")
}

func FetchById(context *gin.Context) {
	if !auth.Auth(context, enum.UserRoleEnum().EMPTY) {
		return
	}

	id := context.Param("id")
	if !util.IsNumber(id) {
		util.Response400(context, "", "user id harus angka")
		return
	}
	idI, _ := strconv.Atoi(id)

	util.Response200(context, repository.FetchUserById(idI), "")
}
