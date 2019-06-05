package main

import (
	"pachong/conn"
	"pachong/controller/gamersky"
	"pachong/proxy"
)

func main() {
	conn.Init()
	proxy.Init()
	if proxy.Count() < 100 {
		// 抓取最新的代理ip
		ips := proxy.Get()
		proxy.CheckIP(ips)
	} else {
		// 不抓取直接使用数据库中的ip
		proxy.CheckDBIP()
	}
	// 爬虫demo
	gamersky.Init()
	// hoperun.Init()
	// umei.Init()
}
