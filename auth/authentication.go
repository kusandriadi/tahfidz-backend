package auth

import (
	"fmt"
	"github.com/sirupsen/logrus"
	"golang.org/x/crypto/bcrypt"
	"net/http"
	"tahfidz-backend/model"
	"tahfidz-backend/repository"
	"tahfidz-backend/util"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v4"
)

func Auth(context *gin.Context) bool {
	tokenString := context.Request.Header.Get("Authorization")
	if len(tokenString) <= 0 {
		util.Response(context, http.StatusBadRequest,
			"Anda tidak diperbolehkan untuk melihat halaman ini. Silakan hubungi Admin.")
		return false
	}

	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		if jwt.GetSigningMethod("HS256") != token.Method {
			return nil, fmt.Errorf("Unexpected signing method: %v", token.Header["alg"])
		}

		return []byte("secret"), nil
	})

	if token != nil && err == nil {
		logrus.Info("Token verified")
	} else {
		util.Response(context, http.StatusUnauthorized,
			"Anda tidak diperbolehkan untuk masuk. Silakan hubungi Admin.")
		context.Abort()
		return false
	}

	return true
}

func Login(context *gin.Context) {
	var user model.User
	err := context.ShouldBindJSON(&user)

	if err != nil {
		util.Response400(context, "Bad Request.", "")
		return
	}

	if !(len(user.Username) > 0 && len(user.Password) > 0) {
		util.Response400(context, "Username dan Password tidak boleh kosong.", "")
		return
	}

	var existingUser = repository.FetchUserByUsername(user.Username)
	if err := bcrypt.CompareHashAndPassword([]byte(existingUser.Password), []byte(user.Password)); err != nil {
		util.Response400(context, "Username dan/atau Password tidak sesuai.", "hash db :"+existingUser.Password+",hash user:"+user.Password)
		return
	}

	sign := jwt.New(jwt.SigningMethodHS256)
	token, err := sign.SignedString([]byte("secret"))

	if err != nil {
		util.Response400(context, "Gagal login, silakan kontak admin.", "")
		return
	}

	context.JSON(http.StatusOK, gin.H{
		"code":    http.StatusOK,
		"message": http.StatusText(http.StatusOK),
		"token":   token,
	})
}
