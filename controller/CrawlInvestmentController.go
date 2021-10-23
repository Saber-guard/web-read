package controller

import (
	"github.com/gin-gonic/gin"
	"web-read/request/crawlRequest"
	"web-read/service"
)

type CrawlInvestmentController struct {
	BaseController
}

func (c CrawlInvestmentController) CrawlCompanyList(context *gin.Context) {
	var inputs crawlRequest.CompanyListRequest
	err := context.ShouldBindQuery(&inputs)
	if err != nil {
		c.ApiError(context, -1, "参数错误")
	} else {
		count := service.CrawlService{}.CrawlCompanyList(inputs)
		context.JSON(200, map[string]interface{}{
			"count": count,
		})
	}
}

func (c CrawlInvestmentController) CrawlCompany(context *gin.Context) {
	code := context.Param("code")
	count := service.CrawlService{}.CrawlCompany(code)
	context.JSON(200, map[string]interface{}{
		"count": count,
	})
}
