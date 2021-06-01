package wechat

type TextXmlRequest struct {
	BaseXmlRequest
	Content string `xml:"Content"`
}
