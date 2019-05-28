package client

import (
	"log"
	"pachong/conf"
	"pachong/conn"
	"pachong/model"
	"pachong/proxy"
	"sync"
	"time"

	"github.com/parnurzeal/gorequest"
)

const (
	db  = "IPTABLE"
	col = "ips"
)

// CheckIP 检查ip代理池的有效ip,
func CheckIP(ips []*model.IP) {
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
			agent.Client.Timeout = 5 * time.Second
			resp, _, errs := agent.Proxy(testIP).Get(pollURL).End()
			if errs != nil {
				wg.Done()
				return
			}
			defer resp.Body.Close()
			if resp.StatusCode == 200 {
				ips[i].Speed = time.Now().Sub(begin).Nanoseconds() / 1000 / 1000 //ms
				if ips[i].Speed < 1500 {
					proxy.IPCh <- ips[i]
					ips[i].Insert()
				}
			}
			wg.Done()
		}(i)
	}
	wg.Wait()
}

// CheckDBIP 检查数据库中的有效ip
func CheckDBIP() {
	conn.SetDB(db)
	conn.SetCol(col)
	ip := &model.IP{}
	ips := ip.FindAll()
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
			agent.Client.Timeout = 5 * time.Second
			resp, _, errs := agent.Proxy(testIP).Get(pollURL).End()
			if errs != nil {
				wg.Done()
				log.Println(errs[0])
				return
			}
			defer resp.Body.Close()
			if resp.StatusCode == 200 {
				ips[i].Speed = time.Now().Sub(begin).Nanoseconds() / 1000 / 1000 //ms
				if ips[i].Speed < 1500 {
					proxy.IPCh <- ips[i]
					ips[i].Insert()
				}
			}
			wg.Done()
		}(i)
	}
	wg.Wait()
}
