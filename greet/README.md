protoc --go_out=plugins=grpc:. *.proto

    --go_out=plugins=grpc:.： 告诉 protoc 编译器使用 protoc-gen-go 插件来生成与 gRPC 兼容的 Go 接口代码，并保存到当前目录 . 下
    *.proto： 表明在当前目录下查找以 .proto 结尾的 Protocol Buffers 文件
