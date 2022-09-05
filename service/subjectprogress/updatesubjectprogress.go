package subjectprogress

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

	var subjectProgress model.SubjectProgress

	err := context.ShouldBindJSON(&subjectProgress)
	if err != nil {
		util.Response400(context, "Failed to transform request body.", err.Error())
		return
	}

	passValidation, message := util.ValidateSubjectProgress(&subjectProgress, false)
	if !passValidation {
		util.Response400(context, message, "")
		return
	}

	db := context.MustGet("db").(*gorm.DB)

	updateResult := db.Model(&subjectProgress).
		Where("id = ?", subjectProgress.Id).
		Update("presence", subjectProgress.Presence)

	if updateResult.Error != nil {
		util.Response400(context, "Gagal membuat progress pelajaran.", updateResult.Error.Error())
		return
	}

	util.Response200(context, subjectProgress, "")
	return
}
