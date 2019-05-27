package client

import (
	"log"
	"pachong/conf"
	"pachong/conn"
	"pachong/model"
	"pachong/proxy"
	"sync"
	"time"

	"github.com/PuerkitoBio/goquery"
	"github.com/parnurzeal/gorequest"
)

// Request 客户端发起请求,返回doc对象
func Request(url string) (*goquery.Document, error) {
	agent := gorequest.New()
	agent.Client.Timeout = 15 * time.Second
	tempIP := <-proxy.IPCh
	defer func() { proxy.IPCh <- tempIP }()
	ip := "http://" + tempIP.Data
	var (
		res  gorequest.Response
		errs []error
	)
	for res == nil {
		res, _, errs = agent.Proxy(ip).Get(url).End()
		if errs != nil {
			log.Println(errs[0])
			proxy.IPCh <- tempIP
			tempIP = <-proxy.IPCh
			ip = "http://" + tempIP.Data
		}
	}
	if res.StatusCode == 200 {
		doc, err := goquery.NewDocumentFromReader(res.Body)
		res.Body.Close()
		return doc, err
	}
	res.Body.Close()
	err := model.NewError("unkonw error!!!")
	return nil, err
}

// CheckIP 检查ip代理池的有效ip,
func CheckIP(ips []*model.IP) {
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
				pollURL = conf.CheckURL
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
				ips[i].Speed = time.Now().Sub(begin).Nanoseconds() / 1000 / 1000 //ms
				if ips[i].Speed < 1000 {
					proxy.IPCh <- ips[i]
					ips[i].Insert()
				}
			}
			wg.Done()
		}(i)
	}
	wg.Wait()
}
