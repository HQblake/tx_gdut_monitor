protoc -I . --go_out=paths=source_relative:./gen manage.proto

protoc -I . --go-grpc_out=paths=source_relative:./gen manage.proto