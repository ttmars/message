package core

import (
	"encoding/json"
	"io/ioutil"
	"log"
	"message/model"
	"net/http"
	"strings"
)

func Juejin() {
	FF("https://juejin.cn/?sort=three_days_hottest", "3")
	FF("https://juejin.cn/?sort=weekly_hottest", "7")
	FF("https://juejin.cn/?sort=monthly_hottest", "30")
}

func FF(url, k string) {
	str := `{"id_type":2,"client_type":2608,"sort_type":` + k + `,"cursor":"0","limit":20}`
	req, err := http.NewRequest("POST", "https://api.juejin.cn/recommend_api/v1/article/recommend_all_feed?aid=2608&uuid=7078976481116390915", strings.NewReader(str))
	if err != nil {
		log.Printf("juejin err:%v\n", err)
		return
	}
	req.Header.Set("content-type", "application/json")
	resp, err := model.DFClient.Do(req)
	if err != nil {
		log.Printf("juejin err:%v\n", err)
		return
	}
	defer resp.Body.Close()
	bodyText, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Printf("juejin err:%v\n", err)
		return
	}

	v := &model.JuejinStruct{}
	err = json.Unmarshal(bodyText, v)
	var result []model.Item
	for _,i := range v.Data {
		title := i.ItemInfo.ArticleInfo.Title
		link := "https://juejin.cn/post/" + i.ItemInfo.ArticleInfo.ArticleID
		//description := i.ItemInfo.ArticleInfo.BriefContent
		if title != ""{
			result = append(result, model.Item{Name: title, Link: link})
		}
	}
	name := "Juejin" + k
	if len(result) >= model.Num {
		model.M[name] = result[:model.Num]
	}
	model.M[name] = append(model.M[name], model.Item{Name: "更多", Link: url})
	return
}