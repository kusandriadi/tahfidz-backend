package user

import (
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
	"tahfidz-backend/model"
	"tahfidz-backend/util"
)

func Update(context *gin.Context) {
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

	createResult := db.Create(&user)

	if createResult.Error != nil {
		util.Response400(context, "Gagal mengubah user baru.", err.Error())
		return
	}

	util.Response200(context, user, "")
	return

}
