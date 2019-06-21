package client

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"pachong/utils"
	"sync"
)

// DownLoadImg 图片下载，需要图片地址和图片的名字。
func DownLoadImg(imgURL string, fileName string, wg *sync.WaitGroup) {
	res, err := GetResponse(imgURL)
	if err != nil {
		fmt.Println(err)
		wg.Done()
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
		wg.Done()
		return
	}
	// 获得文件的writer对象
	writer := bufio.NewWriter(file)
	// copy写入文件
	written, _ := io.Copy(writer, reader)
	fmt.Println(path+fileName+" Total length:", written)
	wg.Done()
}
