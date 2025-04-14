package main

import (
	"log"
	"os"
)

func main() {
	log.Println("Hello, World!")

	// above code is the same as below, defining it explicitly

	logger := log.Default()
	logger.SetOutput(os.Stderr)
	logger.Println("Hello World")
}
