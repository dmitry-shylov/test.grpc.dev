generate_grpc:
	protoc -I ecommerce ecommerce/product_info.proto \
		 --go_out={go_out_path} \
         --go-grpc_out={go-grpc_out} \
         --plugin=protoc-gen-go-grpc={plugin}

start_grpc_product_info_server:
	go run cmd/product_info_server.go

start_grpc_product_info_client:
	go run cmd/product_info_client.go
