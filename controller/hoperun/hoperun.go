package hoperun

import (
	"fmt"
	"pachong/client"
	"pachong/conn"

	"github.com/PuerkitoBio/goquery"
)

const (
	index = "http://www.hoperun.com"
	db    = "HopeRun"
	col   = "index"
)

// Init .
func Init() {
	conn.SetDB(db)
	conn.SetCol(col)
	doc, err := client.Request(index)
	if err != nil {
		fmt.Println(err)
	}
	var urlList []string
	getUrlList(doc, &urlList)
	for i := range urlList {
		fmt.Println(urlList[i])
	}
}

func getUrlList(doc *goquery.Document, urlList *[]string) {
	doc.Find("a").Each(func(i int, selection *goquery.Selection) {
		url, _ := selection.Attr("href")
		*urlList = append(*urlList, url)
	})
}
