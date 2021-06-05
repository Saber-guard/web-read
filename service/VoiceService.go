package service

import (
	"encoding/json"
	"os"
	"regexp"
	"web-read/channel/voiceChannel"
	"web-read/request/voiceRequest"
	"web-read/request/wechatRequest"
	"web-read/response/voiceResponse"
)

type VoiceService struct {
}

func (v VoiceService) UrlToVoiceListener() {
	for c := range voiceChannel.UrlToVoiceChan {
		// 生成声音
		voiceUrlPrefix := "http://api.codingwork.cn/voices/"
		fileName, err := VoiceService{}.urlToVoice(c.Url)
		if err == nil {
			// 通过客服消息发送给用户
			url := "https://api.weixin.qq.com/cgi-bin/message/custom/send?access_token=" + os.Getenv("WECHAT_ACCESS_TOKEN")
			var textJson = wechatRequest.TextJsonRequest{
				ToUser:  c.ToUserName,
				MsgType: "text",
				Text:    wechatRequest.TextJson{Content: "声音链接：" + voiceUrlPrefix + fileName},
			}
			textJsonBytes, _ := json.Marshal(textJson)
			_, _ = CurlService{}.PostJson(url, textJsonBytes)
		}
	}

}

func (v VoiceService) urlToVoice(url string) (fileName string, err error) {
	res, err := CurlService{}.Get(url)
	if err == nil && res.code == 200 {
		re, _ := regexp.Compile("<head>(.*\n*)+</head>")
		tmp := re.ReplaceAllString(res.text, " ")
		re, _ = regexp.Compile("<!DOCTYPE html>")
		tmp = re.ReplaceAllString(tmp, "")
		re, _ = regexp.Compile("(?U)<script[^>]*>(.*\n*)+</script>")
		tmp = re.ReplaceAllString(tmp, "")
		re, _ = regexp.Compile("(?U)<style[^>]*>(.*\n*)+</style>")
		tmp = re.ReplaceAllString(tmp, "")
		re, _ = regexp.Compile("(?U)<[a-zA-Z0-9]+[^>]*>")
		tmp = re.ReplaceAllString(tmp, " ")
		re, _ = regexp.Compile("(?U)</[a-zA-Z0-9]+>")
		tmp = re.ReplaceAllString(tmp, " ")
		re, _ = regexp.Compile("<!--[^>]+-->")
		tmp = re.ReplaceAllString(tmp, " ")
		re, _ = regexp.Compile("\n")
		tmp = re.ReplaceAllString(tmp, " ")
		re, _ = regexp.Compile(" +")
		tmp = re.ReplaceAllString(tmp, " ")

		// 调用接口转语音
		strToVoiceUrl := "http://api.codingwork.cn/voice/strToVoice"
		dataBytes, _ := json.Marshal(voiceRequest.StrToVoiceData{Text: tmp})
		jsonBytes, _ := json.Marshal(voiceRequest.StrToVoiceRequest{
			Sig:  "ceshisigceshisig",
			Data: string(dataBytes),
		})
		res, err := CurlService{}.PostJson(strToVoiceUrl, jsonBytes)
		if err == nil {
			var response voiceResponse.StrToVoiceResponse
			if err = json.Unmarshal([]byte(res.json), &response); err == nil {
				fileName = response.Data.FileName
			}
		}
	}
	return
}
