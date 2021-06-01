package main

import (
	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
	"web-read/route"
	"web-read/service"
)

func main() {
	// 加载配置文件
	_ = godotenv.Load()

	router := gin.Default()
	// 中间件
	//router.Use(middleware.LogToFile())// 记录日志
	// 加载路由
	route.LoadRoute(router)

	// 循环获取微信access_token
	go service.WechatService{}.AccessToken()

	_ = router.Run(":2222")
}
