package main

import "fmt"

func main() {

	var myList []int = []int{16, 34, 64, 76, 88}

	for i := 0; i < len(myList); i++ {
		fmt.Println(myList[i])
	}

	fmt.Println()

	// like in python
	// for i, element in enumerate(myList):
	for i, element := range myList {
		fmt.Printf("index: %d, element: %d \n", i, element)
	}

	fmt.Println()

	// _ stands as "ignore" = not store. So u can't even use that variable
	for _, element := range myList {
		fmt.Printf("element: %d \n", element)
	}

	fmt.Println()

	for _, element := range myList {
		fmt.Println(element)

		for _, element2 := range myList {
			fmt.Println(element, element2)
		}
	}

	fmt.Println()
	fmt.Println()

	for i := 0; i < len(myList); i++ {
		var stack []int

		stack = append(stack, myList[i])

		for j := i + 1; j < len(myList); j++ {
			// fmt.Println(myList[j])

			stack = append(stack, myList[j])
		}

		fmt.Println(stack)
	}
}
