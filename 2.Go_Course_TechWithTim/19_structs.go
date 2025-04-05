package main

import "fmt"

// Structs, are a way to group related data together
type Person struct {
	Name       string // public field, because it starts with an uppercase letter
	Age        int
	Country    string
	occupation string // private field, because it starts with a lowercase letter
}

func modify_struct_value(structPointer *Person) {
	structPointer.Age = 100
}

// Embedding structs
type Employee struct {
	person *Person
	Salary int
}

func main() {
	var johan Person = Person{"Johan", 23, "México", "Software Engineer"}

	fmt.Printf("%s is %d years old and lives in %s. \n", johan.Name, johan.Age, johan.Country)

	fmt.Println()

	person2 := Person{Name: "person2"}
	fmt.Println(person2) // {person2 0  }
	fmt.Printf("%s is %d years old and lives in %s. \n", person2.Name, person2.Age, person2.Country)

	fmt.Println()

	// ###################################3
	// Now, if we want to send an instance of a struct to a function
	// to modify one of its values for example,
	// we must sent the pointer to the struct,
	// otherwise it will not modify the original struct, will create a copy of it
	// like a local variable
	person3 := Person{Name: "person3", Age: 30, Country: "USA"}
	fmt.Println(person3) // {person3 30 USA }

	modify_struct_value(&person3)
	fmt.Println(person3) // {person3 100 USA }

	fmt.Println()

	// ###################################3
	// Embedding structs
	var employee1 Employee = Employee{&johan, 1000000}
	fmt.Println(employee1)                          // {0xc00019a0c0 100000}
	fmt.Println(employee1.person, employee1.Salary) // &{Johan 23 México } 1000000
	fmt.Println(employee1.person)                   // &{Johan 23 México }
	fmt.Println(employee1.person.Name)              // Johan

}
