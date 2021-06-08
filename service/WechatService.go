package service

import (
	"encoding/json"
	"fmt"
	"os"
	"regexp"
	"time"
	"web-read/request/wechatRequest"
	"web-read/response/wechatResponse"
)

type WechatService struct {
}

// 获取AccessToken
func (w WechatService) AccessToken() {
	url := "https://api.weixin.qq.com/cgi-bin/token?grant_type=client_credential&appid=" + os.Getenv("WECHAT_APPID") + "&secret=" + os.Getenv("WECHAT_SECRET")
	for {
		var response wechatResponse.AccessTokenResponse
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
func (w WechatService) ReceiveText(inputs wechatRequest.TextXmlRequest) (response wechatResponse.TextXmlResponse, err error) {

	response.FromUserName = inputs.ToUserName
	response.ToUserName = inputs.FromUserName
	response.CreateTime = time.Now().Unix()

	// 如果是http或https开头，调用在线语音合成
	re, err := regexp.Compile("^http(s)?://")
	if err == nil && re.MatchString(inputs.Content) {
		voiceUrlPrefix := "http://voice.codingwork.cn:8080/"
		fileName, err := VoiceService{}.urlToVoice(inputs.Content)
		if err == nil {
			response.MsgType = "text"
			response.Content = "声音链接：" + voiceUrlPrefix + fileName
		}
	} else {
		response.MsgType = "text"
		response.Content = inputs.Content
	}

	return
}
