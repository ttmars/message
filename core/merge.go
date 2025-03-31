package core

import (
	"bytes"
	"encoding/json"
	"fmt"
	"github.com/antchfx/htmlquery"
	"io"
	"io/ioutil"
	"log"
	"log/slog"
	"message/model"
	"net/http"
	"strconv"
	"strings"
	"time"
)

func GetAll() {
	defer func() {
		if err := recover(); err != nil {
			slog.Error("panic occurred", "error", err)
		}
	}()

	Weibo()
	Baidu()
	Bilibili()
	Acfun()
	CSDN()
	Zhihu()
	CCTV()
	ZWTX()
}

func CSDN() (result []model.Item, err error) {
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
	for _, i := range v.Data {
		title := i.ArticleTitle
		link := i.ArticleDetailURL
		//badge := i.ViewCount + "浏览"
		result = append(result, model.Item{Name: title, Link: link})
	}
	if len(result) >= model.Num {
		model.M["CSDN"] = result[:model.Num]
	} else {
		model.M["CSDN"] = result
	}
	model.M["CSDN"] = append(model.M["CSDN"], model.Item{Name: "更多", Link: "https://blog.csdn.net/rank/list"})
	return
}

func Weibo() (result []model.Item, err error) {
	url := "https://s.weibo.com/top/summary"
	req, _ := http.NewRequest("GET", url, nil)
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

	doc, err := htmlquery.Parse(bytes.NewReader(bodyText))
	if err != nil {
		log.Printf("Weibo err:%v\n", err)
		return
	}
	nodes, err := htmlquery.QueryAll(doc, "//*[@id=\"pl_top_realtimehot\"]/table/tbody/tr/td[2]/a")
	if err != nil {
		log.Printf("Weibo err:%v\n", err)
		return
	}
	for _, node := range nodes {
		title := htmlquery.InnerText(node)
		href := htmlquery.SelectAttr(node, "href")
		link := "https://s.weibo.com" + href
		if strings.HasPrefix(href, "/weibo") {
			result = append(result, model.Item{Name: title, Link: link})
		}
	}
	if len(result) >= model.Num {
		model.M["Weibo"] = result[:model.Num]
	} else {
		model.M["Weibo"] = result
	}
	model.M["Weibo"] = append(model.M["Weibo"], model.Item{Name: "更多", Link: url})
	return
}

func Baidu() (result []model.Item, err error) {
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

	doc, err := htmlquery.Parse(bytes.NewReader(bodyText))
	if err != nil {
		log.Printf("Baidu err:%v\n", err)
		return
	}
	nodes, err := htmlquery.QueryAll(doc, "//*[@id=\"sanRoot\"]/main/div[2]/div/div[2]/div/div[2]/a/div[1]")
	if err != nil {
		log.Printf("Baidu err:%v\n", err)
		return
	}
	for _, node := range nodes {
		title := strings.TrimSpace(htmlquery.InnerText(node))
		link := htmlquery.SelectAttr(node.Parent, "href")
		result = append(result, model.Item{Name: title, Link: link})
	}
	if len(result) >= model.Num {
		model.M["Baidu"] = result[:model.Num]
	} else {
		model.M["Baidu"] = result
	}
	model.M["Baidu"] = append(model.M["Baidu"], model.Item{Name: "更多", Link: url})
	return
}

func Bilibili() (result []model.Item, err error) {
	url := "https://api.bilibili.com/x/web-interface/ranking/v2?rid=0&type=all&web_location=333.934&w_rid=53b20296ec99dd4db6409ea8bcf0ad45&wts=1743417372"
	method := "GET"

	client := &http.Client{}
	req, err := http.NewRequest(method, url, nil)

	if err != nil {
		fmt.Println(err)
		return
	}
	req.Header.Add("accept", "*/*")
	req.Header.Add("accept-language", "zh-CN,zh;q=0.9")
	req.Header.Add("cache-control", "no-cache")
	req.Header.Add("origin", "https://www.bilibili.com")
	req.Header.Add("pragma", "no-cache")
	req.Header.Add("priority", "u=1, i")
	req.Header.Add("referer", "https://www.bilibili.com/v/popular/rank/all")
	req.Header.Add("sec-ch-ua", "\"Chromium\";v=\"134\", \"Not:A-Brand\";v=\"24\", \"Google Chrome\";v=\"134\"")
	req.Header.Add("sec-ch-ua-mobile", "?0")
	req.Header.Add("sec-ch-ua-platform", "\"Windows\"")
	req.Header.Add("sec-fetch-dest", "empty")
	req.Header.Add("sec-fetch-mode", "cors")
	req.Header.Add("sec-fetch-site", "same-site")
	req.Header.Add("user-agent", "Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/134.0.0.0 Safari/537.36")

	res, err := client.Do(req)
	if err != nil {
		fmt.Println(err)
		return
	}
	defer res.Body.Close()

	body, err := ioutil.ReadAll(res.Body)
	if err != nil {
		fmt.Println(err)
		return
	}
	//fmt.Println(string(body))

	var v model.BiliStruct
	json.Unmarshal(body, &v)

	for _, node := range v.Data.List {
		title := strings.TrimSpace(node.Title)
		link := node.ShortLinkV2
		result = append(result, model.Item{Name: title, Link: link})
		//fmt.Println(title, link)
	}
	if len(result) >= model.Num {
		model.M["Bilibili"] = result[:model.Num]
	} else {
		model.M["Bilibili"] = result
	}
	model.M["Bilibili"] = append(model.M["Bilibili"], model.Item{Name: "更多", Link: "https://www.bilibili.com/v/popular/rank/all"})
	return
}

