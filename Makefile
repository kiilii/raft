rpc:
	protoc --proto_path=/usr/local/Cellar/protobuf/21.5/include \
				--proto_path=./proto \
				--go_out=paths=source_relative:./proto \
				--go-grpc_out=paths=source_relative:./proto \

.PHONY: rpc
