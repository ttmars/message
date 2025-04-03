package core

import (
	"bytes"
	"fmt"
	"github.com/antchfx/htmlquery"
	"io"
	"message/model"
	"net/http"
)

func ITHome() error {
	client := &http.Client{}
	req, err := http.NewRequest("GET", "https://www.ithome.com/", nil)
	if err != nil {
		return err
	}
	req.Header.Set("authority", "www.ithome.com")
	req.Header.Set("accept", "text/html,application/xhtml+xml,application/xml;q=0.9,image/avif,image/webp,image/apng,*/*;q=0.8,application/signed-exchange;v=b3;q=0.7")
	req.Header.Set("accept-language", "zh-CN,zh-TW;q=0.9,zh;q=0.8,en-US;q=0.7,en;q=0.6")
	req.Header.Set("cache-control", "no-cache")
	req.Header.Set("cookie", "Hm_lvt_f2d5cbe611513efcf95b7f62b934c619=1690791318; Hm_lvt_cfebe79b2c367c4b89b285f412bf9867=1690791228,1690852535; Hm_lpvt_cfebe79b2c367c4b89b285f412bf9867=1690852535")
	req.Header.Set("pragma", "no-cache")
	req.Header.Set("sec-ch-ua", `"Not/A)Brand";v="99", "Google Chrome";v="115", "Chromium";v="115"`)
	req.Header.Set("sec-ch-ua-mobile", "?0")
	req.Header.Set("sec-ch-ua-platform", `"Windows"`)
	req.Header.Set("sec-fetch-dest", "document")
	req.Header.Set("sec-fetch-mode", "navigate")
	req.Header.Set("sec-fetch-site", "none")
	req.Header.Set("sec-fetch-user", "?1")
	req.Header.Set("upgrade-insecure-requests", "1")
	req.Header.Set("user-agent", "Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/115.0.0.0 Safari/537.36")
	resp, err := client.Do(req)
	if err != nil {
		return err
	}
	defer resp.Body.Close()
	bodyText, err := io.ReadAll(resp.Body)
	if err != nil {
		return err
	}

	var result0 []model.Item
	var result1 []model.Item
	var result2 []model.Item
	doc, err := htmlquery.Parse(bytes.NewReader(bodyText))
	if err != nil {
		return err
	}
	nodes0 := htmlquery.Find(doc, "//*[@id=\"d-1\"]/li/a")
	nodes1 := htmlquery.Find(doc, "//*[@id=\"d-2\"]/li/a")
	nodes2 := htmlquery.Find(doc, "//*[@id=\"d-3\"]/li/a")

	for _, node := range nodes0 {
		title := htmlquery.SelectAttr(node, "title")
		link := htmlquery.SelectAttr(node, "href")
		if title != "" {
			result0 = append(result0, model.Item{Name: title, Link: link})
		}
	}

	for _, node := range nodes1 {
		title := htmlquery.SelectAttr(node, "title")
		link := htmlquery.SelectAttr(node, "href")
		if title != "" {
			result1 = append(result1, model.Item{Name: title, Link: link})
		}
	}

	for _, node := range nodes2 {
		title := htmlquery.SelectAttr(node, "title")
		link := htmlquery.SelectAttr(node, "href")
		if title != "" {
			result2 = append(result2, model.Item{Name: title, Link: link})
		}
	}

	for i := 0; i <= 2; i++ {
		var result []model.Item
		if i == 0 {
			result = result0
		} else if i == 1 {
			result = result1
		} else {
			result = result2
		}
		key := fmt.Sprintf("ithome%v", i)
		if len(result) >= model.Num {
			model.M[key] = result[:model.Num]
		} else {
			model.M[key] = result
		}
		model.M[key] = append(model.M[key], model.Item{Name: "更多", Link: "https://www.ithome.com/"})
	}

	return err
}
