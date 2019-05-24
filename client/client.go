package client

import (
	"net/http"
	"pachong/conn"
	"pachong/model"
	"sync"
	"time"

	"github.com/PuerkitoBio/goquery"
	"github.com/parnurzeal/gorequest"
)

// Request 客户端发起请求,返回doc对象
func Request(url string) (*goquery.Document, error) {
	res, err := http.Get(url)
	if err != nil {
		return nil, err
	}
	if res.StatusCode == 200 {
		doc, err := goquery.NewDocumentFromReader(res.Body)
		res.Body.Close()
		return doc, err
	}
	return nil, err
}

// CheckIP 检查ip代理池的有效ip,
func CheckIP(ips []*model.IP, ipCh chan *model.IP) {
	const (
		db  = "IPTABLE"
		col = "ips"
	)
	conn.SetDB(db)
	conn.SetCol(col)
	var wg sync.WaitGroup
	len := len(ips)
	wg.Add(len)
	for i := 0; i < len; i++ {
		go func(i int) {
			const (
				pollURL = "http://httpbin.org/get"
			)
			var testIP string
			if ips[i].Type2 == "https" {
				testIP = "https://" + ips[i].Data
			} else {
				testIP = "http://" + ips[i].Data
			}

			begin := time.Now()
			agent := gorequest.New()
			// 设置2s超时时间，以防长时间不响应
			agent.Client.Timeout = 2 * time.Second
			resp, _, errs := agent.Proxy(testIP).Get(pollURL).End()
			if errs != nil {
				wg.Done()
				return
			}
			defer resp.Body.Close()
			if resp.StatusCode == 200 {
				//harrybi 20180815 判断返回的数据格式合法性
				// _, err := sj.NewFromReader(resp.Body)
				// if err != nil {
				// 	fmt.Println(testIP, pollURL, err)
				// 	return
				// }
				//harrybi 计算该代理的速度，单位毫秒
				ips[i].Speed = time.Now().Sub(begin).Nanoseconds() / 1000 / 1000 //ms
				if ips[i].Speed < 1000 {
					ipCh <- ips[i]
					ips[i].Insert()
				}
			}
			wg.Done()
		}(i)
	}
	wg.Wait()
}
