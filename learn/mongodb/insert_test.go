package mongodb

import (
	"context"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"log"
	"testing"
	"time"
)


type TimePoint struct {
	StartTime int64 `bson:startTime`
	EndTime int64 `bson:endTime`
}

type LogRecord struct {
	JobName string `bson:jobName`
	Command string `bson:command`
	Err string `bson:err`
	Content string `bson:content`
	TimePoint TimePoint `bson:timePoint`
}

func TestInsert(t *testing.T){
	var (
		client *mongo.Client
		err error
		dataBase *mongo.Database
		collection  *mongo.Collection
		record LogRecord
		ret *mongo.InsertOneResult
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
	record = LogRecord{
		JobName: "job1",
		Command: "echo hello",
		Err: "",
		Content: "test",
		TimePoint: TimePoint{StartTime:time.Now().Unix() , EndTime: time.Now().Unix()+10},
	}
	if ret, err = collection.InsertOne(context.TODO(), &record); err != nil {
		t.Log(err)
	}
	// _id: 默认生成一个全局唯一ID, ObjectID：12字节的二进制
	docId := ret.InsertedID.(primitive.ObjectID)
	t.Log(docId.Hex())
}