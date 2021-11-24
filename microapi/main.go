package main

import (
	_ "github.com/micro/go-plugins/registry/consul/v2"
	"github.com/micro/micro/v2/cmd"
)

/*
micro v2 不再支持 --registry consul
执行：go build -o micro
再使用 micro --registry=consul --registry_address=127.0.0.1:8500 --api_handler=api api...执行命令
*/

func main() {
	cmd.Init()
}
