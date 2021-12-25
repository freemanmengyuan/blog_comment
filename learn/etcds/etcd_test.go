package etcds

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
