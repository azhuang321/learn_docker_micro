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
		delResp *clientv3.DeleteResponse
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
	if delResp, err = kv.Delete(context.TODO(), "/cron/jobs/job1", clientv3.WithPrevKV()); err != nil {
		fmt.Println(err)
		return
	}
	// 删除多个值
	if delResp, err = kv.Delete(context.TODO(), "/cron/jobs/", clientv3.WithPrefix()); err != nil {
		fmt.Println(err)
		return
	}

	//从什么开始删除,删除两个值
	if delResp, err = kv.Delete(context.TODO(), "/cron/jobs/", clientv3.WithFromKey(), clientv3.WithLimit(2)); err != nil {
		fmt.Println(err)
		return
	}

	if len(delResp.PrevKvs) != 0 {
		for _, kvpair := range delResp.PrevKvs {
			fmt.Println("deleted:", kvpair.Value)
		}
	}

}
