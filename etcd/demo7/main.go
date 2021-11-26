package main

import (
	"context"
	"fmt"
	"go.etcd.io/etcd/api/v3/mvccpb"
	"go.etcd.io/etcd/client/v3"
	"time"
)

func main() {
	var (
		config  clientv3.Config
		client  *clientv3.Client
		getResp *clientv3.GetResponse
		err     error
	)

	//客户端配置
	config = clientv3.Config{
		Endpoints:   []string{"127.0.0.1:2379"},
		DialTimeout: 5 * time.Second,
	}
	//建立连接
	if client, err = clientv3.New(config); err != nil {
		fmt.Println(err)
		return
	}

	//kv
	kv := clientv3.NewKV(client)

	// 模拟kv的变化
	go func() {
		for {
			kv.Put(context.TODO(), "/cron/jobs/job7", "i am job7")
			kv.Delete(context.TODO(), "/cron/jobs/job7")
			time.Sleep(time.Second)
		}
	}()

	if getResp, err = kv.Get(context.TODO(), "/cron/jobs/job7"); err != nil {
		fmt.Println(err)
		return
	}

	//现在key是存在的
	if len(getResp.Kvs) != 0 {
		fmt.Println("当前值:", string(getResp.Kvs[0].Value))
	}

	//当前etcd集群事务id,单调递增
	watchStartRevision := getResp.Header.Revision + 1

	//创建一个watcher
	watcher := clientv3.NewWatcher(client)

	// 启动监听
	fmt.Println("从当前版本开始监听:", watchStartRevision)

	ctx, cancleFunc := context.WithCancel(context.TODO())
	time.AfterFunc(time.Second*5, func() {
		cancleFunc()
	})

	watchRespChan := watcher.Watch(ctx, "/cron/jobs/job7", clientv3.WithRev(watchStartRevision))

	//处理kv变化
	for watchResp := range watchRespChan {
		for _, event := range watchResp.Events {
			switch event.Type {
			case mvccpb.PUT:
				fmt.Println("修改为", string(event.Kv.Value), "revision:", event.Kv.CreateRevision)
			case mvccpb.DELETE:
				fmt.Println("修改为", string(event.Kv.Value), "revision:", event.Kv.CreateRevision)
			}
		}
	}
}
