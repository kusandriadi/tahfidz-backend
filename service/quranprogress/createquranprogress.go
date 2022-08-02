package quranprogress

import (
	"tahfidz-backend/auth"
	"tahfidz-backend/model"
	"tahfidz-backend/util"
	"time"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func Create(context *gin.Context) {
	if !auth.Auth(context) {
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

	quranProgress.CreatedDate = time.Now()

	createResult := db.Create(&quranProgress)

	if createResult.Error != nil {
		util.Response400(context, "Gagal membuat quran progress.", createResult.Error.Error())
		return
	}

	util.Response200(context, quranProgress, "")
	return
}
