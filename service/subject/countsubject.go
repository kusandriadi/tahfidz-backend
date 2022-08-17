package subject

import (
	"tahfidz-backend/auth"
	"tahfidz-backend/model/enum"
	"tahfidz-backend/repository"
	"tahfidz-backend/util"

	"github.com/gin-gonic/gin"
)

func Count(context *gin.Context) {
	if !auth.Auth(context, enum.UserRoleEnum().EMPTY) {
		return
	}

	util.Response200(context, repository.CountSubject(), "")
}
