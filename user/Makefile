GOPATH:=$(shell go env GOPATH)
.PHONY: init
init:
	go get -u github.com/golang/protobuf/proto
	go get -u github.com/golang/protobuf/protoc-gen-go
	go get github.com/micro/micro/v3/cmd/protoc-gen-micro
	go get github.com/micro/micro/v3/cmd/protoc-gen-openapi

.PHONY: api
api:
	protoc --openapi_out=. --proto_path=. proto/user.proto

.PHONY: proto
proto:
	protoc --proto_path=. --micro_out=proto/user/ --go_out=:proto/user/ proto/user/user.proto

.PHONY: build
build:
	CGO_ENABLE=0 GOOS=linux GOARCH=amd64 go build -tags netgo -o user main.go

.PHONY: test
test:
	go test -v ./... -cover

.PHONY: docker
docker:
	docker build -t user:latest .

.PHONY: run
run:
	docker rm user && docker build -t user:latest . && docker run --name user user