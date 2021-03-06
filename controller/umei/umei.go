package umei

import (
	"fmt"
	"pachong/client"
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
			tmp := reg.FindAllString(urlList[i], -1)
			fileName := fmt.Sprintf("%s.%s", tmp[8], tmp[9])
			client.DownLoadImg(urlList[i], fileName, &wg)
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
