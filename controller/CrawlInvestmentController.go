package controller

import (
	"github.com/gin-gonic/gin"
	"web-read/service"
)

type CrawlInvestmentController struct {
	BaseController
}

func (c CrawlInvestmentController) CrawlCompany(context *gin.Context) {
	service.CrawlService{}.CrawlCompany()
	context.String(200, "123")
}
