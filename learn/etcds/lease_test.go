package etcds

import (
	"context"
	"go.etcd.io/etcd/clientv3"
	"testing"
	"time"
)

func TestLease(t *testing.T) {
	var (
		config clientv3.Config
		err error
		cli *clientv3.Client
		lease clientv3.Lease
		leaseGrantResp *clientv3.LeaseGrantResponse
		leaseId clientv3.LeaseID
		kv clientv3.KV
		putResq *clientv3.PutResponse
		getResp *clientv3.GetResponse
	)
	// 客户端配置
	config = clientv3.Config{
		Endpoints: []string{"176.122.158.130:2379"},
		DialTimeout: 5*time.Second,
	}
	// 建立链接
	if cli,err = clientv3.New(config); err != nil{
		t.Log(err)
		return
	}
	// 申请一个lease(租约)
	lease = clientv3.NewLease(cli)
	// 申请一个10秒的租约
	if leaseGrantResp, err = lease.Grant(context.TODO(), 10); err != nil {
		t.Log(err)
	}
	// 拿到租约的id
	leaseId = leaseGrantResp.ID

	// 用于读写etcd的键值对
	kv = clientv3.NewKV(cli)
	putResq, err = kv.Put(context.TODO(), "/cron/lock/job1", "", clientv3.WithLease(leaseId))
	if err != nil {
		t.Log(err)
	}
	t.Log("写入带租期的键值成功:", putResq.Header.Revision)
	for {
		if getResp, err = kv.Get(context.TODO(), "/cron/lock/job1"); err != nil {
			t.Log(err)
		}
		if getResp.Count >0 {
			t.Log("未过期", getResp.Kvs)
		} else{
			t.Log("已过期")
			break
		}
	}
}