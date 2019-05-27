package proxy

import (
	"log"
	"pachong/model"
	"sync"
)

// IPCh 使用通道作为本爬虫框架的代理ip池，可全局调用.
var IPCh chan *model.IP

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
		//Xici,
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
