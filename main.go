package main

import (
	"encoding/json"
	"fmt"
	"github.com/gin-gonic/gin"
	"io/ioutil"
	"log"
	"message/core"
	"message/model"
	"time"
)

type S2 struct {
	Title string
	Content []model.Item
}

type S1 struct {
	Flag int
	Style int
	Data []S2
}

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
	r.Static("bootstrap/img", "./web/static/bootstrap/img")
	r.Static("/static/bootstrap", "./web/static/bootstrap")
	r.LoadHTMLGlob("./web/static/template/**/*")

	// 加载hugo
	r.Static("/public", "./hugo/site/public/")

	r.GET("/", func(c *gin.Context) {
		c.Request.URL.Path = "/weibo"
		r.HandleContext(c)
	})

	r.GET("/hot", func(c *gin.Context) {
		c.Request.URL.Path = "/weibo"
		r.HandleContext(c)
	})

	r.GET("/weibo", func(c *gin.Context) {
		tmp := make(map[string]interface{})
		tmp["Flag"] = 0
		tmp["Data"] = model.M["Weibo"]
		c.HTML(200, "onelist.html",tmp)
	})

	r.GET("/baidu", func(c *gin.Context) {
		tmp := make(map[string]interface{})
		tmp["Flag"] = 0
		tmp["Data"] = model.M["Baidu"]
		c.HTML(200, "onelist.html",tmp)
	})

	r.GET("/cctv", func(c *gin.Context) {
		tmp := make(map[string]interface{})
		tmp["Flag"] = 0
		tmp["Data"] = model.M["CCTV"]
		c.HTML(200, "onelist.html",tmp)
	})

	r.GET("/csdn", func(c *gin.Context) {
		tmp := make(map[string]interface{})
		tmp["Flag"] = 0
		tmp["Data"] = model.M["CSDN"]
		c.HTML(200, "onelist.html",tmp)
	})

	r.GET("/zhihu", func(c *gin.Context) {
		tmp := make(map[string]interface{})
		tmp["Flag"] = 0
		tmp["Data"] = model.M["Zhihu"]
		c.HTML(200, "onelist.html",tmp)
	})

	r.GET("/acfun", func(c *gin.Context) {
		tmp := make(map[string]interface{})
		tmp["Flag"] = 0
		tmp["Data"] = model.M["Acfun"]
		c.HTML(200, "onelist.html",tmp)
	})

	r.GET("/blbl", func(c *gin.Context) {
		tmp := make(map[string]interface{})
		tmp["Flag"] = 0
		tmp["Data"] = model.M["Bilibili"]
		c.HTML(200, "onelist.html", tmp)
	})

	r.GET("/github", func(c *gin.Context) {
		var R S1
		R.Flag = 1
		R.Style = 20
		R.Data = append(R.Data,
			S2{Title: "Daily", Content: model.M["Github3"]},
			S2{Title: "Weekly", Content: model.M["Github4"]},
			S2{Title: "Monthly", Content: model.M["Github5"]},
		)
		c.HTML(200, "threelist.html",R)
	})

	r.GET("/githubGo", func(c *gin.Context) {
		var R S1
		R.Flag = 1
		R.Style = 20
		R.Data = append(R.Data,
			S2{Title: "Daily(Go)", Content: model.M["Github0"]},
			S2{Title: "Weekly(Go)", Content: model.M["Github1"]},
			S2{Title: "Monthly(Go)", Content: model.M["Github2"]},
		)
		c.HTML(200, "threelist.html",R)
	})

	r.GET("/githubPy", func(c *gin.Context) {
		var R S1
		R.Flag = 1
		R.Style = 20
		R.Data = append(R.Data,
			S2{Title: "Daily(Py)", Content: model.M["Github6"]},
			S2{Title: "Weekly(Py)", Content: model.M["Github7"]},
			S2{Title: "Monthly(Py)", Content: model.M["Github8"]},
		)
		c.HTML(200, "threelist.html",R)
	})

	r.GET("/juejin", func(c *gin.Context) {
		var R S1
		R.Flag = 1
		R.Style = 10
		R.Data = append(R.Data,
			S2{Title: "Triduum", Content: model.M["Juejin3"]},
			S2{Title: "Weekly", Content: model.M["Juejin7"]},
			S2{Title: "Monthly", Content: model.M["Juejin30"]},
		)
		c.HTML(200, "threelist.html",R)
	})

	r.GET("/kr36", func(c *gin.Context) {
		var R S1
		R.Flag = 1
		R.Style = 10
		R.Data = append(R.Data,
			S2{Title: "人气榜", Content: model.M["Kr360"]},
			S2{Title: "综合榜", Content: model.M["Kr361"]},
			S2{Title: "收藏榜", Content: model.M["Kr362"]},
		)
		c.HTML(200, "threelist.html",R)
	})

	r.GET("/zwtx", func(c *gin.Context) {
		var R S1
		R.Flag = 1
		R.Style = 10
		R.Data = append(R.Data,
			S2{Title: "06:00", Content: model.M["zwtx0"]},
			S2{Title: "07:00", Content: model.M["zwtx1"]},
			S2{Title: "08:00", Content: model.M["zwtx2"]},
		)
		c.HTML(200, "threelist.html",R)
	})

	r.GET("/tool", func(c *gin.Context) {
		var v model.ToolStruct
		b,_ := ioutil.ReadFile("./data/tool.txt")
		err := json.Unmarshal(b, &v)
		if err != nil{
			log.Printf("tool err:%v\n", err)
			return
		}
		c.HTML(200, "tool.html",v)
	})

	r.Run("127.0.0.1:8080")
}

func PrintM()  {
	fmt.Println("===============================")
	for k,v := range model.M{
		log.Println(k, len(v))
	}
}


