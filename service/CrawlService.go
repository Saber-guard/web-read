package service

import (
	"encoding/json"
	"fmt"
	"math"
	"reflect"
	"strconv"
	"strings"
	"time"
	"web-read/enum"
	"web-read/model"
	"web-read/response/crawlResponse"
	"web-read/util"
)

type CrawlService struct {
}

func (c CrawlService) CrawlCompanyList() int {
	count := 0
	page := 1
	size := 20
	totalPage := 100
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
		var response crawlResponse.CrawlCompanyListResponse
		if err = json.Unmarshal([]byte(res.text), &response); err != nil {
			fmt.Println(err)
			LogService.Log("ERROR", "上市公司返回信息解析失败", LogData{"text": res.text, "error": err})
			continue
		}
		// 第一次循环时重置一下总页数
		if totalPage == 100 {
			totalPage = int(math.Ceil(float64(response.Data.Total) / float64(size)))
			LogService.Log("INFO", "总页数重置", LogData{"totalPage": totalPage})
		}
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
					Date: time.Now().Format(enum.DataZone),
					Code: item["f12"].(string),
					Ext:  string(itemStr),
				}
				dayData.TodayStart, _ = util.StringUtil{}.AllToStr(item["f17"])
				dayData.TodayEnd, _ = util.StringUtil{}.AllToStr(item["f2"])
				dayData.YestdayEnd, _ = util.StringUtil{}.AllToStr(item["f18"])
				dayData.Highest, _ = util.StringUtil{}.AllToStr(item["f15"])
				dayData.Minimum, _ = util.StringUtil{}.AllToStr(item["f16"])
				dayData.DealNum, _ = util.StringUtil{}.AllToStr(item["f5"])
				dayData.DealMoney, _ = util.StringUtil{}.AllToStr(item["f6"])
				dayData.IncreaseRange, _ = util.StringUtil{}.AllToStr(item["f3"])
				dayData.IncreaseMoney, _ = util.StringUtil{}.AllToStr(item["f4"])
				dayData.TurnoverRate, _ = util.StringUtil{}.AllToStr(item["f8"])
				dayData.ProfitRatio, _ = util.StringUtil{}.AllToStr(item["f9"])
				dayData.ValueRatio, _ = util.StringUtil{}.AllToStr(item["f23"])
				_ = DbService.Db.Create(&dayData)
			} else {
				dayData.Date = time.Now().Format(enum.DataZone)
				dayData.Code, _ = util.StringUtil{}.AllToStr(item["f12"])
				dayData.TodayStart, _ = util.StringUtil{}.AllToStr(item["f17"])
				dayData.TodayEnd, _ = util.StringUtil{}.AllToStr(item["f2"])
				dayData.YestdayEnd, _ = util.StringUtil{}.AllToStr(item["f18"])
				dayData.Highest, _ = util.StringUtil{}.AllToStr(item["f15"])
				dayData.Minimum, _ = util.StringUtil{}.AllToStr(item["f16"])
				dayData.DealNum, _ = util.StringUtil{}.AllToStr(item["f5"])
				dayData.DealMoney, _ = util.StringUtil{}.AllToStr(item["f6"])
				dayData.IncreaseRange, _ = util.StringUtil{}.AllToStr(item["f3"])
				dayData.IncreaseMoney, _ = util.StringUtil{}.AllToStr(item["f4"])
				dayData.TurnoverRate, _ = util.StringUtil{}.AllToStr(item["f8"])
				dayData.ProfitRatio, _ = util.StringUtil{}.AllToStr(item["f9"])
				dayData.ValueRatio, _ = util.StringUtil{}.AllToStr(item["f23"])
				dayData.Ext = string(itemStr)
				DbService.Db.Save(&dayData)
			}

			count++
		}
	}
	return count
}

