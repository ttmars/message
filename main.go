package main

import (
	"encoding/json"
	"github.com/gin-gonic/gin"
	"io/ioutil"
	"log"
	"message/core"
	"message/model"
	"time"
)

type S2 struct {
	Title   string
	Content []model.Item
}

type S1 struct {
	Flag  int
	Style int
	Data  []S2
}

func main() {
	Run()
}

func handleErr(err error, msg string) {
	if err != nil {
		log.Printf("%s: %v\n", msg, err)
	}
}

func Run() {
	// 采集GitHub
	go func() {
		for {
			handleErr(core.Github(), "Github")
			time.Sleep(time.Minute * 30)
		}
	}()
	// 其他页面
	go func() {
		for {
			handleErr(core.Tieba(), "Tieba")
			handleErr(core.ITHome(), "ITHome")
			handleErr(core.Weibo(), "Weibo")
			handleErr(core.Baidu(), "Baidu")
			handleErr(core.Bilibili(), "Bilibili")
			handleErr(core.Acfun(), "Acfun")
			handleErr(core.CSDN(), "CSDN")
			handleErr(core.Zhihu(), "Zhihu")
			handleErr(core.CCTV(), "CCTV")
			handleErr(core.ZWTX(), "ZWTX")
			handleErr(core.Juejin(), "Juejin")
			handleErr(core.Kr36(), "Kr36")
			//core.Douyin()
			handleErr(core.Douban(), "Douban")
			time.Sleep(time.Minute * 10)
		}
	}()
	// 开启服务
	//go core.ExportMetric()
	RunServer()
}

