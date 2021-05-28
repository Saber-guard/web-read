package controller

import "github.com/gin-gonic/gin"

type BaseController struct {
}

func (b BaseController) ApiSucc(context *gin.Context, data interface{}, message string) {
	context.JSON(200, gin.H{
		"code":    0,
		"message": message,
		"data":    data,
	})
}

func (b BaseController) ApiError(context *gin.Context, code int, message string) {
	context.JSON(200, gin.H{
		"code":    code,
		"message": message,
	})
}
