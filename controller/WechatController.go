package controller

import (
	"encoding/xml"
	"github.com/gin-gonic/gin"
	"web-read/request/wechat"
	"web-read/service"
)

type WechatController struct {
	BaseController
}

func (w WechatController) Init(context *gin.Context) {
	var inputs wechat.InitRequest
	if err := context.ShouldBind(&inputs); err != nil {
		w.ApiError(context, -1, "参数错误")
	} else {
		context.String(200, inputs.EchoStr)
	}
}

func (w WechatController) Callback(context *gin.Context) {
	var inputs wechat.BaseXmlRequest
	body, err := context.GetRawData()
	if err == nil {
		// 先解析成基础xml的struct
		if err := xml.Unmarshal(body, &inputs); err == nil {
			// 再根据MsgType解析成不同的struct
			switch inputs.MsgType {
			case "text":
				var textInputs wechat.TextXmlRequest
				if err := xml.Unmarshal(body, &textInputs); err == nil {
					res, err := service.WechatService{}.ReceiveText(textInputs)
					if err == nil {
						context.XML(200, res)
					}
				}
			}
		}
	}
}
