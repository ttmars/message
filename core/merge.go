package core

import (
	"bytes"
	"encoding/json"
	"github.com/antchfx/htmlquery"
	"io"
	"io/ioutil"
	"log"
	"message/model"
	"net/http"
	"strconv"
	"strings"
	"sync"
	"time"
)

var wg = &sync.WaitGroup{}

func GetAll()  {
	wg.Add(8)
	go Weibo()
	go Baidu()
	go Bilibili()
	go Acfun()
	go CSDN()
	go Zhihu()
	go CCTV()
	go ZWTX()
	wg.Wait()
}

func CSDN()(result []model.Item, err error)  {
	defer wg.Done()
	url := "https://blog.csdn.net/phoenix/web/blog/hot-rank?page=0&pageSize=25&type="
	resp, err := model.DFClient.Get(url)
	if err != nil {
		log.Printf("CSDN err:%v\n", err)
		return
	}
	defer resp.Body.Close()
	bodyText, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Printf("CSDN err:%v\n", err)
		return
	}

	v := &model.CSDNStruct{}
	err = json.Unmarshal(bodyText, v)
	for _,i := range v.Data {
		title := i.ArticleTitle
		link := i.ArticleDetailURL
		//badge := i.ViewCount + "浏览"
		result = append(result, model.Item{Name: title, Link: link})
	}
	if len(result) >= model.Num {
		model.M["CSDN"] = result[:model.Num]
	}else{
		model.M["CSDN"] = result
	}
	model.M["CSDN"] = append(model.M["CSDN"], model.Item{Name: "更多", Link: "https://blog.csdn.net/rank/list"})
	return
}

func Weibo() (result []model.Item, err error) {
	defer wg.Done()
	url := "https://s.weibo.com/top/summary"
	req,_ := http.NewRequest("GET", url, nil)
	req.Header.Set("cookie", "SUB=_2AkMVZUSbf8NxqwFRmP0WxG_jb41yyAvEieKjObVAJRMxHRl-yT9jqlULtRB6PuVqdHCUYilcOycjbJG7lmJA3vZSjI41; SUBP=0033WrSXqPxfM72-Ws9jqgMF55529P9D9WFlfySZ479DB-M3cI_uFeOE")
	resp, err := model.DFClient.Do(req)
	if err != nil {
		log.Printf("Weibo err:%v\n", err)
		return
	}
	defer resp.Body.Close()
	bodyText, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Printf("Weibo err:%v\n", err)
		return
	}

	doc,err := htmlquery.Parse(bytes.NewReader(bodyText))
	if err != nil {
		log.Printf("Weibo err:%v\n", err)
		return
	}
	nodes,err := htmlquery.QueryAll(doc, "//*[@id=\"pl_top_realtimehot\"]/table/tbody/tr/td[2]/a")
	if err != nil {
		log.Printf("Weibo err:%v\n", err)
		return
	}
	for _,node := range nodes{
		title := htmlquery.InnerText(node)
		href := htmlquery.SelectAttr(node, "href")
		link := "https://s.weibo.com" + href
		if strings.HasPrefix(href, "/weibo"){
			result = append(result, model.Item{Name:title, Link: link})
		}
	}
	if len(result) >= model.Num {
		model.M["Weibo"] = result[:model.Num]
	}else{
		model.M["Weibo"] = result
	}
	model.M["Weibo"] = append(model.M["Weibo"], model.Item{Name: "更多", Link: url})
	return
}

func Baidu()(result []model.Item, err error){
	defer wg.Done()
	url := "https://top.baidu.com/board?tab=realtime"
	resp, err := model.DFClient.Get(url)
	if err != nil {
		log.Printf("Baidu err:%v\n", err)
		return
	}
	defer resp.Body.Close()
	bodyText, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Printf("Baidu err:%v\n", err)
		return
	}

	doc,err := htmlquery.Parse(bytes.NewReader(bodyText))
	if err != nil {
		log.Printf("Baidu err:%v\n", err)
		return
	}
	nodes,err := htmlquery.QueryAll(doc, "//*[@id=\"sanRoot\"]/main/div[2]/div/div[2]/div/div[2]/a/div[1]")
	if err != nil {
		log.Printf("Baidu err:%v\n", err)
		return
	}
	for _,node := range nodes{
		title := strings.TrimSpace(htmlquery.InnerText(node))
		link := htmlquery.SelectAttr(node.Parent, "href")
		result = append(result, model.Item{Name: title, Link: link})
	}
	if len(result) >= model.Num {
		model.M["Baidu"] = result[:model.Num]
	}else{
		model.M["Baidu"] = result
	}
	model.M["Baidu"] = append(model.M["Baidu"], model.Item{Name: "更多", Link: url})
	return
}


