package controller

import (
	"github.com/gin-gonic/gin"
	"web-read/service"
)

type CrawlInvestmentController struct {
	BaseController
}

func (c CrawlInvestmentController) CrawlCompany(context *gin.Context) {
	count := service.CrawlService{}.CrawlCompany()
	context.JSON(200, map[string]interface{}{
		"count": count,
	})
}
