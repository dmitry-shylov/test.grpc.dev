include .env
export

generate_grpc:
	protoc -I ecommerce ecommerce/product_info.proto \
		 --go_out=$(GO_OUT) \
         --go-grpc_out=$(GO_GRPC_OUT) \
         --plugin=protoc-gen-go-grpc=$(PLUGIN)

start_grpc_product_info_server:
	go run cmd/product_info_server.go

start_grpc_product_info_client:
	go run cmd/product_info_client.go
