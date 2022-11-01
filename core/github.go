package core

import (
	"bytes"
	"github.com/antchfx/htmlquery"
	"io"
	"log"
	"message/model"
	"strconv"
	"strings"
	"sync"
)

var wg1 = &sync.WaitGroup{}

func Github()  {
	urls := []string{"https://github.com/trending/go?since=daily",
					"https://github.com/trending/go?since=weekly",
					"https://github.com/trending/go?since=monthly",
					"https://github.com/trending?since=daily",
					"https://github.com/trending?since=weekly",
					"https://github.com/trending?since=monthly",
					"https://github.com/trending/python?since=daily",
					"https://github.com/trending/python?since=weekly",
					"https://github.com/trending/python?since=monthly",
					"https://github.com/trending/c?since=daily",
					"https://github.com/trending/c?since=weekly",
					"https://github.com/trending/c?since=monthly",
	}
	for i,url := range urls{
		//log.Println("debug", i, url)
		wg1.Add(1)
		key := "Github" + strconv.Itoa(i)
		go F(url, key)
	}
	//log.Println("wait##########################################################################")
	wg1.Wait()
	//log.Println("finish##########################################################################")
}

func F(url string, k string)  {
	defer wg1.Done()
	var result []model.Item
	resp, err := model.DFClient.Get(url)
	if err != nil {
		log.Printf("github err:%v\n", err)
		return
	}
	defer resp.Body.Close()
	bodyText, err := io.ReadAll(resp.Body)
	if err != nil {
		log.Printf("github err:%v\n", err)
		return
	}

	doc,err := htmlquery.Parse(bytes.NewReader(bodyText))
	if err != nil {
		log.Printf("github err:%v\n", err)
		return
	}
	nodes := htmlquery.Find(doc, "//article")
	for _,node := range nodes{
		n1,_ := htmlquery.Query(node, "//h1/a")
		n2,_ := htmlquery.Query(node, "//p")
		n3,_ := htmlquery.Query(node, "//div[last()]/span[last()]")
		star := strings.TrimSpace(htmlquery.InnerText(n3))
		des := ""
		if n2 != nil {
			des = strings.TrimSpace(htmlquery.InnerText(n2))
		}
		title := htmlquery.SelectAttr(n1, "href")
		link := "https://github.com" + title
		result = append(result, model.Item{Name: title[1:], Link: link, Description: des, Badge: star})
	}
	if len(result) >= model.GithubNum {
		model.M[k] = result[:model.GithubNum]
	}else{
		model.M[k] = result
	}
	model.M[k] = append(model.M[k], model.Item{Name: "更多", Link: url})
	log.Println("github success!!")
}