package voiceResponse

type StrToVoiceResponse struct {
	Error int            `json:"error"`
	Info  string         `json:"info"`
	Data  StrToVoiceData `json:"data"`
}
type StrToVoiceData struct {
	FileName string `json:"file_name"`
}
