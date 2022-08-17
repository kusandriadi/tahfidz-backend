package subject

import (
	"tahfidz-backend/auth"
	"tahfidz-backend/model"
	"tahfidz-backend/model/enum"
	"tahfidz-backend/util"
	"time"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func Create(context *gin.Context) {
	if !auth.Auth(context, enum.UserRoleEnum().ADMIN) {
		return
	}

	var newSubject model.Subject

	err := context.ShouldBindJSON(&newSubject)
	if err != nil {
		util.Response400(context, "Failed to transform request body.", err.Error())
		return
	}

	passValidation, message := util.ValidateSubject(&newSubject)
	if !passValidation {
		util.Response400(context, message, "")
		return
	}

	db := context.MustGet("db").(*gorm.DB)
	var existingUser model.User
	existingResult := db.Where("name = ?", newSubject.Name).Find(&existingUser)
	if existingResult.RowsAffected > 0 {
		util.Response400(context, "Kajian/Pelajaran "+newSubject.Name+" sudah ada, silakan cari kajian/pelajaran lain.", "")
		return
	}

	newSubject.CreatedDate = time.Now()

	createResult := db.Create(&newSubject)

	if createResult.Error != nil {
		util.Response400(context, "Gagal membuat subject baru.", createResult.Error.Error())
		return
	}

	util.Response200(context, newSubject, "")
	return
}
