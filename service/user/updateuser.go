package user

import (
	"tahfidz-backend/auth"
	"tahfidz-backend/model"
	"tahfidz-backend/model/enum"
	"tahfidz-backend/util"
	"time"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func Update(context *gin.Context) {
	if !auth.Auth(context, []string{enum.UserRoleEnum().ADMIN}) {
		return
	}

	var updatedUser model.User

	err := context.ShouldBindJSON(&updatedUser)
	if err != nil {
		util.Response400(context, "Failed to transform request body.", err.Error())
		return
	}

	passValidation, message := util.ValidateUser(&updatedUser)
	if !passValidation {
		util.Response400(context, message, "")
		return
	}

	db := context.MustGet("db").(*gorm.DB)
	var existingUser model.User
	existingResult := db.Find(&existingUser, updatedUser.Id)
	if existingResult.RowsAffected == 0 {
		util.Response400(context, "Gagal mengubah data, tidak ada user "+string(updatedUser.Id), "")
		return
	}

	if len(updatedUser.UserBirthDate) > 0 {
		var formattedDate, _ = time.Parse("02-01-2006", updatedUser.UserBirthDate)
		updatedUser.BirthDate = &formattedDate
	}

	existingUser.Name = updatedUser.Name
	existingUser.BirthPlace = updatedUser.BirthPlace
	existingUser.Address = updatedUser.Address
	existingUser.Password = updatedUser.Password
	existingUser.LastEducation = updatedUser.LastEducation
	existingUser.City = updatedUser.City
	existingUser.Guardian = updatedUser.Guardian
	existingUser.Role = updatedUser.Role
	existingUser.GuardianPhone = updatedUser.GuardianPhone
	existingUser.UserPhone = updatedUser.UserPhone

	createResult := db.Save(&existingUser)

	if createResult.Error != nil {
		util.Response400(context, "Gagal mengubah data user.", createResult.Error.Error())
		return
	}

	util.Response200(context, updatedUser, "")
	return
}
