
GOPATH:=$(shell go env GOPATH)
.PHONY: init
init:
	go get -u github.com/golang/protobuf/proto
	go get -u github.com/golang/protobuf/protoc-gen-go
	go get github.com/micro/micro/v3/cmd/protoc-gen-micro
	go get github.com/micro/micro/v3/cmd/protoc-gen-openapi

.PHONY: api
api:
	protoc --openapi_out=. --proto_path=. proto/product.proto

.PHONY: proto
proto:
	protoc --proto_path=. --micro_out=proto/product/ --go_out=:proto/product/ proto/product/product.proto

.PHONY: build
build:
	CGO_ENABLE=0 GOOS=linux GOARCH=amd64 go build -tags netgo -o product main.go

.PHONY: test
test:
	go test -v ./... -cover

.PHONY: docker
docker:
	docker build -t product:latest .

.PHONY: run
run:
	docker rm product && docker build -t product:latest . && docker run --name product product
