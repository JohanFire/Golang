package main

import "fmt"

func main() {

	fmt.Printf("Hello %T %v", 10, 10)

	fmt.Println()

	var storedPrint string = fmt.Sprintf("Hello %T %v, this is a storedPrint", 10, 10)

	fmt.Println(storedPrint)
	fmt.Println(storedPrint)
	fmt.Println(storedPrint[0:5])

	fmt.Println()

	var areYouSmart bool = true
	fmt.Printf("Are you smart? %t \n", areYouSmart)

}
