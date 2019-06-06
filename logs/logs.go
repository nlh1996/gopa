package logs

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
	"pachong/conf"
	"time"
)

// Log .
type Log struct {
	Version string `json:"version"`
	Host    string `json:"host"`
	Level   uint8  `json:"level"`
	Info    string `json:"_some_info"`
	Msg     string `json:"short_message"`
}

var (
	log    Log
	client *http.Client
)

func init() {
	log = Log {
		Version: "1.1",
		Host:    "爬虫",
		Level:   0,
		Info:    "",
		Msg:     "",
	}
	client = &http.Client{Timeout: 10 * time.Second}
}

// SendLog 发送日志信息
func SendLog(err error, info string, lv uint8) {
	log.Msg = err.Error()
	log.Info = info
	log.Level = lv
	bytesData, _ := json.Marshal(log)
	reader := bytes.NewReader(bytesData)
	req, _ := http.NewRequest("POST", conf.LogURL, reader)
	req.Header.Set("content-type", "application/json")
	resp, err := client.Do(req)
	if err != nil {
		fmt.Println(err, resp)
	}
	resp.Body.Close()
}