func Bilibili()(result []model.Item, err error){
	defer wg.Done()
	url := "https://www.bilibili.com/v/popular/rank/all"
	req, _ := http.NewRequest("GET", url, nil)
	req.Header.Set("cookie", "buvid3=5D5418C6-B76A-475E-9EF5-618C6121ED1A167614infoc; b_lsid=DFBF6551_17FB45E250C; bsource=search_baidu; _uuid=967E76C2-AEA1-8B110-FB99-8425ABA3434C43121infoc; LIVE_BUVID=AUTO1016479985451625; buvid_fp=d2c03acdbde02a5f747fa45978e2213e; buvid4=377A79AA-5113-B55B-B36C-5925BFA3E9E745806-022032309-DsswV1JjVn6AGA18EUjH+g%3D%3D; innersign=0; b_ut=7; i-wanna-go-back=2; PVID=2")
	resp, err := model.DFClient.Do(req)
	if err != nil {
		log.Printf("Bilibili err:%v\n", err)
		return
	}
	defer resp.Body.Close()
	bodyText, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Printf("Bilibili err:%v\n", err)
		return
	}

	doc,err := htmlquery.Parse(bytes.NewReader(bodyText))
	if err != nil {
		log.Printf("Bilibili err:%v\n", err)
		return
	}
	nodes,err := htmlquery.QueryAll(doc, "//*[@id=\"app\"]/div/div[2]/div[2]/ul/li/div/div[2]/a")
	if err != nil {
		log.Printf("Bilibili err:%v\n", err)
		return
	}
	for _,node := range nodes{
		title := strings.TrimSpace(htmlquery.InnerText(node))
		link := "https:" + htmlquery.SelectAttr(node, "href")
		result = append(result, model.Item{Name: title, Link: link})
	}
	if len(result) >= model.Num {
		model.M["Bilibili"] = result[:model.Num]
	}else{
		model.M["Bilibili"] = result
	}
	model.M["Bilibili"] = append(model.M["Bilibili"], model.Item{Name: "更多", Link: url})
	return
}

func Acfun()(result []model.Item, err error){
	defer wg.Done()
	url := "https://www.acfun.cn/rest/pc-direct/rank/channel?channelId=&subChannelId=&rankLimit=30&rankPeriod=DAY"
	req, _ := http.NewRequest("GET", url, nil)
	req.Header.Set("user-agent", "Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/99.0.4844.82 Safari/537.36")
	resp, err := model.DFClient.Do(req)
	if err != nil {
		log.Printf("Acfun err:%v\n", err)
		return
	}
	defer resp.Body.Close()
	bodyText, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Printf("Acfun err:%v\n", err)
		return
	}

	v := &model.AcfunStruct{}
	err = json.Unmarshal(bodyText, v)
	for _,i := range v.RankList {
		title := i.ContentTitle
		link := "https://www.acfun.cn/v/ac" + i.DougaID
		result = append(result, model.Item{Name: title, Link: link})
	}
	if len(result) >= model.Num {
		model.M["Acfun"] = result[:model.Num]
	}else{
		model.M["Acfun"] = result
	}
	model.M["Acfun"] = append(model.M["Acfun"], model.Item{Name: "更多", Link: "https://www.acfun.cn/rank/list"})
	return
}

func Zhihu()  {
	defer wg.Done()
	url := "https://www.zhihu.com/billboard"
	r,_ := http.NewRequest("GET", url, nil)
	r.Header.Set("User-Agent", "Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/99.0.4844.84 Safari/537.36")
	resp,err := http.DefaultClient.Do(r)
	if err != nil {
		log.Printf("Zhihu err:%v\n", err)
		return
	}
	defer resp.Body.Close()
	b,err := io.ReadAll(resp.Body)
	if err != nil {
		log.Printf("Zhihu err:%v\n", err)
		return
	}

	doc,err := htmlquery.Parse(bytes.NewReader(b))
	if err != nil {
		log.Printf("Zhihu err:%v\n", err)
		return
	}
	node,err := htmlquery.Query(doc, "//*[@id=\"js-initialData\"]")
	if err != nil {
		log.Printf("Zhihu err:%v\n", err)
		return
	}
	bodyText := []byte(htmlquery.InnerText(node))

	v := &model.ZhihuStruct{}
	err = json.Unmarshal(bodyText, v)
	var result []model.Item
	for _,i := range v.InitialState.Topstory.HotList{
		title := i.Target.TitleArea.Text
		link := "https://www.zhihu.com/question/" + i.CardID[2:]
		result = append(result, model.Item{Name: title, Link: link})
	}

	if len(result) >= model.Num {
		model.M["Zhihu"] = result[:model.Num]
	}else{
		model.M["Zhihu"] = result
	}
	model.M["Zhihu"] = append(model.M["Zhihu"], model.Item{Name: "更多", Link: url})
	return
}

