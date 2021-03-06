package conn

import (
	"context"
	"log" 
	"pachong/conf"
	"time"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

// Mgo .
type Mgo struct {
	client *mongo.Client
	db     *mongo.Database
	col    *mongo.Collection
}

var mgo *Mgo

// Init 数据库连接.
func Init() {
	mgo = &Mgo{}
	var err error
	mgo.client, err = mongo.NewClient(options.Client().ApplyURI(conf.MgoURL))
	ctxWithTimeout, _ := context.WithTimeout(context.Background(), 5*time.Second)
	err = mgo.client.Connect(ctxWithTimeout)
	if err != nil {
		log.Fatal(err)
	}
	// Check the connection
	err = mgo.client.Ping(ctxWithTimeout, nil)
	if err != nil {
		log.Fatal(err)
	}
	log.Println("Connected to MongoDB!")
}

// GetCol .
func GetCol() *mongo.Collection {
	return mgo.col
}

// SetDB .
func SetDB(str string) {
	if mgo.client == nil {
		return
	}
	mgo.db = mgo.client.Database(str)
}

// SetCol .
func SetCol(str string) {
	if mgo.db == nil {
		return
	}
	mgo.col = mgo.db.Collection(str)
}
