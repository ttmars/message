package core

import (
	"bytes"
	"encoding/json"
	"fmt"
	"github.com/antchfx/htmlquery"
	"io"
	"log"
	"message/model"
	"net/http"
)

func Douban() {
	Douban0()
	Douban1()
	Douban2()
}

// Douban2 热门电视剧
func Douban2() {
	client := &http.Client{}
	req, err := http.NewRequest("GET", "https://movie.douban.com/j/search_subjects?type=tv&tag=%E7%83%AD%E9%97%A8&page_limit=50&page_start=0", nil)
	if err != nil {
		log.Fatal(err)
	}
	req.Header.Set("Accept", "*/*")
	req.Header.Set("Accept-Language", "zh-CN,zh-TW;q=0.9,zh;q=0.8,en-US;q=0.7,en;q=0.6")
	req.Header.Set("Cache-Control", "no-cache")
	req.Header.Set("Connection", "keep-alive")
	req.Header.Set("Cookie", `bid=SMcAdHx7Y7c; __utma=30149280.129623236.1690791617.1690791617.1690791617.1; __utmc=30149280; __utmz=30149280.1690791617.1.1.utmcsr=(direct)|utmccn=(direct)|utmcmd=(none); ll="108306"; _pk_id.100001.4cf6=8940335d61e37968.1690792487.; _pk_ses.100001.4cf6=1; __utma=223695111.1615737067.1690792487.1690792487.1690792487.1; __utmb=223695111.0.10.1690792487; __utmc=223695111; __utmz=223695111.1690792487.1.1.utmcsr=(direct)|utmccn=(direct)|utmcmd=(none); __yadk_uid=IR5CxAexo5QHtSmacvWFNLOlUbcUMkdA; _vwo_uuid_v2=DCF4AF867C112C5776549AE2551E67AB8|88907c3bfc1f34f5f2cdaa956ab78706; __utmb=30149280.7.10.1690791617; __gads=ID=e99f5b64396fb6fa-2202e93cb8e70017:T=1690792557:RT=1690797908:S=ALNI_MbhbusJuJJwMW7HKDaZjK11zcczBQ; __gpi=UID=00000d348852e513:T=1690792557:RT=1690797908:S=ALNI_MYdQIC4MT58yzX766DkwwNYEBWXVQ; ap_v=0,6.0`)
	req.Header.Set("Pragma", "no-cache")
	req.Header.Set("Referer", "https://movie.douban.com/")
	req.Header.Set("Sec-Fetch-Dest", "empty")
	req.Header.Set("Sec-Fetch-Mode", "cors")
	req.Header.Set("Sec-Fetch-Site", "same-origin")
	req.Header.Set("User-Agent", "Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/115.0.0.0 Safari/537.36")
	req.Header.Set("X-Requested-With", "XMLHttpRequest")
	req.Header.Set("sec-ch-ua", `"Not/A)Brand";v="99", "Google Chrome";v="115", "Chromium";v="115"`)
	req.Header.Set("sec-ch-ua-mobile", "?0")
	req.Header.Set("sec-ch-ua-platform", `"Windows"`)
	resp, err := client.Do(req)
	if err != nil {
		log.Fatal(err)
	}
	defer resp.Body.Close()
	bodyText, err := io.ReadAll(resp.Body)
	if err != nil {
		log.Fatal(err)
	}

	var v model.Douban1Struct
	json.Unmarshal(bodyText, &v)

	var result []model.Item
	for _, item := range v.Subjects {
		if item.Rate == "" {
			item.Rate = "暂无评分"
		}
		result = append(result, model.Item{Name: item.Title, Link: item.URL, Description: "", Badge: item.Rate})
	}

	if len(result) >= model.GithubNum {
		model.M["Douban2"] = result[:model.Num]
	} else {
		model.M["Douban2"] = result
	}
	model.M["Douban2"] = append(model.M["Douban2"], model.Item{Name: "更多", Link: "https://movie.douban.com/"})

	log.Println("Douban2 success!!")
	return
}

