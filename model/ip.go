package model

import (
	"log"
	"pachong/conn"
	"pachong/utils"

	"gopkg.in/mgo.v2/bson"
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

// FindAll 得到所有ip并且清空集合
func (obj *IP) FindAll() []*IP {
	var results []*IP
	col := conn.GetCol()
	cur, err := col.Find(utils.GetCtx(), bson.M{})
	if err != nil {
		log.Fatal(err)
	}

	for cur.Next(utils.GetCtx()) {
		ip := &IP{}
		if err := cur.Decode(&ip); err != nil {
			log.Fatal(err)
		}
		results = append(results, ip)
	}

	if err := cur.Err(); err != nil {
		log.Fatal(err)
	}

	del, err := col.DeleteMany(utils.GetCtx(), bson.M{})
	if err != nil {
			log.Fatal(err)
	}
	log.Println( del.DeletedCount)
	return results
}
