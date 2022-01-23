package etcds

import (
	"context"
	"fmt"
	"go.etcd.io/etcd/clientv3"
	"testing"
	"time"
)

func TestWatchKey(t *testing.T){
	var (
		config clientv3.Config
		client *clientv3.Client
		err error
		kv clientv3.KV
		getResp *clientv3.GetResponse
		watchStartRevision int64
		watcher clientv3.Watcher
		ctx context.Context
		watchRespChan <-chan clientv3.WatchResponse
		watchResp clientv3.WatchResponse
		event *clientv3.Event
	)
	// 客户端配置
	config = clientv3.Config{
		Endpoints: []string{"176.122.158.130:2379"},
		DialTimeout: 5*time.Second,
	}

	// 建立链接
	if client, err = clientv3.New(config); err!= nil {
		t.Log(err)
		return
	}
	// Kv
	kv = clientv3.NewKV(client)

	// 模拟etcd的kv变化
	go func() {
		for {
			kv.Put(context.TODO(), "/cron/jobs/job7", "i am job7")
			kv.Delete(context.TODO(), "/cron/jobs/job7")
			time.Sleep(1*time.Second)
		}
	}()

	// 先Get到当前的值，并监听后续的变化
	if getResp,err = kv.Get(context.TODO(), "/cron/jobs/job7"); err!= nil {
		t.Log(err)
		return
	}
	// 现在的key是存在的
	if len(getResp.Kvs) > 0 {
		t.Log("当前值：", string(getResp.Kvs[0].Value))
	}
	// 当前etcd的集群的事物id，单调递增
	watchStartRevision = getResp.Header.Revision + 1

	// 创建一个watcher
	watcher = clientv3.NewWatcher(client)
	// 启动监听
	fmt.Println("从该版本向后监听：", watchStartRevision)

	ctx, _ = context.WithCancel(context.TODO())

	watchRespChan = watcher.Watch(ctx, "/cron/jobs/job7", clientv3.WithRev(watchStartRevision))

	for watchResp = range watchRespChan {
		for _, event = range  watchResp.Events {
			switch event.Type {
			case 1: // mvccpb.DELETE
				fmt.Println("删除了：",  "revision:", event.Kv.ModRevision)
			case 0: // mvccpb.PUT
				fmt.Println("修改为：", string(event.Kv.Value), "revision:", event.Kv.CreateRevision, event.Kv.ModRevision)
			}
		}
	}

}