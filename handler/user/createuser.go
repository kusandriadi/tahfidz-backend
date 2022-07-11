package user

import (
	"tahfidz-backend/model"
	"tahfidz-backend/util"
	"time"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func Create(context *gin.Context) {
	var user model.User

	err := context.ShouldBindJSON(&user)
	if err != nil {
		util.Response400(context, "Failed to transform request body.", err.Error())
		return
	}

	passValidation, message := util.ValidateUser(&user)
	if !passValidation {
		util.Response400(context, message, "")
		return
	}

	db := context.MustGet("db").(*gorm.DB)

	user.CreatedDate = time.Now()

	if &user.BirthDate != nil {
		user.BirthDate, err = time.Parse("02-01-2006", user.BirthDate.Format("02-01-2006"))
	}

	createResult := db.Create(&user)

	if createResult.Error != nil {
		util.Response400(context, "Gagal membuat user baru.", err.Error())
		return
	}

	util.Response200(context, user, "")
	return
}
