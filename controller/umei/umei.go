package umei

import (
	"fmt"
	"pachong/client"
	"pachong/conn"
	"regexp"
	"sync"

	"github.com/PuerkitoBio/goquery"
)

const (
	index = "http://www.umei.cc/meinvtupian/"
	db    = "UMei"
	col   = "imgs"
)

// Init .
func Init() {
	conn.SetDB(db)
	conn.SetCol(col)
	doc, err := client.Get(index)
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
		go func(i int) {
			reg := regexp.MustCompile(`\w+`)
			timeString := reg.FindAllString(urlList[i], -1)
			fileName := fmt.Sprintf("%s.%s", timeString[8], timeString[9])
			client.DownLoadImg(urlList[i], fileName)
			wg.Done()
		}(i)
	}
	wg.Wait()
}

func getUrlList(doc *goquery.Document, urlList *[]string) {
	doc.Find("img").Each(func(i int, selection *goquery.Selection) {
		url, _ := selection.Attr("src")
		*urlList = append(*urlList, url)
	})
}
