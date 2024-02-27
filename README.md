# grpc-experiments
gRPC experiments

## Install

```sh
go install google.golang.org/grpc/cmd/protoc-gen-go-grpc@latest

# sudo apt install -y protobuf-compiler protoc-gen-go
```

## Run

Use the protoc compiler to generate Go code from the proto file:

```sh
protoc --go_out=. --go_opt=paths=source_relative \
    --go-grpc_out=. --go-grpc_opt=paths=source_relative \
    sum/sum.proto

```
