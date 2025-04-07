FROM golang:1.24.2-alpine

WORKDIR /app

# COPY . .
COPY . /app

# # Download and install dependencies
# RUN go mod tidy

# Compile the app
# RUN go build -o main .
# RUN go build -o main
RUN go build -o main main.go

EXPOSE 8080

CMD ["./main"]