package controller

import (
	"github.com/gin-gonic/gin"
	"web-read/service"
)

type CrawlInvestmentController struct {
	BaseController
}

func (c CrawlInvestmentController) CrawlCompanyList(context *gin.Context) {
	count := service.CrawlService{}.CrawlCompanyList()
	context.JSON(200, map[string]interface{}{
		"count": count,
	})
}

func (c CrawlInvestmentController) CrawlCompany(context *gin.Context) {
	code := context.Param("code")
	count := service.CrawlService{}.CrawlCompany(code)
	context.JSON(200, map[string]interface{}{
		"count": count,
	})
}
