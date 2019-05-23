package client

import (
	"fmt"
	"net/http"

	"github.com/PuerkitoBio/goquery"
)

// Request 客户端发起请求,返回html对象
func Request(url string) (*goquery.Document, error) {
	res, err := http.Get(url)
	if err == nil && res.StatusCode == 200 {
		html, err := goquery.NewDocumentFromReader(res.Body)
		res.Body.Close()
		return html, err
	}
	fmt.Println(res.StatusCode)
	return nil, err
}