func (c CrawlService) CrawlCompany(code string) bool {
	start := time.Now().AddDate(0, 0, -31).Format(enum.DataZone2)
	end := time.Now().Format(enum.DataZone2)
	url := "http://push2his.eastmoney.com/api/qt/stock/kline/get?" +
		"secid=" + code + "&fields1=f1,f2,f3&fields2=f51,f52,f53,f54,f55,f56,f57,f58,f59,f60,f61" +
		"&klt=101&fqt=0&beg=" + start + "&end=" + end
	res, err := CurlService{}.Get(url)
	if err != nil {
		fmt.Println(err)
		LogService.Log("ERROR", "抓取上市公司信息失败", LogData{"url": url, "error": err})
		return false
	}
	var response crawlResponse.CrawlCompanyResponse
	if err = json.Unmarshal([]byte(res.text), &response); err != nil {
		fmt.Println(err)
		LogService.Log("ERROR", "上市公司返回信息解析失败", LogData{"text": res.text, "error": err})
		return false
	}
	yestdayEnd := "" // 昨收
	for index, item := range response.Data.Klines {
		tmp := strings.Split(item, ",")
		// 第一条只取收盘价,记作昨收
		if index > 0 {
			// 保存日线信息
			var dayData model.DayDataModel
			DbService.Db.Where("code = ?", response.Data.Code).Where("date = ?", tmp[0]).First(&dayData)
			// 判断是否是类型的零值
			if reflect.DeepEqual(dayData, reflect.Zero(reflect.TypeOf(dayData)).Interface()) {
				dayData = model.DayDataModel{
					Date: tmp[0],
					Code: response.Data.Code,
					Ext:  res.text,
				}
				dayData.TodayStart, _ = util.StringUtil{}.AllToStr(tmp[1])
				dayData.TodayEnd, _ = util.StringUtil{}.AllToStr(tmp[2])
				dayData.YestdayEnd, _ = util.StringUtil{}.AllToStr(yestdayEnd)
				dayData.Highest, _ = util.StringUtil{}.AllToStr(tmp[3])
				dayData.Minimum, _ = util.StringUtil{}.AllToStr(tmp[4])
				dayData.DealNum, _ = util.StringUtil{}.AllToStr(tmp[5])
				dayData.DealMoney, _ = util.StringUtil{}.AllToStr(tmp[6])
				dayData.IncreaseRange, _ = util.StringUtil{}.AllToStr(tmp[8])
				dayData.IncreaseMoney, _ = util.StringUtil{}.AllToStr(tmp[9])
				dayData.TurnoverRate, _ = util.StringUtil{}.AllToStr(tmp[10])
				dayData.ProfitRatio = "0"
				dayData.ValueRatio = "0"
				_ = DbService.Db.Create(&dayData)
			} else {
				dayData.Date = tmp[0]
				dayData.Code = response.Data.Code
				dayData.TodayStart, _ = util.StringUtil{}.AllToStr(tmp[1])
				dayData.TodayEnd, _ = util.StringUtil{}.AllToStr(tmp[2])
				dayData.YestdayEnd, _ = util.StringUtil{}.AllToStr(yestdayEnd)
				dayData.Highest, _ = util.StringUtil{}.AllToStr(tmp[3])
				dayData.Minimum, _ = util.StringUtil{}.AllToStr(tmp[4])
				dayData.DealNum, _ = util.StringUtil{}.AllToStr(tmp[5])
				dayData.DealMoney, _ = util.StringUtil{}.AllToStr(tmp[6])
				dayData.IncreaseRange, _ = util.StringUtil{}.AllToStr(tmp[8])
				dayData.IncreaseMoney, _ = util.StringUtil{}.AllToStr(tmp[9])
				dayData.TurnoverRate, _ = util.StringUtil{}.AllToStr(tmp[10])
				//dayData.ProfitRatio = "0"
				//dayData.ValueRatio = "0"
				//dayData.Ext = res.text
				DbService.Db.Save(&dayData)
			}
		}
		yestdayEnd = tmp[2]
	}

	return true
}
