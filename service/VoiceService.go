package service

import (
	"crypto/md5"
	"encoding/base64"
	"fmt"
	"github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/common"
	"github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/common/profile"
	tts "github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/tts/v20190823"
	"io"
	"os"
	"regexp"
	"strings"
	"sync"
	"time"
	"web-read/util"
)

type VoiceService struct {
}

func (v VoiceService) urlToVoice(url string) (fileName string, err error) {
	res, err := CurlService{}.Get(url)
	if err == nil && res.code == 200 {
		// html转换成文本
		text := v.HtmlToText(res.text)

		// 文字内容md5生成文件名称
		m := md5.New()
		_, _ = io.WriteString(m, text)
		fileName = fmt.Sprintf("%x", m.Sum(nil)) + ".mp3"
		filePath := "./tmp/voices/" + fileName

		// 判断文件是否存在，存在则直接返回文件名
		_, fileExistErr := os.Stat(filePath)
		if fileExistErr != nil {
			textArr := util.StringUtil{}.SplitByLen(text, 105)
			voiceArr := make([]string, len(textArr))
			var wg sync.WaitGroup
			for index, item := range textArr {
				wg.Add(1)
				go v.TextToVoice(index, item, fileName, voiceArr, &wg)
				time.Sleep(time.Millisecond * 85)
			}
			wg.Wait()
			// 写入文件
			f, err := os.Create(filePath)
			if err == nil {
				bytes, _ := base64.StdEncoding.DecodeString(strings.Join(voiceArr, ""))
				_, _ = f.Write(bytes)
				_ = f.Close()
			}
		}
	}
	return
}

func (v VoiceService) HtmlToText(html string) (text string) {
	re, _ := regexp.Compile("<head>(.*\n*)+</head>")
	text = re.ReplaceAllString(html, " ")
	re, _ = regexp.Compile("<!DOCTYPE html>")
	text = re.ReplaceAllString(text, "")
	re, _ = regexp.Compile("(?U)<script[^>]*>(.*\n*)+</script>")
	text = re.ReplaceAllString(text, "")
	re, _ = regexp.Compile("(?U)<style[^>]*>(.*\n*)+</style>")
	text = re.ReplaceAllString(text, "")
	re, _ = regexp.Compile("(?U)<[a-zA-Z0-9]+[^>]*>")
	text = re.ReplaceAllString(text, " ")
	re, _ = regexp.Compile("(?U)</[a-zA-Z0-9]+>")
	text = re.ReplaceAllString(text, " ")
	re, _ = regexp.Compile("<!--[^>]+-->")
	text = re.ReplaceAllString(text, "")
	re, _ = regexp.Compile("[\n\t\f]")
	text = re.ReplaceAllString(text, " ")
	re, _ = regexp.Compile(" +")
	text = re.ReplaceAllString(text, " ")
	re, _ = regexp.Compile("(&lt;)|(&gt;)|(&nbsp;)|(&emsp;)|(&ensp;)|(&quot;)")
	text = re.ReplaceAllString(text, "")
	return
}

func (v VoiceService) TextToVoice(index int, text string, fileName string, voiceArr []string, wg *sync.WaitGroup) {
	// 调用腾讯云接口转语音
	credential := common.NewCredential(
		os.Getenv("SECRET_ID"),
		os.Getenv("SECRET_KET"),
	)
	cpf := profile.NewClientProfile()
	cpf.HttpProfile.Endpoint = "tts.tencentcloudapi.com"
	client, _ := tts.NewClient(credential, "ap-chengdu", cpf)
	request := tts.NewTextToVoiceRequest()
	request.Text = common.StringPtr(text)
	request.SessionId = common.StringPtr(fileName)
	request.ModelType = common.Int64Ptr(1)
	request.VoiceType = common.Int64Ptr(4) // 选择声音
	request.Codec = common.StringPtr("mp3")
	request.Speed = common.Float64Ptr(1) // 语速-2代表0.6倍 -1代表0.8倍 0代表1.0倍（默认） 1代表1.2倍 2代表1.5倍
	res, err := client.TextToVoice(request)
	if err == nil {
		voiceArr[index] = *res.Response.Audio
	}
	wg.Done()
}
