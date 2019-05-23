package model

import (
	"log"
	"pachong/conn"
	"pachong/utils"
)

// News 新闻数据
type News struct {
	Title   string
	Media   string
	URL     string
	PubTime string
	Content string
}

// Insert 插入新闻到数据库
func (obj *News) Insert() {
	_, err := conn.GetCol().InsertOne(utils.GetCtx(), obj)
	if err != nil {
		log.Println(err.Error())
	}
}
