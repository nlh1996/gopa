package gamersky

import (
	"fmt"
	"pachong/client"
	"pachong/conn"
	"pachong/model"
	"regexp"
	"sync"

	"github.com/PuerkitoBio/goquery"
)

const (
	index = "https://www.gamersky.com/news/"
	db    = "GamerSky"
	col   = "news"
)

// Init .
func Init() {
	conn.SetDB(db)
	conn.SetCol(col)
	doc, err := client.GetDocument(index)
	if err != nil {
		//logs.SendLog(err, "", 5)
		fmt.Println(err)
	}
	var newsList []string
	getNewsList(doc, &newsList)
	var wg sync.WaitGroup
	for i := range newsList {
		wg.Add(1)
		go getNews(newsList[i], &wg)
	}
	wg.Wait()
}

// 获取所有新闻链接
func getNewsList(doc *goquery.Document, newsList *[]string) {
	// '//a[@class="tt"]/@href'
	doc.Find("a[class=tt]").Each(func(i int, selection *goquery.Selection) {
		url, _ := selection.Attr("href")
		*newsList = append(*newsList, url)
	})
}

// 爬取新闻内容
func getNews(url string, wg *sync.WaitGroup) {
	doc, err := client.GetDocument(url)
	if err != nil {
		//logs.SendLog(err, "", 5)
		fmt.Println(err)
		wg.Done()
		return
	}

	news := &model.News{}
	news.URL = url
	news.Media = "GameSky"

	doc.Find("div[class=Mid2L_tit]>h1").Each(func(i int, selection *goquery.Selection) {
		news.Title = selection.Text()
	})

	if news.Title == "" {
		wg.Done()
		return
	}

	doc.Find("div[class=Mid2L_con]>p").Each(func(i int, selection *goquery.Selection) {
		news.Content = news.Content + selection.Text()
	})

	var tmpTime string
	doc.Find("div[class=detail]").Each(func(i int, selection *goquery.Selection) {
		tmpTime = selection.Text()
	})
	reg := regexp.MustCompile(`\S+`)
	timeString := reg.FindAllString(tmpTime, -1)
	news.PubTime = fmt.Sprintf("%s %s", timeString[0], timeString[1])
	news.Insert()
	wg.Done()
}
