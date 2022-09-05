package quranprogress

import (
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
	"tahfidz-backend/auth"
	"tahfidz-backend/model"
	"tahfidz-backend/util"
	"time"
)

func Create(context *gin.Context) {
	if !auth.Auth(context, nil) {
		return
	}

	var quranProgress model.QuranProgress

	err := context.ShouldBindJSON(&quranProgress)
	if err != nil {
		util.Response400(context, "Failed to transform request body.", err.Error())
		return
	}

	passValidation, message := util.ValidateQuranProgress(&quranProgress, true)
	if !passValidation {
		util.Response400(context, message, "")
		return
	}

	db := context.MustGet("db").(*gorm.DB)

	now := time.Now()
	quranProgress.CreatedDate = &now

	createResult := db.Create(&quranProgress)

	if createResult.Error != nil {
		util.Response400(context, "Gagal membuat progress hapalan.", createResult.Error.Error())
		return
	}

	util.Response200(context, quranProgress, "")
	return
}
