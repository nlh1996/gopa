package main

import (
	"fmt"
	"pachong/client"
	"pachong/conn"
	"pachong/model"
	"pachong/proxy"
)

func main() {
	conn.Init()
	ips := proxy.Init()
	ipCh := make(chan *model.IP, 100)
	client.CheckIP(ips, ipCh)
	for i := 0; i < len(ipCh); i++ {
		ip := <-ipCh
		fmt.Println(ip.Data)
	}
	// gamersky.Init()
	// hoperun.Init()
}
