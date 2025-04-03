package main

import (
	"fmt"
	"strconv"
)

func main() {
	edad := 22

	// convert int to string
	edadStr := strconv.Itoa(edad)

	// convert string to int
	// this func returns 2 values
	// edadInt, err := strconv.Atoi(edadStr)
	edadInt, _ := strconv.Atoi(edadStr)
	// "_" guion bajo means I receive this value but I will drop it, I dont need it,

	fmt.Println("Mi edad es: " + edadStr)
	fmt.Println(edadInt + 10)
}
