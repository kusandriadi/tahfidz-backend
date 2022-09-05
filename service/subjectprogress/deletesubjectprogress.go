package subjectprogress

import (
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
	"net/http"
	"tahfidz-backend/auth"
	"tahfidz-backend/model"
	"tahfidz-backend/util"
)

func Delete(context *gin.Context) {
	if !auth.Auth(context, nil) {
		return
	}

	id := context.Param("id")
	if !util.IsNumber(id) {
		util.Response400(context, "", "subject progress id harus angka")
		return
	}

	db := context.MustGet("db").(*gorm.DB)
	var subjectProgress model.SubjectProgress

	result := db.Find(&subjectProgress, id)

	if result.RowsAffected == 0 {
		util.Response(context, http.StatusBadRequest, "subject progress id "+id+" tidak ditemukan.")
		return
	}

	db.Model(&subjectProgress).Update("markForDelete", false)
	util.Response200(context, "", "Berhasil menghapus subject progress dengan id "+id)
}
