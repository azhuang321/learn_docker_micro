package main

import (
	"context"
	"fmt"
	"go.etcd.io/etcd/client/v3"
	"time"
)

func main() {
	var (
		config       clientv3.Config
		client       *clientv3.Client
		lease        clientv3.Lease
		leaseResp    *clientv3.LeaseGrantResponse
		putResp      *clientv3.PutResponse
		getResp      *clientv3.GetResponse
		keepRespChan <-chan *clientv3.LeaseKeepAliveResponse
		keepResp     *clientv3.LeaseKeepAliveResponse
		err          error
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

	//申请一个less 续租
	lease = clientv3.NewLease(client)

	//申请10s的续租
	if leaseResp, err = lease.Grant(context.TODO(), 10); err != nil {
		fmt.Println(err)
		return
	}

	//拿到less id
	leaseId := leaseResp.ID

	// 5s 后自动取消续租,相当于 先续约了5s 然后10s 后过期  总共存在15s
	//ctx,_ := context.WithTimeout(context.TODO(),5 * time.Second)
	//if keepRespChan,err = lease.KeepAlive(ctx,leaseId); err != nil {
	//	fmt.Println(err)
	//	return
	//}

	//自动续租
	if keepRespChan, err = lease.KeepAlive(context.TODO(), leaseId); err != nil {
		fmt.Println(err)
		return
	}

	go func() {
		for {
			select {
			case keepResp = <-keepRespChan:
				if keepResp == nil {
					fmt.Println("租约已失效")
					goto END
				} else {
					fmt.Println("收到租约自动续约", keepResp.ID)
				}
			}
		}
	END:
	}()

	//获取kv api
	kv := clientv3.NewKV(client)

	// 写入key\value 与 租约关联起来
	if putResp, err = kv.Put(context.TODO(), "/cron/lock/job1", "", clientv3.WithLease(leaseId)); err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println("写入成功", putResp.Header.Revision)

	//定时查看是否过期

	for {
		if getResp, err = kv.Get(context.TODO(), "/cron/lock/job1"); err != nil {
			fmt.Println(err)
			return
		}

		if getResp.Count == 0 {
			fmt.Println("过期了")
			break
		}
		fmt.Println("还没过期", getResp.Kvs)
		time.Sleep(time.Second * 2)
	}
}