func RunServer() {
	gin.SetMode(gin.ReleaseMode)
	r := gin.Default()
	r.Use(gin.Recovery())

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

	r.GET("/douyin", func(c *gin.Context) {
		tmp := make(map[string]interface{})
		tmp["Flag"] = 0
		tmp["Data"] = model.M["DY"]
		c.HTML(200, "onelist.html", tmp)
	})

	r.GET("/weibo", func(c *gin.Context) {
		tmp := make(map[string]interface{})
		tmp["Flag"] = 0
		tmp["Data"] = model.M["Weibo"]
		c.HTML(200, "onelist.html", tmp)
	})

	r.GET("/baidu", func(c *gin.Context) {
		tmp := make(map[string]interface{})
		tmp["Flag"] = 0
		tmp["Data"] = model.M["Baidu"]
		c.HTML(200, "onelist.html", tmp)
	})

	r.GET("/tieba", func(c *gin.Context) {
		tmp := make(map[string]interface{})
		tmp["Flag"] = 0
		tmp["StyleFlag"] = 666
		tmp["Data"] = model.M["tieba"]
		c.HTML(200, "onelist.html", tmp)
	})

	r.GET("/cctv", func(c *gin.Context) {
		tmp := make(map[string]interface{})
		tmp["Flag"] = 0
		tmp["Data"] = model.M["CCTV"]
		c.HTML(200, "onelist.html", tmp)
	})

	r.GET("/csdn", func(c *gin.Context) {
		tmp := make(map[string]interface{})
		tmp["Flag"] = 0
		tmp["Data"] = model.M["CSDN"]
		c.HTML(200, "onelist.html", tmp)
	})

	r.GET("/zhihu", func(c *gin.Context) {
		tmp := make(map[string]interface{})
		tmp["Flag"] = 0
		tmp["Data"] = model.M["Zhihu"]
		c.HTML(200, "onelist.html", tmp)
	})

	r.GET("/acfun", func(c *gin.Context) {
		tmp := make(map[string]interface{})
		tmp["Flag"] = 0
		tmp["Data"] = model.M["Acfun"]
		c.HTML(200, "onelist.html", tmp)
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
		c.HTML(200, "threelist.html", R)
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
		c.HTML(200, "threelist.html", R)
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
		c.HTML(200, "threelist.html", R)
	})

	r.GET("/githubC", func(c *gin.Context) {
		var R S1
		R.Flag = 1
		R.Style = 20
		R.Data = append(R.Data,
			S2{Title: "Daily(C)", Content: model.M["Github9"]},
			S2{Title: "Weekly(C)", Content: model.M["Github10"]},
			S2{Title: "Monthly(C)", Content: model.M["Github11"]},
		)
		c.HTML(200, "threelist.html", R)
	})

	r.GET("/githubJava", func(c *gin.Context) {
		var R S1
		R.Flag = 1
		R.Style = 20
		R.Data = append(R.Data,
			S2{Title: "Daily(Java)", Content: model.M["Github12"]},
			S2{Title: "Weekly(Java)", Content: model.M["Github13"]},
			S2{Title: "Monthly(Java)", Content: model.M["Github14"]},
		)
		c.HTML(200, "threelist.html", R)
	})

	r.GET("/githubRust", func(c *gin.Context) {
		var R S1
		R.Flag = 1
		R.Style = 20
		R.Data = append(R.Data,
			S2{Title: "Daily(Rust)", Content: model.M["Github15"]},
			S2{Title: "Weekly(Rust)", Content: model.M["Github16"]},
			S2{Title: "Monthly(Rust)", Content: model.M["Github17"]},
		)
		c.HTML(200, "threelist.html", R)
	})

	r.GET("/githubCPP", func(c *gin.Context) {
		var R S1
		R.Flag = 1
		R.Style = 20
		R.Data = append(R.Data,
			S2{Title: "Daily(CPP)", Content: model.M["Github18"]},
			S2{Title: "Weekly(CPP)", Content: model.M["Github19"]},
			S2{Title: "Monthly(CPP)", Content: model.M["Github20"]},
		)
		c.HTML(200, "threelist.html", R)
	})

	r.GET("/githubJS", func(c *gin.Context) {
		var R S1
		R.Flag = 1
		R.Style = 20
		R.Data = append(R.Data,
			S2{Title: "Daily(JS)", Content: model.M["Github21"]},
			S2{Title: "Weekly(JS)", Content: model.M["Github22"]},
			S2{Title: "Monthly(JS)", Content: model.M["Github23"]},
		)
		c.HTML(200, "threelist.html", R)
	})

	r.GET("/githubShell", func(c *gin.Context) {
		var R S1
		R.Flag = 1
		R.Style = 20
		R.Data = append(R.Data,
			S2{Title: "Daily(Shell)", Content: model.M["Github24"]},
			S2{Title: "Weekly(Shell)", Content: model.M["Github25"]},
			S2{Title: "Monthly(Shell)", Content: model.M["Github26"]},
		)
		c.HTML(200, "threelist.html", R)
	})

	r.GET("/githubAssembly", func(c *gin.Context) {
		var R S1
		R.Flag = 1
		R.Style = 20
		R.Data = append(R.Data,
			S2{Title: "Daily(Assembly)", Content: model.M["Github27"]},
			S2{Title: "Weekly(Assembly)", Content: model.M["Github28"]},
			S2{Title: "Monthly(Assembly)", Content: model.M["Github29"]},
		)
		c.HTML(200, "threelist.html", R)
	})

	r.GET("/githubCSharp", func(c *gin.Context) {
		var R S1
		R.Flag = 1
		R.Style = 20
		R.Data = append(R.Data,
			S2{Title: "Daily(CSharp)", Content: model.M["Github30"]},
			S2{Title: "Weekly(CSharp)", Content: model.M["Github31"]},
			S2{Title: "Monthly(CSharp)", Content: model.M["Github32"]},
		)
		c.HTML(200, "threelist.html", R)
	})

	r.GET("/githubPHP", func(c *gin.Context) {
		var R S1
		R.Flag = 1
		R.Style = 20
		R.Data = append(R.Data,
			S2{Title: "Daily(PHP)", Content: model.M["Github33"]},
			S2{Title: "Weekly(PHP)", Content: model.M["Github34"]},
			S2{Title: "Monthly(PHP)", Content: model.M["Github35"]},
		)
		c.HTML(200, "threelist.html", R)
	})

	r.GET("/douban", func(c *gin.Context) {
		var R S1
		R.Flag = 1
		R.Style = 20
		R.Data = append(R.Data,
			S2{Title: "正在热映", Content: model.M["Douban0"]},
			S2{Title: "热门电影", Content: model.M["Douban1"]},
			S2{Title: "热门电视剧", Content: model.M["Douban2"]},
		)
		c.HTML(200, "threelist.html", R)
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
		c.HTML(200, "threelist.html", R)
	})

	r.GET("/ithome", func(c *gin.Context) {
		var R S1
		R.Flag = 1
		R.Style = 10
		R.Data = append(R.Data,
			S2{Title: "日榜", Content: model.M["ithome0"]},
			S2{Title: "周榜", Content: model.M["ithome1"]},
			S2{Title: "月榜", Content: model.M["ithome2"]},
		)
		c.HTML(200, "threelist.html", R)
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
		c.HTML(200, "threelist.html", R)
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
		c.HTML(200, "threelist.html", R)
	})

	r.GET("/tool", func(c *gin.Context) {
		var v model.ToolStruct
		b, _ := ioutil.ReadFile("./data/tool.txt")
		err := json.Unmarshal(b, &v)
		if err != nil {
			log.Printf("tool err:%v\n", err)
			return
		}
		c.HTML(200, "tool.html", v)
	})

	r.Run("127.0.0.1:8080")
}
