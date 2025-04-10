# Golang
Learning Go

### Check if go is already install
or also if you wanna check Go's version, open a cmd and type: `go version` <br/>

### GO env
GO has 2 main env variables, Go path & Go routh, you can see them with:
```bash
    go env
```
```bash
...
set GOPATH=C:\Users\johan\go
...
set GOROOT=C:\Program Files\Go
...
```
GOROOT is where Go is installed.
GOPATH is where you will install all your packages

### Start a Go project
similar to `npm init`... <br/>
```bash
go mod init "UNIQUE_URL"
```
_Example usage:
        'go mod init example.com/m' to initialize a v0 or v1 module
        'go mod init example.com/m/v2' to initialize a v2 module_

but in my case, let's put github's url like:
```bash
go mod init github.com/JohanFire/Golang
```

in this [go.mod](./go.mod) we will listing all our packages & dependencies for the project

### Run Main code
```bash
go run .

# also if you want to run a code with a different name:
go run file.go
```

### Install packages
can check all packages in [Go Packages](https://pkg.go.dev/)
<br/>

#### UUID package
_The uuid package generates and inspects UUIDs_
for this case I will install the UUID package:
```bash
go get github.com/google/uuid
```

### Create .exe of the project
of course the .exe executes the Main code as principal
```bash
go build .

# also if you want to build a code with a different name:
go build file.go
```

so you can execute now the .exe, it will run the same, but now compiled.

```bash
./Golang.exe
```


## Go work example
```bash
go work init 
# 
go work init <directory>
```

```go
go 1.24.2

use ./6.gRPC/1.gRPC
// use ./6.gRPC/3.gRPC_coffee_shop
```