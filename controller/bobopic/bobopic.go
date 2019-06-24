package bobopic

import (
	"fmt"
	"pachong/client"
	"regexp"
	"strconv"
	"sync"

	"github.com/PuerkitoBio/goquery"
)

const (
	index = "https://bobopic.com/page/"
)

// Init .
func Init() {
	var wg sync.WaitGroup
	ch := make(chan string, 100)
	for i := 1; i < 20; i++ {
		wg.Add(1)
		go getPages(i, &ch, &wg)
	}
	go func() {
		var index int
		reg := regexp.MustCompile(`[.][j][p][g]|[.][p][n][g]`)
		for {
			url, ok := <-ch
			if ok {
				index ++
				tmp := reg.FindAllString(url, -1)
				fileName := fmt.Sprintf("%d%s", index, tmp[0])
				wg.Add(1)
				go func() {
					client.DownLoadImg(url, fileName, &wg)
				}()
			}
		}
	}()

	wg.Wait()
	for {
		if len(ch) == 0{
			close(ch)
			return
		}
	}
}

func getPages(i int, ch *chan string, wg *sync.WaitGroup) {
	url := index + strconv.FormatInt(int64(i), 10)
	doc, err := client.GetDocument(url)
	if err != nil {
		fmt.Println(err)
		wg.Done()
		return
	}
	getUrlList(doc, ch)
	wg.Done()
}

func getUrlList(doc *goquery.Document, ch *chan string) {
	doc.Find("img").Each(func(i int, selection *goquery.Selection) {
		url, _ := selection.Attr("srcset")
		if url != "" {
			*ch <- url
		}
	})
	// doc.Find("article>div").Each(func(i int, selection *goquery.Selection) {
	// 	url, _ := selection.Attr("style")
	// 	if url != "" {
	// 		url = url[22:]
	// 		reg := regexp.MustCompile(`[^);]*`)
	// 		url := reg.FindAllString(url, -1)
	// 		*ch <- url[0]
	// 	}
	// })
}
