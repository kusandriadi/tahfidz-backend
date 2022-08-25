package subjectprogress

import (
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
	"tahfidz-backend/auth"
	"tahfidz-backend/model"
	"tahfidz-backend/model/enum"
	"tahfidz-backend/util"
	"time"
)

func Create(context *gin.Context) {
	if !auth.Auth(context, enum.UserRoleEnum().EMPTY) {
		return
	}

	var subjectProgress model.SubjectProgress

	err := context.ShouldBindJSON(&subjectProgress)
	if err != nil {
		util.Response400(context, "Failed to transform request body.", err.Error())
		return
	}

	passValidation, message := util.ValidateSubjectProgress(&subjectProgress, true)
	if !passValidation {
		util.Response400(context, message, "")
		return
	}

	db := context.MustGet("db").(*gorm.DB)

	now := time.Now()
	subjectProgress.CreatedDate = &now
	createResult := db.Create(&subjectProgress)

	if createResult.Error != nil {
		util.Response400(context, "Gagal membuat progress pelajaran.", createResult.Error.Error())
		return
	}

	util.Response200(context, subjectProgress, "")
	return
}
