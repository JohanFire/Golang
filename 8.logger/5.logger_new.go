package main

import (
	"log"
	"os"
)

func main() {
	logger := log.New(
		os.Stderr,
		"Prefix: ",
		// log.Ldate|log.Ltime|log.Lshortfile|log.Lmsgprefix,
		log.Ldate|log.Ltime|log.Lshortfile,
	)

	logger.Println("Hello World")
}
