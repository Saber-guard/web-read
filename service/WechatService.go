package service

import (
	"encoding/json"
	"fmt"
	"os"
	"time"
	wechat2 "web-read/request/wechat"
	"web-read/response/wechat"
)

type WechatService struct {
}

// 获取AccessToken
func (w WechatService) AccessToken() {
	url := "https://api.weixin.qq.com/cgi-bin/token?grant_type=client_credential&appid=" + os.Getenv("WECHAT_APPID") + "&secret=" + os.Getenv("WECHAT_SECRET")
	for {
		var response wechat.AccessTokenResponse
		res, err := CurlService{}.GetJson(url)
		if err != nil {
			fmt.Println("AccessToken请求失败")
		}
		if err = json.Unmarshal([]byte(res.json), &response); err != nil {
			fmt.Println("AccessToken解析失败")
		}
		if err = os.Setenv("WECHAT_ACCESS_TOKEN", response.AccessToken); err != nil {
			fmt.Println("AccessToken记录失败")
		}

		// 定时获取
		time.Sleep(time.Minute * 60)
	}
}

// 接收text消息
func (w WechatService) ReceiveText(inputs wechat2.TextXmlRequest) (response wechat.TextXmlResponse, err error) {
	response.FromUserName = inputs.ToUserName
	response.ToUserName = inputs.FromUserName
	response.MsgType = "text"
	response.CreateTime = time.Now().Unix()
	response.Content = inputs.Content
	return
}
