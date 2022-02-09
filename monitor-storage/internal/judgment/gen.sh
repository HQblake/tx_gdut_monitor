protoc -I . --go_out=paths=source_relative:./judgpb judgment2store.proto

protoc -I . --go-grpc_out=paths=source_relative:./judgpb judgment2store.proto