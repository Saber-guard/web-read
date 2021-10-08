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
		"success": true,
		"content": data,
	})
}
