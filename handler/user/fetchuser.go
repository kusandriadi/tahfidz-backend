package user

import (
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
	"tahfidz-backend/model"
	"tahfidz-backend/util"
)

func FetchAll(context *gin.Context) {
	var users []model.User
	db := context.MustGet("db").(*gorm.DB)

	db.Find(&users)

	util.Response200(context, users, "")
}

func FetchByRole(context *gin.Context) {
	role := context.Param("role")
	if util.Role(role) {
		util.Response400(context, "Role "+role+" tidak ditemukan.", "")
	}

	db := context.MustGet("db").(*gorm.DB)
	var users []model.User

	db.Where("role = ?", role).Find(&users)

	util.Response200(context, users, "")
}

func FetchById(context *gin.Context) {
	id := context.Param("id")
	if util.IsNumber(id) {
		util.Response400(context, "", "user id harus angka")
		return
	}

	db := context.MustGet("db").(*gorm.DB)
	var user model.User

	db.Find(&user, id)

	util.Response200(context, user, "")
}
