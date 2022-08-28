
init: 
	go install google.golang.org/protobuf/cmd/protoc-gen-go@latest
	go install google.golang.org/grpc/cmd/protoc-gen-go-grpc@latest
	go install github.com/gogo/protobuf/proto@latest

rpc:
	protoc --proto_path=/usr/local/Cellar/protobuf/21.5/include \
				--proto_path=./proto \
				--gogo_out=paths=source_relative:./proto \
				--go-grpc_out=paths=source_relative:./proto \
				./proto/*.proto

.PHONY: rpc
