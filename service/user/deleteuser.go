package user

import (
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
	"net/http"
	"tahfidz-backend/auth"
	"tahfidz-backend/model"
	"tahfidz-backend/model/enum"
	"tahfidz-backend/util"
)

func Delete(context *gin.Context) {
	if !auth.Auth(context, []string{enum.UserRoleEnum().ADMIN}) {
		return
	}

	id := context.Param("id")
	if !util.IsNumber(id) {
		util.Response400(context, "", "user id harus angka")
		return
	}

	db := context.MustGet("db").(*gorm.DB)
	var user model.User

	result := db.Find(&user, id)

	if result.RowsAffected == 0 {
		util.Response(context, http.StatusBadRequest, "user id "+id+" tidak ditemukan.")
		return
	}

	db.Model(&user).Update("markForDelete", false)
	util.Response200(context, "", "Berhasil menghapus user dengan user id "+id)
}
