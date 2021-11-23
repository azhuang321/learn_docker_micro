package main

import (
	"context"
	"demo/proto"
	"fmt"
	"github.com/micro/go-micro/v2"
)

func main() {
	service := micro.NewService(micro.Name("cap.imooc.client"))
	service.Init()

	capImooc := proto.NewCapService("cap.imooc.server", service.Client())

	res, err := capImooc.SayHello(context.TODO(), &proto.SayHelloRequest{Message: "world"})
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(res.Answer)
}
