package mongodb

import (
	"context"
	"fmt"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"log"
	"testing"
	"time"
)



func TestInsertMany(t *testing.T){
	var (
		client *mongo.Client
		err error
		dataBase *mongo.Database
		collection  *mongo.Collection
		ret *mongo.InsertManyResult
	)
	// 1.建立链接
	clientOptions := options.Client().ApplyURI("mongodb://176.122.158.130:27017")
	// Connect to MongoDB
	client, err = mongo.Connect(context.TODO(), clientOptions)
	if err != nil {
		log.Fatal(err)
	}
	// 2.选择数据库
	dataBase = client.Database("cron")
	// 3.选择collection
	collection = dataBase.Collection("log")

	// 4.插入记录
	var insertMany []interface{}
	for i:= 0; i<50; i++ {
		record := &LogRecord{
			JobName: fmt.Sprintf("job%d", i),
			Command: fmt.Sprintf("echo hello %d", i),
			Err: "",
			Content: "test",
			TimePoint: TimePoint{StartTime:time.Now().Unix() , EndTime: time.Now().Unix()+10},
		}
		insertMany = append(insertMany, record)
	}
	if ret, err = collection.InsertMany(context.TODO(), insertMany); err != nil {
		t.Log(err)
	}
	// _id: 默认生成一个全局唯一ID, ObjectID：12字节的二进制
	for _,insertId := range ret.InsertedIDs {
		t.Log("自增id:", insertId.(primitive.ObjectID).Hex())
	}
}