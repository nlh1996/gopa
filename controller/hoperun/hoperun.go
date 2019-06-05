package hoperun

import (
	"fmt"
	"pachong/client"
	"pachong/conn"
	"pachong/model"
	"sync"

	"github.com/PuerkitoBio/goquery"
)

const (
	index = "http://www.hoperun.com"
	db    = "HopeRun"
	col   = "data"
)

// Init .
func Init() {
	conn.SetDB(db)
	conn.SetCol(col)
	doc, err := client.GetDocument(index)
	if err != nil {
		fmt.Println(err)
		return
	}
	var urlList []string
	getUrlList(doc, &urlList)
	len := len(urlList)
	var wg sync.WaitGroup
	wg.Add(len)
	for i := 0; i < len; i++ {
		go getNews(urlList[i], &wg)
	}
	wg.Wait()
}

func getUrlList(doc *goquery.Document, urlList *[]string) {
	doc.Find("div[class=nav_nl]>dl>dd>a").Each(func(i int, selection *goquery.Selection) {
		url, _ := selection.Attr("href")
		*urlList = append(*urlList, url)
	})
}

func getNews(url string, wg *sync.WaitGroup) {
	if url[:4] != "http" {
		url = "http://www.hoperun.com" + url
	} else {
		wg.Done()
		return
	}
	doc, err := client.GetDocument(url)
	if err != nil {
		fmt.Println(err)
		wg.Done()
		return
	}

	doc.Find("div[class=fa_c]>ul>li>a").Each(func(i int, selection *goquery.Selection) {
		data := &model.Solution{}
		data.Href, _ = selection.Attr("href")
		data.Title = selection.Text()
		data.Content = selection.Next().Text()
		data.Insert()
	})

	wg.Done()
}