// Douban1 热门电影
func Douban1() {
	client := &http.Client{}
	req, err := http.NewRequest("GET", "https://movie.douban.com/j/search_subjects?type=movie&tag=%E7%83%AD%E9%97%A8&page_limit=50&page_start=0", nil)
	if err != nil {
		log.Fatal(err)
	}
	req.Header.Set("Accept", "*/*")
	req.Header.Set("Accept-Language", "zh-CN,zh-TW;q=0.9,zh;q=0.8,en-US;q=0.7,en;q=0.6")
	req.Header.Set("Cache-Control", "no-cache")
	req.Header.Set("Connection", "keep-alive")
	req.Header.Set("Cookie", `bid=SMcAdHx7Y7c; __utma=30149280.129623236.1690791617.1690791617.1690791617.1; __utmc=30149280; __utmz=30149280.1690791617.1.1.utmcsr=(direct)|utmccn=(direct)|utmcmd=(none); ll="108306"; _pk_id.100001.4cf6=8940335d61e37968.1690792487.; _pk_ses.100001.4cf6=1; __utma=223695111.1615737067.1690792487.1690792487.1690792487.1; __utmb=223695111.0.10.1690792487; __utmc=223695111; __utmz=223695111.1690792487.1.1.utmcsr=(direct)|utmccn=(direct)|utmcmd=(none); __yadk_uid=IR5CxAexo5QHtSmacvWFNLOlUbcUMkdA; _vwo_uuid_v2=DCF4AF867C112C5776549AE2551E67AB8|88907c3bfc1f34f5f2cdaa956ab78706; __utmb=30149280.7.10.1690791617; __gads=ID=e99f5b64396fb6fa-2202e93cb8e70017:T=1690792557:RT=1690797908:S=ALNI_MbhbusJuJJwMW7HKDaZjK11zcczBQ; __gpi=UID=00000d348852e513:T=1690792557:RT=1690797908:S=ALNI_MYdQIC4MT58yzX766DkwwNYEBWXVQ; ap_v=0,6.0`)
	req.Header.Set("Pragma", "no-cache")
	req.Header.Set("Referer", "https://movie.douban.com/")
	req.Header.Set("Sec-Fetch-Dest", "empty")
	req.Header.Set("Sec-Fetch-Mode", "cors")
	req.Header.Set("Sec-Fetch-Site", "same-origin")
	req.Header.Set("User-Agent", "Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/115.0.0.0 Safari/537.36")
	req.Header.Set("X-Requested-With", "XMLHttpRequest")
	req.Header.Set("sec-ch-ua", `"Not/A)Brand";v="99", "Google Chrome";v="115", "Chromium";v="115"`)
	req.Header.Set("sec-ch-ua-mobile", "?0")
	req.Header.Set("sec-ch-ua-platform", `"Windows"`)
	resp, err := client.Do(req)
	if err != nil {
		log.Fatal(err)
	}
	defer resp.Body.Close()
	bodyText, err := io.ReadAll(resp.Body)
	if err != nil {
		log.Fatal(err)
	}

	var v model.Douban1Struct
	json.Unmarshal(bodyText, &v)

	var result []model.Item
	for _, item := range v.Subjects {
		if item.Rate == "" {
			item.Rate = "暂无评分"
		}
		result = append(result, model.Item{Name: item.Title, Link: item.URL, Description: "", Badge: item.Rate})
	}

	if len(result) >= model.GithubNum {
		model.M["Douban1"] = result[:model.Num]
	} else {
		model.M["Douban1"] = result
	}
	model.M["Douban1"] = append(model.M["Douban1"], model.Item{Name: "更多", Link: "https://movie.douban.com/"})

	log.Println("Douban1 success!!")
	return
}

