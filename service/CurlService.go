package service

import (
	"bytes"
	"io/ioutil"
	"net/http"
)

type CurlService struct {
}

func (c CurlService) Get(url string) (result TextResult, err error) {
	res, err := http.Get(url)
	if err == nil {
		result.code = res.StatusCode
		body, err := ioutil.ReadAll(res.Body)
		if err == nil {
			result.text = string(body)
		}
	}
	return
}

type TextResult struct {
	code int
	text string
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

func (c CurlService) PostJson(url string, jsonBytes []byte) (result JsonResult, err error) {
	res, err := http.Post(url, "application/json", bytes.NewBuffer(jsonBytes))
	if err == nil {
		result.code = res.StatusCode
		body, err := ioutil.ReadAll(res.Body)
		if err == nil {
			result.json = string(body)
		}
	}
	return
}
