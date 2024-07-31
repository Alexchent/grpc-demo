# 生成pb文件和grpc文件，输出到当前路径，相对与proto/person.proto，也就是输出在proto目录
.PHONY: p
proto:
	protoc --go_out=. --go_opt=paths=source_relative \
            --go-grpc_out=. --go-grpc_opt=paths=source_relative \
            proto/*.proto

# -go_out=. 等同于--go_out=paths=import:. 表示按照go_package的全路径创建目录层级
.PHONY: p2
p2:
	protoc --proto_path=. --go_out=. proto/person.proto
	#protoc --proto_path=. --go_out=paths=import:. proto1/greeter/greeter_v2.proto

# 启动grpc客户端
.PHONY: ui
ui:
	grpcui -plaintext localhost:50051

