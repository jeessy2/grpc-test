# A simple test for golang grpc
[reference]https://grpc.io/docs/quickstart/go/
## Steps
1. go mod init grpc-test
2. [download protoc-3.11.4-win64.zip]https://github.com/protocolbuffers/protobuf/releases
3. move protoc.exe to c:\Go\bin
4. go get github.com/golang/protobuf/protoc-gen-go
5. protoc --go_out=plugins=grpc:./proto ./proto/simple.proto
6. go run server/main.go
7. go run clint/main.go