.PHONY: generate-client

generate-client:
	protoc -I rpc/ rpc/login.proto --go_out=../rpc/ --go-grpc_out=../rpc/
	protoc -I rpc/ rpc/tokenValidate.proto --go_out=../rpc/ --go-grpc_out=../rpc/