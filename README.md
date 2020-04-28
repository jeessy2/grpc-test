# Golang grpc
https://grpc.io/docs/quickstart/go/
 0. go mod init grpc-test
 0. [download protoc-3.11.4-win64.zip]https://github.com/protocolbuffers/protobuf/releases
 0. move protoc.exe to c:\Go\bin
 0. go get github.com/golang/protobuf/protoc-gen-go
 0. protoc --go_out=plugins=grpc:./proto ./proto/simple.proto
 0. go run server/main.go
 0. go run clint/main.go