// Douban0 正在热映
func Douban0() {
	req, err := http.NewRequest("GET", "https://movie.douban.com/", nil)
	if err != nil {
		log.Fatal(err)
	}
	req.Header.Set("Accept", "text/html,application/xhtml+xml,application/xml;q=0.9,image/avif,image/webp,image/apng,*/*;q=0.8,application/signed-exchange;v=b3;q=0.7")
	req.Header.Set("Accept-Language", "zh-CN,zh-TW;q=0.9,zh;q=0.8,en-US;q=0.7,en;q=0.6")
	req.Header.Set("Cache-Control", "no-cache")
	req.Header.Set("Connection", "keep-alive")
	req.Header.Set("Cookie", `bid=SMcAdHx7Y7c; ap_v=0,6.0; __utma=30149280.129623236.1690791617.1690791617.1690791617.1; __utmc=30149280; __utmz=30149280.1690791617.1.1.utmcsr=(direct)|utmccn=(direct)|utmcmd=(none); ll="108306"; _pk_id.100001.4cf6=8940335d61e37968.1690792487.; _pk_ses.100001.4cf6=1; __utma=223695111.1615737067.1690792487.1690792487.1690792487.1; __utmb=223695111.0.10.1690792487; __utmc=223695111; __utmz=223695111.1690792487.1.1.utmcsr=(direct)|utmccn=(direct)|utmcmd=(none); __yadk_uid=IR5CxAexo5QHtSmacvWFNLOlUbcUMkdA; _vwo_uuid_v2=DCF4AF867C112C5776549AE2551E67AB8|88907c3bfc1f34f5f2cdaa956ab78706; __gads=ID=e99f5b64396fb6fa-2202e93cb8e70017:T=1690792557:RT=1690792557:S=ALNI_MbhbusJuJJwMW7HKDaZjK11zcczBQ; __gpi=UID=00000d348852e513:T=1690792557:RT=1690792557:S=ALNI_MYdQIC4MT58yzX766DkwwNYEBWXVQ; __utmt_douban=1; __utmt=1; __utmb=30149280.6.10.1690791617`)
	req.Header.Set("Pragma", "no-cache")
	req.Header.Set("Referer", "https://www.douban.com/")
	req.Header.Set("Sec-Fetch-Dest", "document")
	req.Header.Set("Sec-Fetch-Mode", "navigate")
	req.Header.Set("Sec-Fetch-Site", "same-site")
	req.Header.Set("Sec-Fetch-User", "?1")
	req.Header.Set("Upgrade-Insecure-Requests", "1")
	req.Header.Set("User-Agent", "Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/115.0.0.0 Safari/537.36")
	req.Header.Set("sec-ch-ua", `"Not/A)Brand";v="99", "Google Chrome";v="115", "Chromium";v="115"`)
	req.Header.Set("sec-ch-ua-mobile", "?0")
	req.Header.Set("sec-ch-ua-platform", `"Windows"`)
	resp, err := model.DFClient.Do(req)
	if err != nil {
		log.Fatal(err)
	}
	defer resp.Body.Close()
	bodyText, err := io.ReadAll(resp.Body)
	if err != nil {
		log.Fatal(err)
	}
	//fmt.Printf("%s\n", bodyText)

	var result []model.Item
	doc, err := htmlquery.Parse(bytes.NewReader(bodyText))
	if err != nil {
		log.Printf("Kr36 err:%v\n", err)
		return
	}
	nodes := htmlquery.Find(doc, "//*[@id=\"screening\"]/div[2]/ul/li")
	for _, node := range nodes {
		dataRate := htmlquery.SelectAttr(node, "data-rate")
		if dataRate == "" {
			dataRate = "暂无评分"
		}
		dataRegion := htmlquery.SelectAttr(node, "data-region")
		dataRelease := htmlquery.SelectAttr(node, "data-release")
		dataDuration := htmlquery.SelectAttr(node, "data-duration")
		dataDirector := htmlquery.SelectAttr(node, "data-director")
		dataActors := htmlquery.SelectAttr(node, "data-actors")
		description := fmt.Sprintf("%v %v %v\n%v %v", dataRegion, dataRelease, dataDuration, dataDirector, dataActors)

		//dataTrailer := htmlquery.SelectAttr(node, "data-trailer")
		dataTitle := htmlquery.SelectAttr(node, "data-title")

		link := htmlquery.SelectAttr(htmlquery.FindOne(node, "//ul/li[1]/a"), "href")
		//title := htmlquery.SelectAttr(htmlquery.FindOne(node, "//ul/li[1]/a/img"), "alt")

		if dataTitle != "" {
			//fmt.Println(dataTitle, dataRelease, dataRate, dataDuration, dataRegion, dataDirector, dataActors, link)
			result = append(result, model.Item{Name: dataTitle, Link: link, Description: description, Badge: dataRate})
		}
	}
	if len(result) >= model.Num {
		model.M["Douban0"] = result[:model.Num]
	} else {
		model.M["Douban0"] = result
	}
	model.M["Douban0"] = append(model.M["Douban0"], model.Item{Name: "更多", Link: "https://movie.douban.com/"})

	log.Println("Douban0 success!!")
	return
}
