package client

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"pachong/conf"
)

// DownLoadImg 图片下载.
func DownLoadImg(imgURL string, fileName string) {
	res, err := GetResponse(imgURL)
	if err != nil {
		fmt.Println(err)
		return
	}
	// 获得get请求响应的reader对象
	reader := bufio.NewReaderSize(res.Body, 32*1024)
	// 判断目录是否存在
	info, err := os.Stat(conf.ImgPath)
	if os.IsNotExist(err) {
		// 创建目录
		if err := os.MkdirAll(conf.ImgPath, os.ModePerm); err != nil {
			fmt.Println(err,info)
		}
	}
	file, err := os.Create(conf.ImgPath + fileName)
	if err != nil {
		fmt.Println(err)
		return
	}
	// 获得文件的writer对象
	writer := bufio.NewWriter(file)
	// copy写入文件
	written, _ := io.Copy(writer, reader)
	fmt.Printf("Total length: %d", written)
}
