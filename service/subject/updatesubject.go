package subject

import (
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
	"tahfidz-backend/auth"
	"tahfidz-backend/model"
	"tahfidz-backend/model/enum"
	"tahfidz-backend/util"
)

func Update(context *gin.Context) {
	if !auth.Auth(context, []string{enum.UserRoleEnum().ADMIN}) {
		return
	}

	var updatedSubject model.Subject

	err := context.ShouldBindJSON(&updatedSubject)
	if err != nil {
		util.Response400(context, "Failed to transform request body.", err.Error())
		return
	}

	passValidation, message := util.ValidateSubject(&updatedSubject)
	if !passValidation {
		util.Response400(context, message, "")
		return
	}

	db := context.MustGet("db").(*gorm.DB)
	var existingSubject model.Subject
	existingResult := db.Find(&existingSubject, updatedSubject.Id)
	if existingResult.RowsAffected > 0 {
		util.Response400(context, "Nama Kajian/Pelajaran "+updatedSubject.Name+" sudah dipakai, silakan cari kajian/pelajaran lain.", "")
		return
	}

	existingSubject.Name = updatedSubject.Name
	existingSubject.Author = updatedSubject.Author
	existingSubject.Book = updatedSubject.Book
	existingSubject.Duration = updatedSubject.Duration

	createResult := db.Create(&existingSubject)

	if createResult.Error != nil {
		util.Response400(context, "Gagal mengubah kajian/pelajaran.", createResult.Error.Error())
		return
	}

	util.Response200(context, updatedSubject, "")
	return

}
