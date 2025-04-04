/*
Golang Maps are a built-in data type that associates keys with values.
Maps are unordered collections of key-value pairs, where each key is unique and maps to a specific value.

So, basically a Map is equivalent to a Dictionary in Python.

Maps in Go does not keep the order of the keys.
Maps just know each key and its corresponding value.
*/
package main

import "fmt"

func main() {

	// 1st way to define a map
	// Define map type
	// 		  map[key]value
	var myMap map[string]int
	fmt.Println(myMap) // <nil>

	var myMap2 map[string]int = map[string]int{
		"apple":  5,
		"banana": 10,
		"orange": 15,
		"grape":  20,
		"kiwi":   25,
	}
	fmt.Println("myMap2:", myMap2) // response keys, won't be in order

	fmt.Println()

	// 2nd way to define a map
	myMap3 := map[string]string{
		"apple":  "red",
		"banana": "yellow",
		"orange": "orange",
		"grape":  "purple",
		"kiwi":   "green",
	}
	fmt.Println("myMap3:", myMap3)

	fmt.Println()

	// 3rd way to define a map, with make()
	myMap4 := make(map[string]int)
	// add values to the map
	myMap4["apple"] = 5
	myMap4["banana"] = 10
	myMap4["orange"] = 15

	fmt.Println("myMap4:", myMap4)
	fmt.Println("myMap4['banana']:", myMap4["banana"])

	myMap4["banana"] = 11
	fmt.Println("myMap4['banana']:", myMap4["banana"])

	// delete a key-value pair
	delete(myMap4, "banana")
	fmt.Println("myMap4['banana'] deleted")
	fmt.Println("myMap4:", myMap4)

	fmt.Println()

	// check if a key exists
	// if the key exists:
	// 	- store the value in 'value'
	// 	- ok will be true
	// if the key doesn't exist:
	// 	- store the zero value of the value type in 'value'
	// 	- ok will be false
	fmt.Println("Check if key exists")

	value, ok := myMap4["apple"]
	fmt.Println(value, ok)

	value2, ok := myMap4["banana"]
	fmt.Println(value2, ok)

	if value, ok := myMap4["kiwi"]; ok {
		fmt.Println("Key 'kiwi' exists in myMap4:", value)
	} else {
		fmt.Println("Key 'kiwi' does not exist in myMap4")
	}

	fmt.Println()

	// iterate over a map
	fmt.Println("Iteration over myMap3")
	for key, value := range myMap3 {
		fmt.Printf("Key: %s, Value: %s \n", key, value)
	}

	fmt.Println()

	// get length of a map
	fmt.Println("Length of myMap3:", len(myMap3))

}
