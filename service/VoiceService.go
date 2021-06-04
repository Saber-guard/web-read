package service

import (
	"encoding/json"
	"regexp"
	"web-read/request/voiceRequest"
	"web-read/response/voiceResponse"
)

type VoiceService struct {
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
