package main

import (
	"fmt" // allows show message on console = format strings

	"github.com/google/uuid"
)

func main() {
	println("Hello world")
	fmt.Println("Hello world fmt")
	fmt.Println(uuid.New().String())
}
