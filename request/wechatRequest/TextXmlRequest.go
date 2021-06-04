package wechatRequest

type TextXmlRequest struct {
	BaseXmlRequest
	Content string `xml:"Content"`
}
