package main

import (
	"context"
	"fmt"
	"go.etcd.io/etcd/client/v3"
	"time"
)

func main() {
	var (
		config clientv3.Config
		client *clientv3.Client
		opResp clientv3.OpResponse
		err    error
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

	kv := clientv3.NewKV(client)

	putOp := clientv3.OpPut("/cron/jobs/job8", "")

	//执行OP
	if opResp, err = kv.Do(context.TODO(), putOp); err != nil {
		fmt.Println(err)
		return
	}

	fmt.Println("写入revision:", opResp.Put().Header.Revision)

	getOp := clientv3.OpGet("/cron/jobs/job8")

	//执行OP
	if opResp, err = kv.Do(context.TODO(), getOp); err != nil {
		fmt.Println(err)
		return
	}

	fmt.Println("数据revision:", opResp.Get().Kvs[0].ModRevision) //create rev == mod rev
	fmt.Println("数据value:", string(opResp.Get().Kvs[0].Value))

}
