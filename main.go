package main

import (
	"github.com/gin-gonic/gin"
	"go-cloud-space/config"
	"go-cloud-space/routes"
)

func main() {
	config.InitDB()
	r := gin.Default()
	r.Static("/static", "./public/static/")
	// 注册模板引擎
	r.LoadHTMLGlob("./resources/views/*")
	//中间件
	//r.Use(Middleware.LoggerMiddleware())
	//路由注册
	routes.Init(r)
	r.Run("0.0.0.0:8000")
}
