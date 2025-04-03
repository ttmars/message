package core

import (
	"bytes"
	"github.com/antchfx/htmlquery"
	"io/ioutil"
	"message/model"
	"strconv"
	"time"
)

func Kr36() error {
	t := time.Now().Format("2006-01-02")
	urls := []string{
		"https://www.36kr.com/hot-list/renqi/" + t + "/1",
		"https://www.36kr.com/hot-list/zonghe/" + t + "/1",
		"https://www.36kr.com/hot-list/shoucang/" + t + "/1",
	}
	for i, url := range urls {
		k := "Kr36" + strconv.Itoa(i)
		err := FFF(url, k)
		if err != nil {
			return err
		}
	}
	return nil
}

func FFF(url, k string) error {
	resp, err := model.DFClient.Get(url)
	if err != nil {
		return err
	}
	defer resp.Body.Close()
	bodyText, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return err
	}

	var result []model.Item
	doc, err := htmlquery.Parse(bytes.NewReader(bodyText))
	if err != nil {
		return err
	}
	nodes := htmlquery.Find(doc, "//*[@id=\"app\"]/div/div[2]/div[3]/div/div/div[2]/div[1]/div/div/div/div/div[2]/div[2]/p/a")
	for _, node := range nodes {
		title := htmlquery.InnerText(node)
		link := "https://www.36kr.com" + htmlquery.SelectAttr(node, "href")
		result = append(result, model.Item{Name: title, Link: link})
	}
	if len(result) >= model.Num {
		model.M[k] = result[:model.Num]
	} else {
		model.M[k] = result
	}
	model.M[k] = append(model.M[k], model.Item{Name: "更多", Link: url})
	return nil
}
