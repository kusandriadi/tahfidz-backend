package subjectprogress

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
	var existingSubjectProgress model.SubjectProgress
	existingResult := db.Find(&existingSubjectProgress, subjectProgress.Id)
	if existingResult.RowsAffected == 0 {
		util.Response400(context, "Gagal mengubah progress pelajaran karena belum tersedia", "")
		return
	}

	existingSubjectProgress.Presence = subjectProgress.Presence

	updateResult := db.Create(&existingSubjectProgress)

	if updateResult.Error != nil {
		util.Response400(context, "Gagal membuat progress pelajaran.", updateResult.Error.Error())
		return
	}

	util.Response200(context, subjectProgress, "")
	return
}
