# FROM golang:1.20
# FROM golang:1.24
FROM golang:1.24-alpine
# alpine are much slimmer images in general

WORKDIR /app

# COPY . .
COPY . /app

# Compile single script, no modules
RUN go build -o main main.go

CMD ["./main"]