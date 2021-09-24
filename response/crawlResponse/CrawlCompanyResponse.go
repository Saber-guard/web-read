package crawlResponse

type CrawlCompanyResponse struct {
	Data struct {
		Code   string   `json:"code"`
		Klines []string `json:"klines"`
	} `json:"data"`
}
