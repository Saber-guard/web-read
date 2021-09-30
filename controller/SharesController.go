package controller

import (
	"github.com/gin-gonic/gin"
	"web-read/service"
)

type SharesController struct {
	BaseController
}

func (c SharesController) List(context *gin.Context) {
	data := service.SharesService{}.List()
	context.JSON(200, map[string]interface{}{
		"code":    0,
		"message": "成功",
		"data":    data,
	})
}
