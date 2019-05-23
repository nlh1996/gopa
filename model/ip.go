package model

import (
	"log"
	"pachong/conn"
	"pachong/utils"
)

// IP struct
type IP struct {
	Data  string
	Type1 string
	Type2 string
	Speed int64
}


// Insert 插入ip到数据库
func (obj *IP) Insert() {
	_, err := conn.GetCol().InsertOne(utils.GetCtx(), obj)
	if err != nil {
		log.Println(err.Error())
	}
}
