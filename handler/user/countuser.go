package user

import (
	"tahfidz-backend/model"
	"tahfidz-backend/util"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func Count(context *gin.Context) {
	var userCount model.UserCount
	db := context.MustGet("db").(*gorm.DB)
	db.Select("DISTINCT ON (role) role", "total").Find(&userCount)

	util.Response200(context, userCount, "")
}
