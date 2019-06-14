package weather

import (
	"fmt"
	"pachong/client"
)

const (
	index = "http://www.weather.com.cn/data/cityinfo/101010100.html"
	db    = "Weather"
	col   = "weather"
)

// Init .
func Init() {
	res, err := client.GetResponse(index)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(res.Body)
}
