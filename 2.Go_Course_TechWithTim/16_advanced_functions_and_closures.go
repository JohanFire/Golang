package main

import "fmt"

// say_hi receives a name and prints a greeting
func say_hi(name string) {
	fmt.Println("Hello " + name)
}

func pass_a_function_to_a_function(myFunc func(string) string, name string) {
	// here we can call the function passed as a parameter
	fmt.Printf("%s, this is a function passed to a function. \n", myFunc(name))
}

func returnFunction(name string) func() {
	return func() { fmt.Printf("Hi %s, this is a function that returns a function. \n", name) } // this is a Function Closure

	// this is a Function Closure
	return func() {
		fmt.Printf("Hi %s, this is a function that returns a function. \n", name)
	}
}

func main() {
	var name string = "Johan"

	say_hi(name)

	// ############################
	// calling a function without () means we are refering to the function, not calling it
	// something like this would be a function pointer
	fmt.Println(say_hi)

	// we can assign a function to a variable and it will work as well
	myFunc := say_hi
	myFunc(name)

	fmt.Println()

	// ############################
	// Anonymous functions
	// also make a function inside a function, assigned to a variable
	wuau := func() {
		fmt.Println("This is an Anonymous Function, a function inside a function, assigned to a variable")
	}
	wuau()

	// calling the function directly from where defined
	// here is not assigned to a variable, but can be as well
	func(num3, num4 int) {
		fmt.Printf("This is an Anonymous Function, a function inside a function, with params %d and %d \n", num3, num4)
	}(5, 10)

	fmt.Println()

	// ############################
	// Pass a function as a parameter

	anonFunction := func(yourName string) string {
		return "Hello " + yourName
	}

	pass_a_function_to_a_function(anonFunction, name)

	fmt.Println()

	// ############################
	// Function Closure
	// a function that returns another function

	returnFunction(name)()
	// the reason why needs a double (): returnFunction()()
	// is because the first one is the function that returns the function,
	// and the second one is the function returned by the first one

	// easier to understand with a variable
	closure := returnFunction("Tristán")
	closure()

}
