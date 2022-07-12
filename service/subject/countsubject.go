package subject

import (
	"tahfidz-backend/auth"
	"tahfidz-backend/model"
	"tahfidz-backend/util"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func Count(context *gin.Context) {
	if !auth.Auth(context) {
		return
	}

	var subjectCount model.UserCount
	db := context.MustGet("db").(*gorm.DB)
	db.Select("DISTINCT ON (type) type", "total").Find(&subjectCount)

	util.Response200(context, subjectCount, "")
}
