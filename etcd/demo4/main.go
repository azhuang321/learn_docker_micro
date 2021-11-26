package main

import (
	"context"
	"fmt"
	"go.etcd.io/etcd/client/v3"
	"time"
)

func main() {
	var (
		config  clientv3.Config
		client  *clientv3.Client
		err     error
		getResp *clientv3.GetResponse
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

	//用于读写/cron/jobs 下所有的键值
	kv := clientv3.NewKV(client)

	kv.Put(context.TODO(), "/cron/jobs", "test")

	if getResp, err = kv.Get(context.TODO(), "/cron/jobs", clientv3.WithPrefix()); err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(getResp.Kvs, getResp.Count)
}
