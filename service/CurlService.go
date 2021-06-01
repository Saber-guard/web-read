package service

import (
	"io/ioutil"
	"net/http"
)

type CurlService struct {
}

func (c CurlService) GetJson(url string) (result JsonResult, err error) {
	res, err := http.Get(url)
	if err == nil {
		result.code = res.StatusCode
		body, err := ioutil.ReadAll(res.Body)
		if err == nil {
			result.json = string(body)
		}
	}
	return
}

type JsonResult struct {
	code int
	json string
}
