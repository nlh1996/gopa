package proxy

import (
	"log"
	"pachong/model"
	"sync"
)

// Init 初始化ip代理池,返回所有代理ip对象
func Init() []*model.IP {
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
			//log.Println("[run] get into loop")
			for _, v := range temp {
				//log.Println("[run] len of ipChan %v",v)
				ips = append(ips, v)
			}
			wg.Done()
		}(f)
	}
	wg.Wait()
	log.Println("All getters finished.")
	return ips
}
