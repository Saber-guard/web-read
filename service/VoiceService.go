package service

import (
	"bytes"
	"crypto/md5"
	"encoding/base64"
	"fmt"
	"github.com/PuerkitoBio/goquery"
	htmlQuery "github.com/antchfx/xquery/html"
	"github.com/faiface/beep"
	"github.com/faiface/beep/wav"
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
		re, err := regexp.Compile("^http(s)?://mp\\.weixin\\.qq\\.com")
		var text string
		if err == nil && re.MatchString(url) {
			text = v.WechatHtmlToText(res.text)
		} else {
			text = v.HtmlToText(res.text)
		}

		// 文字内容md5生成文件名称
		m := md5.New()
		_, _ = io.WriteString(m, text)
		fileName = fmt.Sprintf("%x", m.Sum(nil)) + ".mp3"
		filePath := os.Getenv("VOICE_PATH") + fileName

		// todo::为了避免并发请求导致重复调用腾讯云，需要一个分布式锁

		// 判断文件是否存在，存在则直接返回文件名
		fileInfo, fileExistErr := os.Stat(filePath)
		// err==nil文件存在，单err!=nil不一定是报文件不存在的错误，需要os.IsNotExist()判断
		if fileExistErr == nil {
			LogService.Log("INFO", "os.Stat success,file already exists", LogData{"filePath": filePath})
			return fileInfo.Name(), fileExistErr
		}

		// 文件不存在先创建，避免微信重试机制导致重复转换
		file, err := os.Create(filePath)
		if err != nil {
			LogService.Log("ERROR", "os.Create error", LogData{
				"error": err, "filePath": filePath,
			})
			return "", err
		}
		defer func() {
			if err = file.Close(); err != nil {
				LogService.Log("ERROR", "os.Close error", LogData{
					"error": err, "filePath": filePath,
				})
			}
		}()

		// 生成音频
		var (
			textArr   = util.StringUtil{}.SplitByLen(strings.TrimSpace(text), 105)
			voiceArr  = make([]beep.Streamer, len(textArr))
			formatArr = make([]beep.Format, len(textArr))
			wg        sync.WaitGroup
		)
		for index, item := range textArr {
			wg.Add(1)
			go v.TextToVoice(index, strings.TrimSpace(item), fileName, voiceArr, formatArr, &wg)
			time.Sleep(time.Millisecond * 85)
		}
		wg.Wait()

		if err = v.writeToFile(file, textArr, voiceArr, formatArr); err != nil {
			return "", err
		}
	}

	LogService.Log("INFO", "~~~urlToVoice end~~~", LogData{"url": url, "fileName": fileName})
	return
}

func (v VoiceService) writeToFile(
	file *os.File, textArr []string, voiceArr []beep.Streamer, formatArr []beep.Format,
) error {
	streamer := beep.Seq(voiceArr...)

	// 取第一个format值
	var format = beep.Format{}
	if len(textArr) > 0 {
		format = formatArr[0]
	}
	err := wav.Encode(file, streamer, format)
	if err != nil {
		LogService.Log("ERROR", "wav.Encode error", LogData{
			"error": err, "fileName": file.Name(),
		})
		return err
	}
	return nil
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

func (v VoiceService) WechatHtmlToText(html string) (text string) {
	doc, _ := htmlQuery.Parse(strings.NewReader(html))
	// 标题
	title := htmlQuery.FindOne(doc, "//h1[@id='activity-name']/text()")
	text = title.Data
	// 文章内容
	article := htmlQuery.FindOne(doc, "//div[@id='js_content']")
	articleHtml, _ := goquery.NewDocumentFromNode(article).Html()
	text += v.HtmlToText(articleHtml)

	return
}

func (v VoiceService) TextToVoice(
	index int, text string, fileName string, voiceArr []beep.Streamer, formatArr []beep.Format, wg *sync.WaitGroup,
) {
	defer wg.Done()

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
	request.Codec = common.StringPtr("wav")
	// request.Speed = common.Float64Ptr(0) // 语速-2代表0.6倍 -1代表0.8倍 0代表1.0倍（默认） 1代表1.2倍 2代表1.5倍
	res, err := client.TextToVoice(request)
	if err != nil {
		LogService.Log("ERROR", "client.TextToVoice error", LogData{
			"error": err, "text": text, "index": index, "fileName": fileName,
		})
		return
	}

	if res != nil && res.Response != nil && res.Response.Audio != nil {
		b, _ := base64.StdEncoding.DecodeString(*res.Response.Audio)
		streamer, format, err := wav.Decode(bytes.NewReader(b))
		if err != nil {
			LogService.Log("ERROR", "wav.Decode error", LogData{
				"error": err, "text": text, "index": index, "fileName": fileName,
			})
		}
		defer streamer.Close()

		voiceArr[index] = streamer
		formatArr[index] = format
	}
}
