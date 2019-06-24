package client

import (
	"bufio"
	"fmt"
	"os"
	"pachong/utils"
	"sync"
)

// var mutex = sync.Mutex{}

// DownLoadImg 图片下载，需要图片地址和图片的名字。
func DownLoadImg(imgURL string, fileName string, wg *sync.WaitGroup) {
	defer	wg.Done()
	res, err := GetResponse(imgURL)
	if err != nil {
		fmt.Println(err)
		return
	}
	path := utils.GetCurrentDirectory()
	path = path + "/download/"
	// 获得get请求响应的reader对象
	reader := bufio.NewReaderSize(res.Body, 32*1024)
	// 创建目录
	if err := os.MkdirAll(path, os.ModePerm); err != nil {
		fmt.Println(err)
	}
	file, err := os.Create(path + fileName)
	if err != nil {
		fmt.Println(err)
		return
	}
	bytes := make([]byte, 128*1024)
	len, err := reader.Read(bytes)
	if len < 0 || err != nil {
		return
	}
	// 注意这里byte数组后的[0:len]，不然可能会导致写入多余的数据
	_, _ = file.Write(bytes[:len])
	// 写入文件
	// written, _ := io.Copy(writer, reader)
	fmt.Println(path + fileName)
}
