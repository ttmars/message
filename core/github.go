package core

import (
	"bytes"
	"github.com/antchfx/htmlquery"
	"io"
	"log"
	"log/slog"
	"message/model"
	"strconv"
	"strings"
	"sync"
)

var wg1 = &sync.WaitGroup{}

func Github() {
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
		"https://github.com/trending/java?since=daily",
		"https://github.com/trending/java?since=weekly",
		"https://github.com/trending/java?since=monthly",
		"https://github.com/trending/rust?since=daily",
		"https://github.com/trending/rust?since=weekly",
		"https://github.com/trending/rust?since=monthly",
		"https://github.com/trending/c++?since=daily",
		"https://github.com/trending/c++?since=weekly",
		"https://github.com/trending/c++?since=monthly",
		"https://github.com/trending/javascript?since=daily",
		"https://github.com/trending/javascript?since=weekly",
		"https://github.com/trending/javascript?since=monthly",
		"https://github.com/trending/shell?since=daily",
		"https://github.com/trending/shell?since=weekly",
		"https://github.com/trending/shell?since=monthly",
		"https://github.com/trending/assembly?since=daily",
		"https://github.com/trending/assembly?since=weekly",
		"https://github.com/trending/assembly?since=monthly",
		"https://github.com/trending/c%23?since=daily",
		"https://github.com/trending/c%23?since=weekly",
		"https://github.com/trending/c%23?since=monthly",
		"https://github.com/trending/php?since=daily",
		"https://github.com/trending/php?since=weekly",
		"https://github.com/trending/php?since=monthly",
	}
	for i, url := range urls {
		//log.Println("debug", i, url)
		wg1.Add(1)
		key := "Github" + strconv.Itoa(i)
		go F(url, key)
	}
	//log.Println("wait##########################################################################")
	wg1.Wait()
	//log.Println("finish##########################################################################")
}

func F(url string, k string) {
	defer func() {
		if err := recover(); err != nil {
			slog.Error("panic occurred", "error", err)
		}
	}()

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

	doc, err := htmlquery.Parse(bytes.NewReader(bodyText))
	if err != nil {
		log.Printf("github err:%v\n", err)
		return
	}
	nodes := htmlquery.Find(doc, "//article")
	for _, node := range nodes {
		n1, _ := htmlquery.Query(node, "//h2/a") // h1改为h2
		n2, _ := htmlquery.Query(node, "//p")
		n3, _ := htmlquery.Query(node, "//div[last()]/span[last()]")
		star := strings.TrimSpace(htmlquery.InnerText(n3))
		des := ""
		if n2 != nil {
			des = strings.TrimSpace(htmlquery.InnerText(n2))
		}
		title := htmlquery.SelectAttr(n1, "href")
		link := "https://github.com" + title

		if len(title) > 0 {
			title = title[1:]
		}
		result = append(result, model.Item{Name: title, Link: link, Description: des, Badge: star})
	}
	if len(result) >= model.GithubNum {
		model.M[k] = result[:model.GithubNum]
	} else {
		model.M[k] = result
	}
	model.M[k] = append(model.M[k], model.Item{Name: "更多", Link: url})
	log.Println("github success!!")
}
