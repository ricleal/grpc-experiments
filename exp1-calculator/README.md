## Run

Use the protoc compiler to generate Go code from the proto file:

```sh
protoc --go_out=common --go_opt=paths=source_relative \
    --go-grpc_out=common --go-grpc_opt=paths=source_relative \
    sum.proto

```