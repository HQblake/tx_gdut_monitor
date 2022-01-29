protoc -I . --go_out=paths=source_relative:./gen send.proto

protoc -I . --go-grpc_out=paths=source_relative:./gen send.proto