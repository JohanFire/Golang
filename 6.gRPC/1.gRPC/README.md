Protocol Buffers (commonly referred to as Protobuf) is a language-neutral, platform-neutral mechanism for serializing structured data. It is widely used for defining and exchanging data between services in a compact and efficient format.

Before proceeding with a project that uses Protocol Buffers, you need to have Protocol Buffer Compiler (protoc) version 3 or higher installed on your machine. This is because:

Code Generation: Protobuf definitions are written in .proto files. The protoc compiler is required to generate code (e.g., Go, Python, Java, etc.) from these .proto files. Without protoc, you cannot generate the necessary code for your project.

Compatibility: Version 3 of protoc introduced several improvements over earlier versions, including support for optional fields, better syntax, and more efficient serialization. Many modern projects depend on features available only in version 3 or higher.

Integration with Go: If your project is in Go, you will likely use plugins like protoc-gen-go to generate Go-specific code from .proto files. These plugins work in conjunction with protoc.

To install protoc:

On Linux: Use your package manager (e.g., apt, yum) or download the binary from the official Protobuf GitHub releases.
On macOS: Use Homebrew (brew install protobuf).
On Windows: Download the binary from the official Protobuf GitHub releases.
Once installed, verify the installation by running:

This should output the installed version, which should be libprotoc 3.x.x or higher.

## Prerequisites

Before proceeding, ensure that Protocol Buffers (protoc) version 3 is installed on your machine. You can download it from the [official Protocol Buffers GitHub releases page](https://github.com/protocolbuffers/protobuf/releases). Follow the installation instructions for your operating system.

### Installation
To install Protocol Buffers (protoc) on your machine:
[https://youtu.be/ES_GI-lmhEU](https://youtu.be/ES_GI-lmhEU)


Verify the installation by running the following command in your terminal:
```bash
protoc --version
```
This should output the installed version of protoc.

### Install gRPC package for Go
To install the gRPC package for Go, you can use the following command:
```bash
# use this, but now deprecated
go get -u github.com/golang/protobuf/protoc-gen-go

# instead
google.golang.org/protobuf
```


## Generating Go Code from Protobuf Definitions
To generate Go code from your .proto files, you can use the protoc command with the protoc-gen-go plugin. Here's a step-by-step guide:
1. **Create a .proto file**: Define your message types and services in a .proto file. As shown in [`classification.proto`](./classification.proto).

2. **Run the protoc command**: Use the following command to generate Go code from your .proto file:
```bash
# protoc --go_out=. --go-grpc_out=. classification.proto
protoc --go_out=plugins=grpc:service.proto

#or if you have a specific path for the proto file

# protoc --go_out=plugins=grpc:directoryPath proto/classification.proto
protoc --go_out=plugins=grpc:proto proto/classification.proto

# protoc --go_out=. --go-grpc_out=. --proto_path=./ classification.proto

# Final command
protoc --go_out=. --go-grpc_out=. classification.proto

```


