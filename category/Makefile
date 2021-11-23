
GOPATH:=$(shell go env GOPATH)
.PHONY: init
init:
	go get -u github.com/golang/protobuf/proto
	go get -u github.com/golang/protobuf/protoc-gen-go
	go get github.com/micro/micro/v3/cmd/protoc-gen-micro
	go get github.com/micro/micro/v3/cmd/protoc-gen-openapi

.PHONY: api
api:
	protoc --openapi_out=. --proto_path=. proto/category.proto

.PHONY: proto
proto:
	protoc --proto_path=. --micro_out=proto/category/ --go_out=:proto/category/ proto/category/category.proto

.PHONY: build
build:
	CGO_ENABLE=0 GOOS=linux GOARCH=amd64 go build -tags netgo -o category main.go

.PHONY: test
test:
	go test -v ./... -cover

.PHONY: docker
docker:
	docker build -t category:latest .

.PHONY: run
run:
	docker rm category && docker build -t category:latest . && docker run --name category category
