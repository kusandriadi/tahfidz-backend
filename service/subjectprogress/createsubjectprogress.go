package subjectprogress

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

	var subjectProgress model.SubjectProgress

	err := context.ShouldBindJSON(&subjectProgress)
	if err != nil {
		util.Response400(context, "Failed to transform request body.", err.Error())
		return
	}

	passValidation, message := util.ValidateSubjectProgress(&subjectProgress)
	if !passValidation {
		util.Response400(context, message, "")
		return
	}

	db := context.MustGet("db").(*gorm.DB)

	subjectProgress.CreatedDate = time.Now()

	createResult := db.Create(&subjectProgress)

	if createResult.Error != nil {
		util.Response400(context, "Gagal membuat subject progress.", createResult.Error.Error())
		return
	}

	util.Response200(context, subjectProgress, "")
	return
}
