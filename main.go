package main

import (
	"pachong/client"
	"pachong/conn"
	"pachong/controller/gamersky"
	"pachong/proxy"
)

func main() {
	conn.Init()
	ips := proxy.Init()
	client.CheckIP(ips)
	gamersky.Init()
	// hoperun.Init()
}
