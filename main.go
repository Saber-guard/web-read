package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
	"os"
	"path/filepath"
	"web-read/middleware"
	"web-read/route"
	"web-read/service"
)

func main() {
	// 加载配置文件
	_ = godotenv.Load()

	// 获取项目根目录
	dir, err := filepath.Abs(filepath.Dir(os.Args[0]))
	if err != nil {
		fmt.Println("根目录获取失败")
		fmt.Println(err)
		os.Exit(0)
	}
	_ = os.Setenv("ROOT_DIR", dir)

	router := gin.Default()
	// 中间件
	router.Use(middleware.Cors())                           // 允许跨域
	router.Use(middleware.RequestLog())                     // 记录请求日志
	service.LogService.Log = service.LogService.LogRegist() // 记录逻辑日志
	// 连接数据库
	service.DbService.InitDb()
	// 加载路由
	route.LoadRoute(router)

	// 循环获取微信access_token
	go service.WechatService{}.AccessToken()

	_ = router.Run(":2222")
}
