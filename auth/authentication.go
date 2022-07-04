package auth

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v4"
	"net/http"
	"tahfidz-backend/model"
)

func auth(context *gin.Context) {
	tokenString := context.Request.Header.Get("Authorization")
	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		if jwt.GetSigningMethod("HS256") != token.Method {
			return nil, fmt.Errorf("Unexpected signing method: %v", token.Header["alg"])
		}

		return []byte("secret"), nil
	})

	if token != nil && err == nil {
		fmt.Println("token verified")
	} else {
		result := gin.H{
			"message": "You are not authorized",
			"error":   err.Error(),
		}
		context.JSON(http.StatusUnauthorized, result)
		context.Abort()
	}
}

func Login(context *gin.Context) {
	var user model.User
	err := context.Bind(&user)
	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{
			"status":  http.StatusBadRequest,
			"message": "can't bind struct",
		})
	}
	if user.Username != "admin" && user.Password != "admin" {
		context.JSON(http.StatusUnauthorized, gin.H{
			"status":  http.StatusUnauthorized,
			"message": "wrong username or password",
		})
	}

	sign := jwt.New(jwt.SigningMethodHS256)
	token, err := sign.SignedString([]byte("secret"))

	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{
			"code":    http.StatusBadRequest,
			"message": "'username' and 'password' combination is wrong.",
			"status":  http.StatusText(http.StatusBadRequest),
		})
	}

	context.JSON(http.StatusOK, gin.H{
		"code":    http.StatusOK,
		"message": http.StatusText(http.StatusOK),
		"token":   token,
	})
}
