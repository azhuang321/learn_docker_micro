
GOPATH:=$(shell go env GOPATH)
.PHONY: init
init:
	go get -u github.com/golang/protobuf/proto
	go get -u github.com/golang/protobuf/protoc-gen-go
	go get github.com/micro/micro/v3/cmd/protoc-gen-micro
	go get github.com/micro/micro/v3/cmd/protoc-gen-openapi

.PHONY: api
api:
	protoc --openapi_out=. --proto_path=. proto/cartApi.proto

.PHONY: proto
proto:
	protoc --proto_path=. --micro_out=proto/cartApi/ --go_out=:proto/cartApi/ proto/cartApi/cartApi.proto
	protoc --proto_path=. --micro_out=proto/cart/ --go_out=:proto/cart/ proto/cart/cart.proto

.PHONY: build
build:
	CGO_ENABLE=0 GOOS=linux GOARCH=amd64 go build -tags netgo -o cartApi main.go

.PHONY: test
test:
	go test -v ./... -cover

.PHONY: docker
docker:
	docker build -t cartApi:latest .

.PHONY: run
run:
	docker rm cartApi && docker build -t cartApi:latest . && docker run --name cartApi cartApi
