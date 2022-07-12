package user

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

	var newUser model.User

	err := context.ShouldBindJSON(&newUser)
	if err != nil {
		util.Response400(context, "Failed to transform request body.", err.Error())
		return
	}

	passValidation, message := util.ValidateUser(&newUser)
	if !passValidation {
		util.Response400(context, message, "")
		return
	}

	db := context.MustGet("db").(*gorm.DB)
	var existingUser model.User
	existingResult := db.Where("username = ?", newUser.Username).Find(&existingUser)
	if existingResult.RowsAffected > 0 {
		util.Response400(context, "Username "+newUser.Username+" sudah dipakai, silakan cari username lain.", "")
		return
	}

	newUser.CreatedDate = time.Now()

	if newUser.BirthDate != nil {
		var formattedDate, _ = time.Parse("02-01-2006", newUser.BirthDate.Format("02-01-2006"))
		newUser.BirthDate = &formattedDate
	}

	createResult := db.Create(&newUser)

	if createResult.Error != nil {
		util.Response400(context, "Gagal membuat newUser baru.", createResult.Error.Error())
		return
	}

	util.Response200(context, newUser, "")
	return
}
