package model

import (
	"log"
	"pachong/conn"
	"pachong/utils"

	"go.mongodb.org/mongo-driver/bson"
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
	return results
}

// Del 从数据库中删除
func (obj *IP) Del() {
	del, err := conn.GetCol().DeleteOne(utils.GetCtx(), bson.D{{"data", obj.Data}})
	if err != nil {
		log.Println(err.Error(), del)
	}
}

// Count 统计ip数量
func (obj *IP) Count() int64 {
	num, err := conn.GetCol().CountDocuments(utils.GetCtx(), bson.M{})
	if err != nil {
		log.Fatal(err)
	}
	return num
}
