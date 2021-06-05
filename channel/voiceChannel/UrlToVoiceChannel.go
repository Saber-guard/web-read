package voiceChannel

type UrlToVoiceChannel struct {
	Url          string
	FromUserName string
	ToUserName   string
}

var UrlToVoiceChan chan UrlToVoiceChannel
