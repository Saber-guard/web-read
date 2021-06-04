package wechatResponse

import (
	"encoding/xml"
)

type TextXmlResponse struct {
	XMLName      xml.Name `xml:"xml"`
	ToUserName   string
	FromUserName string
	CreateTime   int64
	MsgType      string
	Content      string
}
