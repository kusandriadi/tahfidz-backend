package quranprogress

import (
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
	"tahfidz-backend/auth"
	"tahfidz-backend/model"
	"tahfidz-backend/util"
)

func Update(context *gin.Context) {
	if !auth.Auth(context, nil) {
		return
	}

	var quranProgress model.QuranProgress

	err := context.ShouldBindJSON(&quranProgress)
	if err != nil {
		util.Response400(context, "Failed to transform request body.", err.Error())
		return
	}

	passValidation, message := util.ValidateQuranProgress(&quranProgress, false)
	if !passValidation {
		util.Response400(context, message, "")
		return
	}

	db := context.MustGet("db").(*gorm.DB)

	updateResult := db.Model(&quranProgress).
		Where("id = ?", quranProgress.Id).
		Updates(model.QuranProgress{Surat: quranProgress.Surat, Ayat: quranProgress.Ayat, Juz: quranProgress.Juz})

	if updateResult.Error != nil {
		util.Response400(context, "Gagal mengubah progress hapalan.", updateResult.Error.Error())
		return
	}

	util.Response200(context, quranProgress, "")
	return
}
