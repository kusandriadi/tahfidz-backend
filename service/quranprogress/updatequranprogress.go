package quranprogress

import (
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
	"tahfidz-backend/auth"
	"tahfidz-backend/model"
	"tahfidz-backend/model/enum"
	"tahfidz-backend/util"
)

func Update(context *gin.Context) {
	if !auth.Auth(context, enum.UserRoleEnum().EMPTY) {
		return
	}

	var quranProgress model.QuranProgress

	err := context.ShouldBindJSON(&quranProgress)
	if err != nil {
		util.Response400(context, "Failed to transform request body.", err.Error())
		return
	}

	passValidation, message := util.ValidateQuranProgress(&quranProgress)
	if !passValidation {
		util.Response400(context, message, "")
		return
	}

	db := context.MustGet("db").(*gorm.DB)
	var existingQuranProgress model.QuranProgress
	existingResult := db.Find(&existingQuranProgress, quranProgress.Id)
	if existingResult.RowsAffected == 0 {
		util.Response400(context, "Gagal mengubah progress hapalan karena belum tersedia", "")
		return
	}

	existingQuranProgress.Surat = quranProgress.Surat
	existingQuranProgress.Ayat = quranProgress.Ayat
	existingQuranProgress.Juz = quranProgress.Juz

	updateResult := db.Create(&existingQuranProgress)

	if updateResult.Error != nil {
		util.Response400(context, "Gagal membuat progress hapalan.", updateResult.Error.Error())
		return
	}

	util.Response200(context, quranProgress, "")
	return
}
