package util

import (
	"github.com/sirupsen/logrus"
	"net/http"

	"github.com/gin-gonic/gin"
)

func Response200(context *gin.Context, data any, message string) {
	if len(message) > 0 {
		context.JSON(http.StatusOK, gin.H{
			"code":    http.StatusOK,
			"status":  http.StatusText(http.StatusOK),
			"message": message,
			"data":    data,
		})
		return
	}

	context.JSON(http.StatusOK, gin.H{
		"code":   http.StatusOK,
		"status": http.StatusText(http.StatusOK),
		"data":   data,
	})

}

func Response(context *gin.Context, httpStatusNumber int, message string) {
	context.JSON(httpStatusNumber, gin.H{
		"code":    httpStatusNumber,
		"status":  http.StatusText(httpStatusNumber),
		"message": message,
	})
}

func Response400(context *gin.Context, message string, err string) {
	if len(err) > 0 {
		logrus.Error(message + " " + err)
	} else {
		logrus.Error(message)
	}

	context.JSON(http.StatusBadRequest, gin.H{
		"code":    http.StatusBadRequest,
		"status":  http.StatusText(http.StatusBadRequest),
		"message": message,
	})
}
