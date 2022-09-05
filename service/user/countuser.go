package user

import (
	"tahfidz-backend/auth"
	"tahfidz-backend/repository"
	"tahfidz-backend/util"

	"github.com/gin-gonic/gin"
)

func Count(context *gin.Context) {
	if !auth.Auth(context, nil) {
		return
	}

	util.Response200(context, repository.CountUser(), "")
}
