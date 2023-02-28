# grpc-demo

> [协议缓冲区编译器安装](https://www.grpc.io/docs/protoc-installation/)


## 准备工作：

### 首先需要安装proto

1. 安装go [去下载](https://golang.google.cn/)
2. 安装协议缓存区编译器protobuf `brew install protobuf`
```
protoc --version # 确保版本为 3+
```
3. 安装protoc-gen-go-grpc `brew install protoc-gen-go-grpc`
4. 协议编译器的go插件
    - 使用以下命令为 Go 安装协议编译器插件
    ```bash
    go install google.golang.org/protobuf/cmd/protoc-gen-go@v1.26
    go install google.golang.org/grpc/cmd/protoc-gen-go-grpc@v1.1
    ```
    - 更新您的，PATH以便protoc编译器可以找到插件：
    ```bash
    export PATH="$PATH:$(go env GOPATH)/bin"
    ```

## 开始

### 1. 用 protocol buffer 定义服务 `helloworld.proto`

例子如下：
```protobuf
syntax = "proto3";

option go_package = "grpc/examples/helloworld";

package helloworld;

// The greeting service definition.
service Greeter {
  // Sends a greeting
  rpc SayHello (HelloRequest) returns (HelloReply) {}

  rpc SayHelloAgain (HelloRequest) returns (HelloReply) {}
}

// The request message containing the user's name.
message HelloRequest {
  string name = 1;
}

// The response message containing the greetings
message HelloReply {
  string message = 1;
}

```

### 2. 编译成 go 文件
```bash
protoc --go_out=. --go_opt=paths=source_relative \
    --go-grpc_out=. --go-grpc_opt=paths=source_relative \
    helloworld/helloworld.proto
```

说明：

`--go_out=` go文件的生成路径

`--go-grpc_out=` grpc-go文件的生成路径

目录的末尾是proto文件的路径

这会生成 `helloworld/helloworld.pb.go ` 和 `helloworld/helloworld_grpc.pb.go` 文件， 其中包含：
- Code for populating, serializing, and retrieving HelloRequest and HelloReply message types.
- 生成 client 和 server 代码


生成后的文件如下：
```
.
├── greeter_client
│   └── main.go
├── greeter_server
│   └── main.go
└── helloworld
    ├── helloworld.pb.go
    ├── helloworld.proto
    └── helloworld_grpc.pb.go
```