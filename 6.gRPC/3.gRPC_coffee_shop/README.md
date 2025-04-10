
# Run

```bash
# this generates pb files in the same directory as the proto file
protoc --go_opt=paths=source_relative --go-grpc_opt=paths=source_relative --go_out=. --go-grpc_out=. coffee_shop.proto

# this generates pb files in the ./coffee_shop_protos directory,
# but directory must be created first
protoc --go_opt=paths=source_relative --go-grpc_opt=paths=source_relative --go_out=./coffee_shop_protos/ --go-grpc_out=./coffee_shop_protos/ coffee_shop.proto

# this generates pb codes and creates directory if it doesn't exist
protoc --go_out=. --go-grpc_out=. coffee_shop.proto

```