package wechatRequest

type TextJsonRequest struct {
	ToUser  string   `json:"touser"`
	MsgType string   `json:"msgtype"`
	Text    TextJson `json:"text"`
}

type TextJson struct {
	Content string `json:"content"`
}
