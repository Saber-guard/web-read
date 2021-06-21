package service

import (
	"encoding/json"
	"fmt"
	"math"
	"reflect"
	"strconv"
	"time"
	"web-read/enum"
	"web-read/model"
	"web-read/response/crawlResponse"
)

type CrawlService struct {
}

func (c CrawlService) CrawlCompany() bool {
	page := 1
	size := 20
	totalPage := 200
	date := time.Now().Format(enum.DataZone)
	fields := "f1,f2,f3,f4,f5,f6,f7,f8,f9,f10,f12,f11,f13,f14,f15,f16,f17,f18,f19,f20,f21,f22,f23,f24,f25," +
		"f26,f27,f28,f29,f30,f31,f32,f33,f34,f35,f36,f37,f38,f39,f40,f41,f42,f43,f44,f45,f62,f128,f136,f115,f152"
	for page <= totalPage {
		timestamp := time.Now().Unix() * 1000
		url := "http://82.push2.eastmoney.com/api/qt/clist/get?" +
			"pn=" + strconv.Itoa(page) + "&pz=" + strconv.Itoa(size) + "&po=1&np=1&ut=bd1d9ddb04089700cf9c27f6f7426281&fltt=2&invt=2&fid=f3&" +
			"fs=m:0+t:6,m:0+t:80,m:1+t:2,m:1+t:23&fields=" + fields + "&_=" + strconv.FormatInt(timestamp, 10)
		page++
		res, err := CurlService{}.Get(url)
		if err != nil {
			fmt.Println(err)
			LogService.Log("ERROR", "抓取上市公司信息失败", LogData{"url": url, "error": err})
			continue
		}
		var response crawlResponse.CrawlCompanyResponse
		if err = json.Unmarshal([]byte(res.text), &response); err != nil {
			fmt.Println(err)
			LogService.Log("ERROR", "上市公司返回信息解析失败", LogData{"text": res.text, "error": err})
			continue
		}
		totalPage = int(math.Ceil(float64(response.Data.Total) / float64(size)))
		for _, item := range response.Data.Diff {
			itemStr, _ := json.Marshal(item)

			var company model.CompanyInfoModel
			DbService.Db.Where("code = ?", item["f12"].(string)).First(&company)
			// 保存公司信息
			// 判断是否是类型的零值
			if reflect.DeepEqual(company, reflect.Zero(reflect.TypeOf(company)).Interface()) {
				company = model.CompanyInfoModel{
					Code: item["f12"].(string),
					Name: item["f14"].(string),
				}
				_ = DbService.Db.Create(&company)
			} else {
				company.Name = item["f14"].(string)
				DbService.Db.Save(&company)
			}

			// 保存当天日线信息
			var dayData model.DayDataModel
			DbService.Db.Where("code = ?", item["f12"].(string)).Where("date = ?", date).First(&dayData)
			// 判断是否是类型的零值
			if reflect.DeepEqual(dayData, reflect.Zero(reflect.TypeOf(dayData)).Interface()) {
				dayData = model.DayDataModel{
					Date:          time.Now().Format(enum.DataZone),
					Code:          item["f12"].(string),
					TodayStart:    item["f17"].(float32),
					YestdayEnd:    item["f18"].(float32),
					Highest:       item["f15"].(float32),
					Minimum:       item["f16"].(float32),
					DealNum:       item["f5"].(int32),
					DealMoney:     item["f6"].(float64),
					IncreaseRange: item["f3"].(float32),
					IncreaseMoney: item["f4"].(float32),
					TurnoverRate:  item["f8"].(float32),
					ProfitRatio:   item["f9"].(float32),
					ValueRatio:    item["f23"].(float32),
					Ext:           string(itemStr),
				}
				_ = DbService.Db.Create(&dayData)
			} else {
				dayData.Date = time.Now().Format(enum.DataZone)
				dayData.Code = item["f12"].(string)
				dayData.TodayStart = item["f17"].(float32)
				dayData.YestdayEnd = item["f18"].(float32)
				dayData.Highest = item["f15"].(float32)
				dayData.Minimum = item["f16"].(float32)
				dayData.DealNum = item["f5"].(int32)
				dayData.DealMoney = item["f6"].(float64)
				dayData.IncreaseRange = item["f3"].(float32)
				dayData.IncreaseMoney = item["f4"].(float32)
				dayData.TurnoverRate = item["f8"].(float32)
				dayData.ProfitRatio = item["f9"].(float32)
				dayData.ValueRatio = item["f23"].(float32)
				dayData.Ext = string(itemStr)
				DbService.Db.Save(&dayData)
			}

			break // todo::调试代码
		}

		break // todo::调试代码
	}
	return true
}
