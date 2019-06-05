package client

import (
	"log"
	"pachong/proxy"
	"time"

	"github.com/PuerkitoBio/goquery"
	"github.com/parnurzeal/gorequest"
)

// GetDocument 使用代理ip发起get请求,返回doc对象
func GetDocument(url string) (*goquery.Document, error) {
	agent := gorequest.New()
	agent.Client.Timeout = 10 * time.Second
	tempIP := <-proxy.IPCh
	defer func() { proxy.IPCh <- tempIP }()
	ip := "http://" + tempIP.Data
	var (
		res   gorequest.Response
		errs  []error
		index int
	)

	// 更换代理ip重复请求,直到请求成功或者超过请求限制（防止请求死循环）
	for res == nil || res.StatusCode != 200 {
		index++
		res, _, errs = agent.Proxy(ip).Get(url).End()
		if errs != nil {
			proxy.IPCh <- tempIP
			tempIP = <-proxy.IPCh
			ip = "http://" + tempIP.Data
			if index > 5 {
				return nil, errs[0]
			}
		}
	}
	res.Body.Close()
	doc, err := goquery.NewDocumentFromReader(res.Body)
	if err != nil {
		return nil, err
	}
	return doc, nil
}

// GetResponse 使用代理ip发起请求，返回response.
func GetResponse(url string) (gorequest.Response, error) {
	agent := gorequest.New()
	agent.Client.Timeout = 30 * time.Second
	tempIP := <-proxy.IPCh
	defer func() { proxy.IPCh <- tempIP }()
	ip := "http://" + tempIP.Data
	var (
		res   gorequest.Response
		errs  []error
		index int
	)
	// 更换代理ip重复请求,直到请求成功或者超过请求限制（防止请求死循环）
	for res == nil || res.StatusCode != 200 {
		index++
		res, _, errs = agent.Proxy(ip).Get(url).End()
		if errs != nil {
			proxy.IPCh <- tempIP
			tempIP = <-proxy.IPCh
			ip = "http://" + tempIP.Data
			if index > 5 {
				return nil, errs[0]
			}
		}
	}
	res.Body.Close()
	return res, nil
}

// Get 使用本机ip发起请求,返回doc对象
func Get(url string) (*goquery.Document, error) {
	agent := gorequest.New()
	agent.Client.Timeout = 10 * time.Second
	res, _, errs := agent.Get(url).End()
	if errs != nil {
		log.Fatal(errs[0])
	}
	res.Body.Close()
	doc, err := goquery.NewDocumentFromReader(res.Body)
	if err != nil {
		return nil, err
	}
	return doc, nil
}
