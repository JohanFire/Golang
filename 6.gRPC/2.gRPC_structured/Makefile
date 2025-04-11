.PHONY: proto clean

# Generate protobuf files
proto:
# only for linux
# mkdir -p internal/protobuf
# protoc --go_out=. --go-grpc_out=. .\internal\proto\classification.proto

# # this generates files in route directory
# protoc --proto_path=internal/proto \
# 	--go_out=. \
# 	--go_opt=paths=source_relative \
# 	--go-grpc_out=. \
# 	--go-grpc_opt=paths=source_relative \
# 	internal/proto/classification.proto

# # windows
# protoc --proto_path=internal\proto --go_out=. --go_opt=paths=source_relative --go-grpc_out=. --go-grpc_opt=paths=source_relative internal\proto\classification.proto

	protoc --proto_path=internal/proto \
		--go_out=internal/protobuf --go_opt=module=go-router/internal/protobuf \
		--go-grpc_out=internal/protobuf --go-grpc_opt=module=go-router/internal/protobuf \
		internal/proto/classification.proto

run-server-dev:
	go run cmd/server/main.go

run-client-dev:
	go run cmd/client/main.go