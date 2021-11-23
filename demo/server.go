package main

import (
	"context"
	"demo/proto"
	"fmt"
	"github.com/micro/go-micro/v2"
)

type CapServer struct{}

func (c *CapServer) SayHello(ctx context.Context, req *proto.SayHelloRequest, resp *proto.SayHelloResponse) error {
	resp.Answer = "hello" + req.Message
	return nil
}

func main() {
	//创建新的服务
	service := micro.NewService(
		micro.Name("cap.imooc.server"),
	)
	//初始方法
	service.Init()
	//注册服务
	proto.RegisterCapHandler(service.Server(), new(CapServer))
	//运行服务
	if err := service.Run(); err != nil {
		fmt.Println(err)
	}
}
