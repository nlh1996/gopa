package main

import (
	"pachong/conn"
	"pachong/controller/gamersky"
	"pachong/proxy"
)

func main() {
	conn.Init()
	ips := proxy.Init()
	/* 以下两种ip代理池二选一 */
	// 抓取最新的代理ip
	proxy.CheckIP(ips)

	// 使用数据库中的ip,并且去除失效的ip
	// proxy.CheckDBIP()

	// 爬虫demo
	gamersky.Init()
	//hoperun.Init()
	// umei.Init()

}