func CCTV() {
	defer wg.Done()
	url := "https://tv.cctv.com/lm/xwlb/"
	t := time.Now().AddDate(0, 0, -1).Format("20060102")
	dataUrl := "https://tv.cctv.com/lm/xwlb/day/" + t +".shtml"
	req,_ := http.NewRequest("GET", dataUrl, nil)
	resp, err := model.DFClient.Do(req)
	if err != nil {
		log.Printf("CCTV err:%v\n", err)
		return
	}
	defer resp.Body.Close()
	bodyText, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Printf("CCTV err:%v\n", err)
		return
	}

	doc,err := htmlquery.Parse(bytes.NewReader(bodyText))
	if err != nil {
		log.Printf("CCTV err:%v\n", err)
		return
	}
	nodes,err := htmlquery.QueryAll(doc, "/html/body/li/div/a")
	if err != nil {
		log.Printf("CCTV err:%v\n", err)
		return
	}
	var result []model.Item
	for _,node := range nodes{
		title := htmlquery.SelectAttr(node, "alt")
		link := htmlquery.SelectAttr(node, "href")
		result = append(result, model.Item{Name:title, Link: link})
	}

	if len(result) >= model.Num {
		model.M["CCTV"] = result[:model.Num]
	}else{
		model.M["CCTV"] = result
	}
	model.M["CCTV"] = append(model.M["CCTV"], model.Item{Name: "更多", Link: url})
	return
}

func ZWTX() {
	defer wg.Done()
	url := "https://tv.cctv.com/lm/zwtx/index.shtml"
	t := time.Now().Format("20060102")
	dataUrl := "https://api.cntv.cn/NewVideo/getVideoListByColumn?id=TOPC1451558496100826&n=100&sort=desc&p=1&bd=" + t + "&mode=2&serviceId=tvcctv&cb=cb"
	req,_ := http.NewRequest("GET", dataUrl, nil)
	resp, err := model.DFClient.Do(req)
	if err != nil {
		log.Printf("ZWTX err:%v\n", err)
		return
	}
	defer resp.Body.Close()
	bodyText, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Printf("ZWTX err:%v\n", err)
		return
	}
	type AutoGenerated struct {
		Data struct {
			Total int `json:"total"`
			List  []struct {
				GUID      string `json:"guid"`
				ID        string `json:"id"`
				Time      string `json:"time"`
				Title     string `json:"title"`
				Length    string `json:"length"`
				Image     string `json:"image"`
				FocusDate int64  `json:"focus_date"`
				Brief     string `json:"brief"`
				URL       string `json:"url"`
				Mode      int    `json:"mode"`
			} `json:"list"`
		} `json:"data"`
	}
	var v AutoGenerated
	err = json.Unmarshal(bodyText[3:len(bodyText)-2], &v)
	if err != nil {
		log.Printf("ZWTX err:%v\n", err)
		return
	}
	var T = make(map[string][]model.Item)
	for _,data := range v.Data.List {
		title := data.Brief
		link := data.URL
		if data.Mode == 0 {
			title = data.Title
		}
		if strings.Contains(data.Time, "06:") {
			T["zwtx0"] = append(model.M["zwtx0"], model.Item{Name:title, Link: link})
		}else if strings.Contains(data.Time, "07:") {
			T["zwtx1"] = append(model.M["zwtx1"], model.Item{Name:title, Link: link})
		}else {
			T["zwtx2"] = append(model.M["zwtx2"], model.Item{Name:title, Link: link})
		}
	}
	model.M["zwtx0"] = T["zwtx0"]
	model.M["zwtx1"] = T["zwtx1"]
	model.M["zwtx2"] = T["zwtx2"]

	for i:=0;i<=2;i++ {
		k := "zwtx" + strconv.Itoa(i)
		if len(model.M[k]) >= model.Num {
			model.M[k] = model.M[k][:model.Num]
		}
		model.M[k] = append(model.M[k], model.Item{Name: "更多", Link: url})
	}
	return
}