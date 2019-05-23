package main

import (
	"pachong/conn"
	"pachong/proxy"
)

func main() {
	conn.Init()
	ips := proxy.Init()
	// gamersky.Init()
	// hoperun.Init()

}
