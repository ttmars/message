package core

import (
	"github.com/chromedp/cdproto/cdp"
	"github.com/chromedp/chromedp"
	"golang.org/x/net/context"
	"message/model"
	"time"
)

func Douyin() {
	// 自定义浏览器选项
	options := append(chromedp.DefaultExecAllocatorOptions[:],
		chromedp.Flag("headless", true),
		chromedp.Flag("hide-scrollbars", false),
		chromedp.Flag("mute-audio", false),
		chromedp.UserAgent(`Mozilla/5.0 (Windows NT 6.3; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/73.0.3683.103 Safari/537.36`),
	)
	ctx, cancel := chromedp.NewExecAllocator(context.Background(), options...)
	defer cancel()

	// 创建实例
	ctx, cancel = chromedp.NewContext(ctx) // chromedp.WithDebugf(log.Printf),

	defer cancel()

	// 设置超时
	ctx, cancel = context.WithTimeout(ctx, 30*time.Second)
	defer cancel()

	var nodes []*cdp.Node
	// 执行动作
	chromedp.Run(ctx,
		// 登录
		chromedp.Navigate(`https://www.douyin.com/hot`),

		// 重新获取流量信息
		chromedp.Sleep(time.Second*10),

		chromedp.Click(`#login-pannel > div.dy-account-close`), // 关闭登录弹窗
		chromedp.Sleep(time.Second*5),
		chromedp.Nodes(`//*[@id="douyin-right-container"]/div[3]/div/div[3]/ul/li/div[2]/div[1]/a/h3`, &nodes),
	)

	//fmt.Println("hello")
	//fmt.Println(len(nodes))
	//time.Sleep(time.Hour)

	var result []model.Item
	for _, node := range nodes {
		title := node.Children[0].NodeValue
		link := "https://www.douyin.com" + node.Parent.Attributes[1]
		result = append(result, model.Item{Name: title, Link: link})
		//fmt.Println(title, link)
	}
	if len(result) >= model.Num {
		model.M["DY"] = result[:model.Num]
	} else {
		model.M["DY"] = result
	}
	model.M["DY"] = append(model.M["DY"], model.Item{Name: "更多", Link: "https://www.douyin.com/hot"})
	return
}
