package route

import (
	"github.com/gin-gonic/gin"
	"web-read/controller"
)

func LoadRoute(route *gin.Engine) *gin.Engine {
	rootGroup := route.Group("")
	{
		rootGroup.GET("/", func(context *gin.Context) {
			context.JSON(200, gin.H{
				"message": "hello",
			})
		})
	}

	// 微信
	wechatGroup := route.Group("/wechat")
	{
		wechatGroup.GET("/callback", controller.WechatController{}.Init)
		wechatGroup.POST("/callback", controller.WechatController{}.Callback)
	}

	// 爬虫
	crawlGroup := route.Group("/crawl")
	{
		crawlGroup.GET("/crawlCompany", controller.CrawlInvestmentController{}.CrawlCompanyList)
		crawlGroup.GET("/crawlCompany/:code", controller.CrawlInvestmentController{}.CrawlCompany)
	}
	return route

}
