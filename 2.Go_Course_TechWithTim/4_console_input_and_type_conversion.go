package main

import (
	// "log"
	"bufio"
	"fmt"
	"os"
	"strconv"
)

func main() {
	fmt.Print("Write your name: ")

	nameScanner := bufio.NewScanner(os.Stdin)
	nameScanner.Scan()

	// var input string = nameScanner.Text()
	nameInput := nameScanner.Text()
	fmt.Println("You name is:", nameInput)
	fmt.Println("\n")

	fmt.Print("Write your age: ")

	ageScanner := bufio.NewScanner(os.Stdin)
	ageScanner.Scan()

	ageInput, _ := strconv.ParseInt(ageScanner.Text(), 10, 64)

	fmt.Println("Your age is:", ageInput)
	fmt.Println("Your age is of type:", fmt.Sprintf("%T", ageInput))

}
