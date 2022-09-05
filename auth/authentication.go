package auth

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v4"
	"github.com/sirupsen/logrus"
	"golang.org/x/crypto/bcrypt"
	"net/http"
	"strings"
	"tahfidz-backend/model"
	"tahfidz-backend/repository"
	"tahfidz-backend/util"
	"time"
)

func Auth(context *gin.Context, expectedRoles []string) bool {
	tokenString := context.Request.Header.Get("Authorization")
	if len(tokenString) <= 0 {
		util.Response(context, http.StatusBadRequest,
			"Anda tidak diperbolehkan untuk melihat halaman ini. Silakan hubungi Admin.")
		return false
	}

	token, _ := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		if jwt.GetSigningMethod("HS256") != token.Method {
			return nil, fmt.Errorf("Unexpected signing method: %v", token.Header["alg"])
		}

		return []byte("secret"), nil
	})

	if claims, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {
		var role = claims["Role"].(string)
		logrus.Info("Sukses login dengan username " + claims["Username"].(string) + " dan role " + role)
		if len(expectedRoles) > 0 {
			for _, expectedRole := range expectedRoles {
				if strings.EqualFold(expectedRole, role) {
					return true
				}
			}
		} else {
			return true
		}
	}

	util.Response(context, http.StatusUnauthorized,
		"Anda tidak diperbolehkan untuk masuk. Silakan hubungi Admin.")
	context.Abort()
	return false
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

	var existingUser = repository.FetchUserByUsername(user.Username, true)
	if err := bcrypt.CompareHashAndPassword([]byte(existingUser.Password), []byte(user.Password)); err != nil {
		util.Response400(context, "Username dan/atau Password tidak sesuai.", "hash db :"+existingUser.Password+",hash user:"+user.Password)
		return
	}

	claims := model.TahfidzClaim{
		RegisteredClaims: jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(time.Now().Add(24 * time.Hour)),
		},
		Id:       existingUser.Id,
		Username: existingUser.Username,
		Role:     existingUser.Role,
	}

	sign := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	token, err := sign.SignedString([]byte("secret"))

	if err != nil {
		util.Response400(context, "Gagal login, silakan kontak admin.", "")
		return
	}

	context.JSON(http.StatusOK, gin.H{
		"code":    http.StatusOK,
		"message": http.StatusText(http.StatusOK),
		"token":   token,
		"id":      claims.Id,
	})
}
