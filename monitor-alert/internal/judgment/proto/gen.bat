@echo off
protoc -I . --go_out=paths=source_relative:./judgpb manage2judgment.proto
protoc -I . --go-grpc_out=paths=source_relative:./judgpb manage2judgment.proto

protoc -I . --go_out=paths=source_relative:./managepb judgment2manage.proto
protoc -I . --go-grpc_out=paths=source_relative:./managepb judgment2manage.proto

protoc -I . --go_out=paths=source_relative:./storagepb judgment2store.proto
protoc -I . --go-grpc_out=paths=source_relative:./storagepb judgment2store.proto