func Acfun() (result []model.Item, err error) {
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
	for _, i := range v.RankList {
		title := i.ContentTitle
		link := "https://www.acfun.cn/v/ac" + i.DougaID
		result = append(result, model.Item{Name: title, Link: link})
	}
	if len(result) >= model.Num {
		model.M["Acfun"] = result[:model.Num]
	} else {
		model.M["Acfun"] = result
	}
	model.M["Acfun"] = append(model.M["Acfun"], model.Item{Name: "更多", Link: "https://www.acfun.cn/rank/list"})
	return
}

func Zhihu() {
	defer func() {
		if err := recover(); err != nil {
			log.Printf("Zhihu panic:%v\n", err)
		}
	}()
	url := "https://www.zhihu.com/billboard"
	r, _ := http.NewRequest("GET", url, nil)
	r.Header.Set("User-Agent", "Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/99.0.4844.84 Safari/537.36")
	resp, err := http.DefaultClient.Do(r)
	if err != nil {
		log.Printf("Zhihu err:%v\n", err)
		return
	}
	defer resp.Body.Close()
	b, err := io.ReadAll(resp.Body)
	if err != nil {
		log.Printf("Zhihu err:%v\n", err)
		return
	}

	doc, err := htmlquery.Parse(bytes.NewReader(b))
	if err != nil {
		log.Printf("Zhihu err:%v\n", err)
		return
	}
	node, err := htmlquery.Query(doc, "//*[@id=\"js-initialData\"]")
	if err != nil {
		log.Printf("Zhihu err:%v\n", err)
		return
	}
	bodyText := []byte(htmlquery.InnerText(node))

	v := &model.ZhihuStruct{}
	err = json.Unmarshal(bodyText, v)
	var result []model.Item
	for _, i := range v.InitialState.Topstory.HotList {
		title := i.Target.TitleArea.Text
		link := "https://www.zhihu.com/question/" + i.CardID[2:]
		result = append(result, model.Item{Name: title, Link: link})
	}

	if len(result) >= model.Num {
		model.M["Zhihu"] = result[:model.Num]
	} else {
		model.M["Zhihu"] = result
	}
	model.M["Zhihu"] = append(model.M["Zhihu"], model.Item{Name: "更多", Link: url})
	return
}

func CCTV() {
	url := "https://tv.cctv.com/lm/xwlb/"
	t := time.Now().AddDate(0, 0, -1).Format("20060102")
	dataUrl := "https://tv.cctv.com/lm/xwlb/day/" + t + ".shtml"
	req, _ := http.NewRequest("GET", dataUrl, nil)
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

	doc, err := htmlquery.Parse(bytes.NewReader(bodyText))
	if err != nil {
		log.Printf("CCTV err:%v\n", err)
		return
	}
	nodes, err := htmlquery.QueryAll(doc, "/html/body/li/div/a")
	if err != nil {
		log.Printf("CCTV err:%v\n", err)
		return
	}
	var result []model.Item
	for _, node := range nodes {
		title := htmlquery.SelectAttr(node, "alt")
		link := htmlquery.SelectAttr(node, "href")
		result = append(result, model.Item{Name: title, Link: link})
	}

	if len(result) >= model.Num {
		model.M["CCTV"] = result[:model.Num]
	} else {
		model.M["CCTV"] = result
	}
	model.M["CCTV"] = append(model.M["CCTV"], model.Item{Name: "更多", Link: url})
	return
}

func ZWTX() {
	url := "https://tv.cctv.com/lm/zwtx/index.shtml"
	t := time.Now().Format("20060102")
	dataUrl := "https://api.cntv.cn/NewVideo/getVideoListByColumn?id=TOPC1451558496100826&n=100&sort=desc&p=1&bd=" + t + "&mode=2&serviceId=tvcctv&cb=cb"
	req, _ := http.NewRequest("GET", dataUrl, nil)
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
	for _, data := range v.Data.List {
		title := data.Brief
		link := data.URL
		if data.Mode == 0 {
			title = data.Title
		}
		if strings.Contains(data.Time, "06:") {
			T["zwtx0"] = append(T["zwtx0"], model.Item{Name: title, Link: link})
		} else if strings.Contains(data.Time, "07:") {
			T["zwtx1"] = append(T["zwtx1"], model.Item{Name: title, Link: link})
		} else {
			T["zwtx2"] = append(T["zwtx2"], model.Item{Name: title, Link: link})
		}
	}
	model.M["zwtx0"] = T["zwtx0"]
	model.M["zwtx1"] = T["zwtx1"]
	model.M["zwtx2"] = T["zwtx2"]

	for i := 0; i <= 2; i++ {
		k := "zwtx" + strconv.Itoa(i)
		if len(model.M[k]) >= model.Num {
			model.M[k] = model.M[k][:model.Num]
		}
		model.M[k] = append(model.M[k], model.Item{Name: "更多", Link: url})
	}
	return
}
