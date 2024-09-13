.PHONY: gen help

PROTO_DIR=E:\minor\Go\src\go_code\rpcs\grpctranscode\proto

gen:
	protoc --proto_path=$(PROTO_DIR)  --go_out=$(PROTO_DIR) --go_opt=paths=source_relative --go-grpc_out=$(PROTO_DIR) --go-grpc_opt=paths=source_relative --grpc-gateway_out=$(PROTO_DIR) --grpc-gateway_opt=paths=source_relative $(PROTO_DIR)/book/books.proto

help:
	@echo "make gen - 生成pb及grpc代码"
