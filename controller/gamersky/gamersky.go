package gamersky

import (
	"fmt"
	"log"
	"pachong/client"
	"pachong/conn"
	"pachong/model"
	"regexp"
	"time"

	"github.com/PuerkitoBio/goquery"
)

const (
	index = "https://www.gamersky.com/news/"
	db    = "GamerSky"
	col		= "news"
)

// Init .
func Init() {
	conn.SetDB(db)
	conn.SetCol(col)
	doc, err := client.Request(index)
	if err != nil {
		log.Println(err)
		return
	}
	var newsList []string
	getNewsList(doc, &newsList, "a[class=tt]")
	for i := range newsList {
		getNews(newsList[i])
		time.Sleep(200 * time.Millisecond)
	}
}

// 获取所有新闻链接
func getNewsList(doc *goquery.Document, newsList *[]string, str string) {
	// '//a[@class="tt"]/@href'
	doc.Find(str).Each(func(i int, selection *goquery.Selection) {
		url, _ := selection.Attr("href")
		*newsList = append(*newsList, url)
	})
}

// 爬取新闻内容
func getNews(url string) {
	doc, err := client.Request(url)
	if err != nil {
		log.Println(err)
		return
	}

	news := &model.News{}
	news.URL = url
	news.Media = "GameSky"

	doc.Find("div[class=Mid2L_tit]>h1").Each(func(i int, selection *goquery.Selection) {
		news.Title = selection.Text()
	})

	if news.Title == "" {
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
}
