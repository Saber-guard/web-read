package voiceRequest

type StrToVoiceRequest struct {
	Sig  string `json:"sig"`
	Data string `json:"data"`
}

type StrToVoiceData struct {
	Text string `json:"text"`
}
