package subject

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
		util.Response400(context, "", "subject id harus angka")
		return
	}

	db := context.MustGet("db").(*gorm.DB)
	var subject model.Subject

	result := db.Find(&subject, id)

	if result.RowsAffected == 0 {
		util.Response(context, http.StatusBadRequest, "subject id "+id+" tidak ditemukan.")
		return
	}

	db.Model(&subject).Update("markForDelete", true)
	util.Response200(context, "", "Berhasil menghapus subject dengan subject id "+id)
}
