package main

import "fmt"

type Student struct {
	Name string
	Age  int
	// Grades []int
	Grades map[string]int8 // "subjetct": 0/100
}

// To create a struct method:
// func (receiver type) methodName(parameters) returnType {
func (s Student) get_name() string {
	return s.Name
}

func (s Student) get_age() int {
	return s.Age
}

// In this case, to set the change, receiverType should be a pointer to the struct type
func (s *Student) set_age(age int) {
	s.Age = age
}

func (s Student) add_grade(subject string, grade int8) {
	s.Grades[subject] = grade
}

func (s Student) get_average_grade() float32 {
	var sum float32

	for _, grade := range s.Grades {
		sum += float32(grade)
	}

	return sum / float32(len(s.Grades))

}

func main() {

	student1 := Student{
		Name: "Johan",
		Age:  23,
		// Grades: []int{100, 90, 80},
		Grades: map[string]int8{
			"math":    100,
			"english": 90,
			"science": 80,
		},
	}

	fmt.Println(student1) // {Johan 23 [100 90 80]}

	fmt.Println()

	fmt.Println(student1.get_name()) // Johan
	fmt.Println(student1.get_age())  // 23
	student1.set_age(24)
	fmt.Println(student1.get_age()) // 24

	fmt.Println()

	fmt.Println(student1.Grades) // map[english:90 math:100 science:80]
	student1.add_grade("history", 85)
	fmt.Println(student1.Grades) // map[english:90 history:85 math:100 science:80]

	fmt.Println()

	fmt.Println(student1)

	fmt.Println()

	fmt.Println("Average grade:", student1.get_average_grade())
}
