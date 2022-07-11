package util

import (
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
)

func Response200(context *gin.Context, data any, message string) {
	if len(message) > 0 {
		context.JSON(http.StatusOK, gin.H{
			"status":  http.StatusText(http.StatusOK),
			"message": message,
			"data":    data,
		})
	}

	context.JSON(http.StatusOK, gin.H{
		"status": http.StatusText(http.StatusOK),
		"data":   data,
	})

}

func Response(context *gin.Context, httpStatusNumber int, message string) {
	context.JSON(httpStatusNumber, gin.H{
		"status":  http.StatusText(httpStatusNumber),
		"message": message,
	})
}

func Response200WithMessage(context *gin.Context, message string) {
	context.JSON(http.StatusOK, gin.H{
		"status":  http.StatusText(http.StatusOK),
		"message": message,
	})
}

func Response400(context *gin.Context, message string, err string) {
	if len(err) > 0 {
		log.Fatal(message + err)
	} else {
		log.Fatal(message)
	}

	context.JSON(http.StatusBadRequest, gin.H{
		"status":  http.StatusText(http.StatusBadRequest),
		"message": message,
	})
}
