LOCAL_BIN:=$(CURDIR)/bin

install-golangci-lint:
	go install github.com/golangci/golangci-lint/cmd/golangci-lint@v1.53.3

lint:
	golangci-lint run ./... --config .golangci.pipeline.yaml

install-deps:
	set GOBIN=$(LOCAL_BIN) && go install google.golang.org/protobuf/cmd/protoc-gen-go@latest
	set GOBIN=$(LOCAL_BIN) && go install -mod=mod google.golang.org/grpc/cmd/protoc-gen-go-grpc@latest

generate: 
	make generate-note-api

generate-note-api:
	protoc --go_out=. --go-grpc_out=. api\user_api_v1\user_api.proto