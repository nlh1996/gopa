package proxy

import (
	"log"
	"net/http"
	"pachong/model"

	"github.com/PuerkitoBio/goquery"
)

// Xici get ip from xicidaili.com
func Xici() (result []*model.IP) {
	pollURL := "http://www.xicidaili.com/nn/"
	client := &http.Client{}
	req, err := http.NewRequest("GET", pollURL, nil)
	req.Header.Add("User-Agent", "Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/74.0.3729.169 Safari/537.36")
	req.Header.Add("If-None-Match", "W/\"25a308c48a1a3215afe70ed6aba0361e\"")
	req.Header.Add("Upgrade-Insecure-Requests", "1")
	resp, err := client.Do(req)
	if err != nil || resp.StatusCode != 200 {
		log.Println(err)
		return
	}
	defer resp.Body.Close()
	doc, err := goquery.NewDocumentFromReader(resp.Body)
	if err != nil {
		log.Println(err.Error())
		return
	}
	doc.Find("tr>td:nth-child(2)").Each(func(i int, selection *goquery.Selection) {
		ip := &model.IP{}
		ip.Data = selection.Text()
		ip.Data = ip.Data + ":" + selection.Next().Text()
		ip.Type1 = "http"
		result = append(result, ip)
	})
	return
}
