package wechatRequest

type InitRequest struct {
	Signature string `form:"signature" binding:"required"`
	Timestamp string `form:"timestamp" binding:"required"`
	Nonce     string `form:"nonce" binding:"required"`
	EchoStr   string `form:"echostr" binding:"required"`
}
