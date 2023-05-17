package routes

import (
	"github.com/gin-gonic/gin"
	"go-cloud-space/ec/egin"
	"go-cloud-space/ec/goefun/ecore"
)

func Init(Router *gin.Engine) {
	Router.GET("/", func(c *gin.Context) {
		//模板渲染
		c.HTML(200, "index.html", gin.H{})
	})
	Router.POST("/data", func(c *gin.Context) {
		type Data struct {
			Content string `i:"content" rule:"required" msg:"内容不能为空"`
		}
		var data = Data{}
		err := egin.Verify(c, &data)
		if err != nil {
			c.HTML(200, "index.html", gin.H{})
			return
		}
		//生成随机字符串作为文件名 将 Content 内容保存到 storage 目录中
		ecore.E置随机数种子(0)

		文件名 := ecore.E文本_取随机字母(4, 1) + ".txt"
		_ = ecore.E文件_保存("./storage/txt/"+文件名, data.Content)

		//地址重定向到 /data/文件名
		c.Redirect(302, "/data/"+文件名)

	})

	Router.GET("/data/:id", func(c *gin.Context) {
		type Data struct {
			Id string `i:"id" rule:"required" msg:"Id 不能为空"`
		}
		var data = Data{}
		err := egin.Verify(c, &data)
		if err != nil {
			c.JSON(200, gin.H{"code": 400, "msg": err.Error()})
			return
		}
		//读取文件内容
		content := ecore.E读入文本("./storage/txt/" + data.Id)
		//url 当前页面的完整url地址
		url := c.Request.URL.String()
		//获取当前访问的域名
		domain := c.Request.Host
		url = "http://" + domain + url

		//渲染模板 show.html
		c.HTML(200, "show.html", gin.H{
			"content": content,
			"id":      data.Id,
			"url":     url,
		})

	})
	Router.GET("/data/:id/delete", func(c *gin.Context) {
		type Data struct {
			Id string `i:"id" rule:"required" msg:"Id 不能为空"`
		}
		var data = Data{}
		err := egin.Verify(c, &data)
		if err != nil {
			c.JSON(200, gin.H{"code": 400, "msg": err.Error()})
			return
		}
		//删除文件
		_ = ecore.E删除文件("./storage/txt/" + data.Id)
		//重定向到首页
		c.Redirect(302, "/")
	})
	Router.POST("/data/:id/update", func(c *gin.Context) {
		type Data struct {
			Id      string `i:"id" rule:"required" msg:"Id 不能为空"`
			Content string `i:"content" rule:"required" msg:"内容不能为空"`
		}
		var data = Data{}
		err := egin.Verify(c, &data)
		if err != nil {
			c.JSON(200, gin.H{"code": 400, "msg": err.Error()})
			return
		}
		//保存文件
		_ = ecore.E文件_保存("./storage/txt/"+data.Id, data.Content)
		//重定向原来的页面
		c.Redirect(302, "/data/"+data.Id)
	})
}
