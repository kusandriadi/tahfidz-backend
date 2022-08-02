package quranprogress

import (
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
	"net/http"
	"tahfidz-backend/auth"
	"tahfidz-backend/model"
	"tahfidz-backend/util"
)

func Delete(context *gin.Context) {
	if !auth.Auth(context) {
		return
	}

	id := context.Param("id")
	if !util.IsNumber(id) {
		util.Response400(context, "", "quran progress id harus angka")
		return
	}

	db := context.MustGet("db").(*gorm.DB)
	var deletedQuranProgress model.QuranProgress

	result := db.Find(&deletedQuranProgress, id)

	if result.RowsAffected == 0 {
		util.Response(context, http.StatusBadRequest, "quran progress id "+id+" tidak ditemukan.")
		return
	}

	db.Model(&deletedQuranProgress).Update("markForDelete", false)
	util.Response200(context, "", "Berhasil menghapus quran progress dengan id "+id)
}
