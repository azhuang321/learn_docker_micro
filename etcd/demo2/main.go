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
		putResp *clientv3.PutResponse
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

	//用于读写etcd的键值对
	kv := clientv3.NewKV(client)
	if putResp, err = kv.Put(context.TODO(), "/cron/jobs/job1", "world", clientv3.WithPrevKV()); err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println("revision:", putResp.Header.Revision)
	if putResp.PrevKv != nil {
		fmt.Println("prevalue", string(putResp.PrevKv.Value))
	}

}
