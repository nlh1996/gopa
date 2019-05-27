package model

import (
	"log"
	"pachong/conn"
	"pachong/utils"
)

// Solution .
type Solution struct {
	Href    string
	Title   string
	Content string
}

// Insert .
func (obj *Solution) Insert() {
	_, err := conn.GetCol().InsertOne(utils.GetCtx(), obj)
	if err != nil {
		log.Println(err.Error())
	}
}
