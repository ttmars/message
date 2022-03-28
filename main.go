package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"log"
	"message/core"
	"message/model"
	"time"
)


func main()  {
	Run()
}

func Run(){
	// 采集GitHub
	go func() {
		for {
			core.Github()
			time.Sleep(time.Second*120)
		}
	}()
	// 其他页面
	go func() {
		for {
			core.GetAll()
			core.Juejin()
			core.Kr36()
			PrintM()
			time.Sleep(time.Second*60)
		}
	}()
	// 开启服务
	RunServer()
}

func RunServer() {
	r := gin.Default()

	// 加载静态文件和模板
	r.Static("/static/bootstrap", "./web/static/bootstrap")
	r.LoadHTMLGlob("./web/static/template/*")

	// 加载hugo
	r.Static("/public", "./hugo/site/public/")

	r.GET("/", func(c *gin.Context) {
		tmp := make(map[string]interface{})
		tmp["Flag"] = 0
		tmp["Data"] = model.M["Weibo"]
		c.HTML(200, "content.html",tmp)
	})

	r.GET("/hot", func(c *gin.Context) {
		tmp := make(map[string]interface{})
		tmp["Flag"] = 0
		tmp["Data"] = model.M["Weibo"]
		c.HTML(200, "content.html",tmp)
	})

	r.GET("/csdn", func(c *gin.Context) {
		tmp := make(map[string]interface{})
		tmp["Flag"] = 0
		tmp["Data"] = model.M["CSDN"]
		c.HTML(200, "content.html",tmp)
	})

	r.GET("/weibo", func(c *gin.Context) {
		tmp := make(map[string]interface{})
		tmp["Flag"] = 0
		tmp["Data"] = model.M["Weibo"]
		c.HTML(200, "content.html",tmp)
	})

	r.GET("/baidu", func(c *gin.Context) {
		tmp := make(map[string]interface{})
		tmp["Flag"] = 0
		tmp["Data"] = model.M["Baidu"]
		c.HTML(200, "content.html",tmp)
	})

	r.GET("/acfun", func(c *gin.Context) {
		tmp := make(map[string]interface{})
		tmp["Flag"] = 0
		tmp["Data"] = model.M["Acfun"]
		c.HTML(200, "content.html",tmp)
	})

	r.GET("/blbl", func(c *gin.Context) {
		tmp := make(map[string]interface{})
		tmp["Flag"] = 0
		tmp["Data"] = model.M["Bilibili"]
		c.HTML(200, "content.html", tmp)
	})

	r.GET("/github", func(c *gin.Context) {
		tmp := make(map[string]interface{})
		tmp["Flag"] = 1
		tmp["Data"] = model.M
		c.HTML(200, "github.html",tmp)
	})

	r.GET("/githubGo", func(c *gin.Context) {
		tmp := make(map[string]interface{})
		tmp["Flag"] = 2
		tmp["Data"] = model.M
		c.HTML(200, "githubGo.html", tmp)
	})

	r.GET("/juejin", func(c *gin.Context) {
		tmp := make(map[string]interface{})
		tmp["Flag"] = 3
		tmp["Data"] = model.M
		c.HTML(200, "juejin.html",tmp)
	})

	r.GET("/kr36", func(c *gin.Context) {
		tmp := make(map[string]interface{})
		tmp["Flag"] = 4
		tmp["Data"] = model.M
		c.HTML(200, "kr36.html",tmp)
	})

	r.Run("127.0.0.1:8080")
}

func PrintM()  {
	fmt.Println("===============================")
	for k,v := range model.M{
		log.Println(k, len(v))
	}
}
