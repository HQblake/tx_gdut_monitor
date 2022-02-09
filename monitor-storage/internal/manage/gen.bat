@echo off
protoc -I . --go_out=paths=source_relative:./managepb manage2store.proto
protoc -I . --go-grpc_out=paths=source_relative:./managepb manage2store.proto