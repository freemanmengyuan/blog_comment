package etcds

import (
	"context"
	"go.etcd.io/etcd/clientv3"
	"testing"
	"time"
)

func TestOperation(t *testing.T){
	var (
		config clientv3.Config
		client *clientv3.Client
		err error
		kv clientv3.KV
		putOp clientv3.Op
		putResp clientv3.OpResponse
		getOp clientv3.Op
		getResp  clientv3.OpResponse
	)

	// 客户端配置
	config = clientv3.Config{
		Endpoints: []string{"176.122.158.130:2379"},
		DialTimeout: 5 * time.Second,
	}

	// 建立连接
	if client, err = clientv3.New(config); err != nil {
		t.Log(err)
		return
	}

	kv = clientv3.NewKV(client)

	// 创建op:operation
	putOp = clientv3.OpPut("/cron/jobs/job8", "12332323")
	// 执行op
	if putResp, err = kv.Do(context.TODO(), putOp); err!= nil {
		t.Log(err)
		return
	}

	// kv.Do(op)
	// kv.Put()
	// kv.Get()
	// kv.Delete()
	t.Log("写入revision:", putResp.Put().Header.Revision)

	// 创建op
	getOp = clientv3.OpGet("/cron/jobs/job8")
	// 执行op
	if getResp, err = kv.Do(context.TODO(), getOp); err != nil {
		t.Log(err)
		return
	}
	// 打印数据
	t.Log("数据Revision:", getResp.Get().Header.Revision)
	t.Log("数据value:", string(getResp.Get().Kvs[0].Value))
}