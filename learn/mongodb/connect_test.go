package mongodb

import (
	"context"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"testing"
)

/**
 * 1.安装mongodb 参考菜鸟教程
 * 2.常用的命令
 *	mongod --dbpath /var/lib/mongo --logpath /var/log/mongodb/mongod.log --bind_ip=0.0.0.0 --fork
 * 3.安装go扩展
 *  go get -u -v go.mongodb.org/mongo-driver
 */

func TestConnect(t *testing.T){
	var (
		client *mongo.Client
		err error
		dataBase *mongo.Database
		collection  *mongo.Collection
	)
	// 1.建立链接
	op := &options.ClientOptions{
		Hosts: []string{"mongodb://176.122.158.130:27017"},
		// ConnectTimeout:,
	}
	if client, err = mongo.Connect(context.TODO(), op); err != nil {
		t.Log(err)
	}
	// 2.选择数据库
	dataBase = client.Database("my_db")
	// 3.选择collection
	collection = dataBase.Collection("my_collection")
	collection = collection
}