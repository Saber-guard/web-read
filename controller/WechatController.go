package controller

import (
	"github.com/gin-gonic/gin"
	"web-read/request/wechat"
)

type WechatController struct {
	BaseController
}

func (w WechatController) Init(context *gin.Context) {
	var inputs wechat.InitRequest
	err := context.ShouldBind(&inputs)
	if err != nil {
		w.ApiError(context, -1, "参数错误")
	} else {
		context.String(200, inputs.EchoStr)
	}
}
