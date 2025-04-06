package core

import (
	"encoding/json"
	"fmt"
	"io"
	"message/model"
	"net/http"
	"strings"
)

func Huxiu() error {
	var err error
	err = HuxiuF("1")
	if err != nil {
		return err
	}
	err = HuxiuF("2")
	if err != nil {
		return err
	}
	return nil
}

func HuxiuF(time_type string) error {
	url := "https://api-article.huxiu.com/web/index/hotArticles"
	method := "POST"

	payload := strings.NewReader(fmt.Sprintf("platform=www&time_type=%v&pagesize=6", time_type))

	client := &http.Client{}
	req, err := http.NewRequest(method, url, payload)

	if err != nil {
		return err
	}
	req.Header.Add("accept", "application/json, text/plain, */*")
	req.Header.Add("accept-language", "zh-CN,zh;q=0.9")
	req.Header.Add("cache-control", "no-cache")
	req.Header.Add("content-type", "application/x-www-form-urlencoded")
	req.Header.Add("origin", "https://www.huxiu.com")
	req.Header.Add("pragma", "no-cache")
	req.Header.Add("priority", "u=1, i")
	req.Header.Add("referer", "https://www.huxiu.com/")
	req.Header.Add("sec-ch-ua", "\"Google Chrome\";v=\"135\", \"Not-A.Brand\";v=\"8\", \"Chromium\";v=\"135\"")
	req.Header.Add("sec-ch-ua-mobile", "?0")
	req.Header.Add("sec-ch-ua-platform", "\"Windows\"")
	req.Header.Add("sec-fetch-dest", "empty")
	req.Header.Add("sec-fetch-mode", "cors")
	req.Header.Add("sec-fetch-site", "same-site")
	req.Header.Add("user-agent", "Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/135.0.0.0 Safari/537.36")
	req.Header.Add("Cookie", "huxiu_analyzer_wcy_id=34eap1nd9ublcg4x48v")

	res, err := client.Do(req)
	if err != nil {
		return err
	}
	defer res.Body.Close()

	body, err := io.ReadAll(res.Body)
	if err != nil {
		return err
	}
	//fmt.Println(string(body))

	var v model.HuxiuStruct
	err = json.Unmarshal(body, &v)
	if err != nil {
		return err
	}

	var result []model.Item
	for _, item := range v.Data {
		result = append(result, model.Item{Name: item.Title, Link: strings.ReplaceAll(item.URL, `\`, ""), Description: "", Badge: fmt.Sprintf("%v评论%v收藏", item.CountInfo.TotalCommentNum, item.CountInfo.FavoriteNum)})
	}

	key := "Huxiu" + time_type
	if len(result) >= model.Num {
		model.M[key] = result[:model.Num]
	} else {
		model.M[key] = result
	}
	model.M[key] = append(model.M[key], model.Item{Name: "更多", Link: "https://www.huxiu.com/"})

	return nil
}
