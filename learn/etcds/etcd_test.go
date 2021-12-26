package etcds

import (
	"context"
	"go.etcd.io/etcd/clientv3"
	"testing"
	"time"
)

/**
 * etcd基础
 * 抽屉理论 reft协议 大部分原则 支持事物 mvcc
 * 搭建单机etcd,熟悉命令行操作
 * goland调用etcd的put get delete lease watch
 * 使用txn事物功能，实现分布式乐观锁
 */
// 命令行示例
// 启动
// nohup ./etcd --listen-client-urls 'http://0.0.0.0:2379' --advertise-client-urls 'http://0.0.0.0:2379' >myout.log  2>&1 &
// 设置值
// ETCDCTL_API=3 ./etcdctl put "mengyuan" "12" 使用v3版接口
// 按目录查找
// ETCDCTL_API=3 ./etcdctl put "/cron/jobs/job1" "{...job1}"
// ETCDCTL_API=3 ./etcdctl put "/cron/jobs/job2" "{...job2}"
// ETCDCTL_API=3 ./etcdctl get "/cron/jobs/" --prefix
// 监听
// ETCDCTL_API=3 ./etcdctl watch "/cron/jobs/" --prefix

// 程序链接
// 下载包
// go get go.etcd.io/etcd/clientv3   也可以借助golang中国来下载
// 如果运行失败 replace google.golang.org/grpc => google.golang.org/grpc v1.26.0
func TestClient(t *testing.T) {
	t.Log("hello")
	var (
		config clientv3.Config
		err error
		cli *clientv3.Client
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
	t.Log(cli)
}

// 添加
func TestKVPut(t *testing.T) {
	var (
		config clientv3.Config
		err error
		cli *clientv3.Client
		kv clientv3.KV
		putResp *clientv3.PutResponse
	)
	// 客户端配置
	config = clientv3.Config{
		Endpoints: []string{"176.122.158.130:2379"},
		DialTimeout: 5*time.Second,
	}
	// 建立一个客户端
	if cli,err = clientv3.New(config); err != nil{
		t.Log(err)
		return
	}
	// 用于读写etcd的键值对
	kv = clientv3.NewKV(cli)
	if putResp, err = kv.Put(context.TODO(), "/cron/jobs/job4", "client test4", clientv3.WithPrevKV()); err != nil {
		t.Log(err)
	} else {
		t.Log("Revision:", putResp.Header.Revision)
		if putResp.PrevKv != nil {
			t.Log("PrevValue:", string(putResp.PrevKv.Value))
		}
	}
}

// 获取
func TestKVGet(t *testing.T) {
	var (
		config clientv3.Config
		err error
		cli *clientv3.Client
		kv clientv3.KV
		getResp *clientv3.GetResponse
	)
	// 客户端配置
	config = clientv3.Config{
		Endpoints: []string{"176.122.158.130:2379"},
		DialTimeout: 5*time.Second,
	}
	// 建立一个客户端
	if cli,err = clientv3.New(config); err != nil{
		t.Log(err)
		return
	}
	// 用于读写etcd的键值对
	kv = clientv3.NewKV(cli)
	// 读取单个键值
	// if getResp, err = kv.Get(context.TODO(), "/cron/jobs/job3" /*clientv3.WithCountOnly()*/); err != nil {
	//	t.Log(err)
	//} else {
	//	t.Log("Value:", getResp.Kvs, getResp.Count)
	//}
	// 读取/cron/jobs/为前缀的所有key
	if getResp, err = kv.Get(context.TODO(), "/cron/jobs/", clientv3.WithPrefix()); err != nil {
		t.Log(err)
	} else {
		t.Log("Value:", getResp.Kvs, getResp.Count)
	}
}

// 删除
func TestKVDel(t *testing.T) {
	var (
		config clientv3.Config
		err error
		cli *clientv3.Client
		kv clientv3.KV
		delResp *clientv3.DeleteResponse
	)
	// 客户端配置
	config = clientv3.Config{
		Endpoints: []string{"176.122.158.130:2379"},
		DialTimeout: 5*time.Second,
	}
	// 建立一个客户端
	if cli,err = clientv3.New(config); err != nil{
		t.Log(err)
		return
	}
	// 用于读写etcd的键值对
	kv = clientv3.NewKV(cli)
	// 删除KV
	// options
	//clientv3.WithPrefix() 删除指定前缀的键值
	// clientv3.WithFromKey()  clientv3.WithLimit() 删除从指定键开始的固定个数的key
	if delResp, err = kv.Delete(context.TODO(), "/cron/jobs/job4", clientv3.WithPrevKV()); err != nil {
		t.Log(err)
	}
	//被删除之前value是什么
	if len(delResp.PrevKvs) > 0 {
		for _, kv := range delResp.PrevKvs {
			t.Log("删除了：",string(kv.Key), string(kv.Value))
		}
	}
}