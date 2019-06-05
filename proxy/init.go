package proxy

import (
	"log"
	"pachong/conf"
	"pachong/conn"
	"pachong/model"
	"sync"
	"time"

	"github.com/parnurzeal/gorequest"
)

// IPCh 使用通道作为本爬虫框架的代理ip池，可全局调用.
var IPCh chan *model.IP

const (
	db  = "IPTABLE"
	col = "ips"
)

// Init 爬取所有的代理ip
func Init() []*model.IP {
	IPCh = make(chan *model.IP, 100)
	var wg sync.WaitGroup
	funs := []func() []*model.IP{
		//Data5u,
		Feiyi,
		//IP66, //need to remove it
		KDL,
		//GBJ,	//因为网站限制，无法正常下载数据
		Xici,
		//XDL,
		//IP181,  // 已经无法使用
		//YDL,	//失效的采集脚本，用作系统容错实验
		PLP, //need to remove it
		IP89,
	}
	var ips []*model.IP
	for _, f := range funs {
		wg.Add(1)
		go func(f func() []*model.IP) {
			temp := f()
			for _, v := range temp {
				ips = append(ips, v)
			}
			wg.Done()
		}(f)
	}
	wg.Wait()
	log.Println("All getters finished.")
	return ips
}

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
					IPCh <- ips[i]
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
					IPCh <- ips[i]
					ips[i].Insert()
				}
			}
			wg.Done()
		}(i)
	}
	wg.Wait()
}
