include .env
export

generate_grpc_product_info:
	protoc -I ecommerce ecommerce/product_info.proto \
		 --go_out=$(GO_OUT) \
         --go-grpc_out=$(GO_GRPC_OUT) \
         --plugin=protoc-gen-go-grpc=$(PLUGIN)

generate_grpc_bank_account_transaction:
	protoc -I ecommerce ecommerce/bank_account_transaction.proto \
		 --go_out=$(GO_OUT) \
         --go-grpc_out=$(GO_GRPC_OUT) \
         --plugin=protoc-gen-go-grpc=$(PLUGIN)

start_grpc_product_info_server:
	go run cmd/product_info_server.go

start_grpc_product_info_client:
	go run cmd/product_info_client.go

start_grpc_bank_account_transaction_server:
	go run cmd/bank_account_transaction_server.go

start_grpc_bank_account_transaction_client:
	go run cmd/bank_account_transaction_client.go