package mongodb

import (
	"context"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"log"
	"testing"
)

type FindByName struct {
	JobName string `bson:jobName`
}

func TestFindLimit(t *testing.T) {
	var (
		client     *mongo.Client
		err        error
		dataBase   *mongo.Database
		collection *mongo.Collection
		cond *FindByName
		cursor *mongo.Cursor
		record *LogRecord
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

	// 4.按照jobName字段过滤，找出记录
	cond = &FindByName{
		JobName: "job1",
	}
	// 5.查询 (过滤加翻页)
	var skip,limit int64
	if cursor, err = collection.Find(context.TODO(), cond, &options.FindOptions{
		Skip: &skip,
		Limit: &limit,
	}); err != nil {
		t.Log()
	}
	defer cursor.Close(context.TODO())

	// 6.遍历结果集
	for cursor.Next(context.TODO()) {
		record = &LogRecord{}

		// 反序列化bson到对象
		if err = cursor.Decode(record); err != nil {
			t.Log(err)
		}
		t.Log(*record)
	}
}