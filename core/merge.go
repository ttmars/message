package core

import (
	"bytes"
	"encoding/json"
	"github.com/antchfx/htmlquery"
	"io"
	"io/ioutil"
	"message/model"
	"net/http"
	"strconv"
	"strings"
	"time"
)

func CSDN() error {
	var result []model.Item
	url := "https://blog.csdn.net/phoenix/web/blog/hot-rank?page=0&pageSize=25&type="
	resp, err := model.DFClient.Get(url)
	if err != nil {
		return err
	}
	defer resp.Body.Close()
	bodyText, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return err
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
	return nil
}

func Weibo() error {
	url := "https://s.weibo.com/top/summary"
	req, _ := http.NewRequest("GET", url, nil)
	req.Header.Set("cookie", "SUB=_2AkMVZUSbf8NxqwFRmP0WxG_jb41yyAvEieKjObVAJRMxHRl-yT9jqlULtRB6PuVqdHCUYilcOycjbJG7lmJA3vZSjI41; SUBP=0033WrSXqPxfM72-Ws9jqgMF55529P9D9WFlfySZ479DB-M3cI_uFeOE")
	resp, err := model.DFClient.Do(req)
	if err != nil {
		return err
	}
	defer resp.Body.Close()
	bodyText, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return err
	}

	doc, err := htmlquery.Parse(bytes.NewReader(bodyText))
	if err != nil {
		return err
	}
	nodes, err := htmlquery.QueryAll(doc, "//*[@id=\"pl_top_realtimehot\"]/table/tbody/tr/td[2]/a")
	if err != nil {
		return err
	}
	var result []model.Item
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
	return nil
}

func Baidu() error {
	var result []model.Item
	url := "https://top.baidu.com/board?tab=realtime"
	resp, err := model.DFClient.Get(url)
	if err != nil {
		return err
	}
	defer resp.Body.Close()
	bodyText, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return err
	}

	doc, err := htmlquery.Parse(bytes.NewReader(bodyText))
	if err != nil {
		return err
	}
	nodes, err := htmlquery.QueryAll(doc, "//*[@id=\"sanRoot\"]/main/div[2]/div/div[2]/div/div[2]/a/div[1]")
	if err != nil {
		return err
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
	return nil
}

func Bilibili() error {
	var result []model.Item
	url := "https://api.bilibili.com/x/web-interface/ranking/v2?rid=0&type=all&web_location=333.934&w_rid=53b20296ec99dd4db6409ea8bcf0ad45&wts=1743417372"
	method := "GET"

	client := &http.Client{}
	req, err := http.NewRequest(method, url, nil)

	if err != nil {
		return err
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
		return err
	}
	defer res.Body.Close()

	body, err := ioutil.ReadAll(res.Body)
	if err != nil {
		return err
	}

	var v model.BiliStruct
	err = json.Unmarshal(body, &v)
	if err != nil {
		return err
	}

	for _, node := range v.Data.List {
		title := strings.TrimSpace(node.Title)
		link := node.ShortLinkV2
		result = append(result, model.Item{Name: title, Link: link})
	}
	if len(result) >= model.Num {
		model.M["Bilibili"] = result[:model.Num]
	} else {
		model.M["Bilibili"] = result
	}
	model.M["Bilibili"] = append(model.M["Bilibili"], model.Item{Name: "更多", Link: "https://www.bilibili.com/v/popular/rank/all"})
	return nil
}

func Acfun() error {
	var result []model.Item
	url := "https://www.acfun.cn/rest/pc-direct/rank/channel?channelId=&subChannelId=&rankLimit=30&rankPeriod=DAY"
	req, _ := http.NewRequest("GET", url, nil)
	req.Header.Set("user-agent", "Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/99.0.4844.82 Safari/537.36")
	resp, err := model.DFClient.Do(req)
	if err != nil {
		return err
	}
	defer resp.Body.Close()
	bodyText, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return err
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
	return err
}

func Zhihu() error {
	url := "https://www.zhihu.com/billboard"
	r, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return err
	}
	r.Header.Set("User-Agent", "Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/99.0.4844.84 Safari/537.36")
	resp, err := http.DefaultClient.Do(r)
	if err != nil {
		return err
	}
	defer resp.Body.Close()
	b, err := io.ReadAll(resp.Body)
	if err != nil {
		return err
	}

	doc, err := htmlquery.Parse(bytes.NewReader(b))
	if err != nil {
		return err
	}
	node, err := htmlquery.Query(doc, "//*[@id=\"js-initialData\"]")
	if err != nil {
		return err
	}
	bodyText := []byte(htmlquery.InnerText(node))

	v := &model.ZhihuStruct{}
	err = json.Unmarshal(bodyText, v)
	if err != nil {
		return err
	}
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
	return nil
}

func CCTV() error {
	url := "https://tv.cctv.com/lm/xwlb/"
	t := time.Now().AddDate(0, 0, -1).Format("20060102")
	dataUrl := "https://tv.cctv.com/lm/xwlb/day/" + t + ".shtml"
	req, _ := http.NewRequest("GET", dataUrl, nil)
	resp, err := model.DFClient.Do(req)
	if err != nil {
		return err
	}
	defer resp.Body.Close()
	bodyText, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return err
	}

	doc, err := htmlquery.Parse(bytes.NewReader(bodyText))
	if err != nil {
		return err
	}
	nodes, err := htmlquery.QueryAll(doc, "/html/body/li/div/a")
	if err != nil {
		return err
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
	return nil
}

func ZWTX() error {
	url := "https://tv.cctv.com/lm/zwtx/index.shtml"
	t := time.Now().Format("20060102")
	dataUrl := "https://api.cntv.cn/NewVideo/getVideoListByColumn?id=TOPC1451558496100826&n=100&sort=desc&p=1&bd=" + t + "&mode=2&serviceId=tvcctv&cb=cb"
	req, _ := http.NewRequest("GET", dataUrl, nil)
	resp, err := model.DFClient.Do(req)
	if err != nil {
		return err
	}
	defer resp.Body.Close()
	bodyText, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return err
	}

	var v model.ZWTXStruct
	err = json.Unmarshal(bodyText[3:len(bodyText)-2], &v)
	if err != nil {
		return err
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
	return nil
}
