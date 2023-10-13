package core

import (
	"crypto/tls"
	"encoding/json"
	"fmt"
	"io"
	"log"
	"log/slog"
	"message/model"
	"net/http"
)

func Tieba() {
	defer func() {
		if err := recover(); err != nil {
			slog.Error("panic occurred", "error", err)
		}
	}()

	tr := &http.Transport{
		TLSClientConfig: &tls.Config{InsecureSkipVerify: true},
	}
	client := &http.Client{Transport: tr}
	req, err := http.NewRequest("GET", "http://tieba.baidu.com/hottopic/browse/topicList", nil)
	if err != nil {
		log.Fatal(err)
	}
	req.Header.Set("Accept", "text/html,application/xhtml+xml,application/xml;q=0.9,image/avif,image/webp,image/apng,*/*;q=0.8,application/signed-exchange;v=b3;q=0.7")
	req.Header.Set("Accept-Language", "zh-CN,zh-TW;q=0.9,zh;q=0.8,en-US;q=0.7,en;q=0.6")
	req.Header.Set("Cache-Control", "no-cache")
	req.Header.Set("Connection", "keep-alive")
	req.Header.Set("Cookie", `XFI=7a394e50-3010-11ee-8f01-eb2b513a1d32; XFCS=5F8DCC5C4BE7E0C3028A53929A0D8EBB42A244D54C59A1D3E1707CB3C0ACBD46; XFT=LL/FEwUyCfiDl3tJo7uZZK8Qv0iy1AyfD9l+9mnCr28=; BIDUPSID=CF54A566E88BCA0E7F4A12E4C0179BB0; PSTM=1688352369; BAIDUID=CF54A566E88BCA0E640A4192E4F47D32:FG=1; BDORZ=B490B5EBF6F3CD402E515D22BCDA1598; BAIDU_WISE_UID=wapp_1690791582030_369; arialoadData=false; RT="z=1&dm=baidu.com&si=8ac0bbaf-f786-4f63-b590-f6c1eb9117ce&ss=lkqlnh9p&sl=5&tt=2mg&bcn=https%3A%2F%2Ffclog.baidu.com%2Flog%2Fweirwood%3Ftype%3Dperf&ld=586&ul=qi8&hd=qtx"; USER_JUMP=-1; XFI=1a817b00-300f-11ee-97a2-5d91025b4eac; XFCS=5BF17D35A1A387E78C63A0AFB0BFFFD5A7021F9B67F83445EE6E38E870E379B8; XFT=kXqobedRPoN3X1pgZgw/y7foGRyGbC7LE0vpY0UWt90=; Hm_lvt_98b9d8c2fd6608d564bf2ac2ae642948=1690791582,1690855196; Hm_lpvt_98b9d8c2fd6608d564bf2ac2ae642948=1690855782`)
	req.Header.Set("Pragma", "no-cache")
	req.Header.Set("Upgrade-Insecure-Requests", "1")
	req.Header.Set("User-Agent", "Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/115.0.0.0 Safari/537.36")
	resp, err := client.Do(req)
	if err != nil {
		log.Fatal(err)
	}
	defer resp.Body.Close()
	bodyText, err := io.ReadAll(resp.Body)
	if err != nil {
		log.Fatal(err)
	}
	//fmt.Printf("%s\n", bodyText)

	var v model.TiebaStruct
	json.Unmarshal(bodyText, &v)

	var result []model.Item
	for _, item := range v.Data.BangTopic.TopicList {
		//fmt.Println(item)
		result = append(result, model.Item{Name: item.TopicName, Link: item.TopicURL, Description: item.TopicDesc, Badge: fmt.Sprintf("%.1fw", float64(item.DiscussNum)/10000)})
	}

	key := "tieba"
	if len(result) >= model.Num {
		model.M[key] = result[:model.Num]
	} else {
		model.M[key] = result
	}
	model.M[key] = append(model.M[key], model.Item{Name: "更多", Link: "https://tieba.baidu.com/hottopic/browse/topicList?res_type=1"})

	log.Println("Tieba success!!")
	return
}
