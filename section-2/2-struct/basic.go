package main

import "fmt"

type Student struct {
	Name  string
	Age   *uint8 // Using pointer to allow nil value
	Grade string
	Major string
}

func main() {
	student1 := Student{
		Name:  "K",
		Grade: "A",
		Major: "Computer Science",
	}

	student2 := Student{
		Name:  "M",
		Grade: "B",
		Major: "Mathematics",
	}

	age := new(uint8)  // Initialize pointer
	*age = 20 // Assign value to the pointer
	student1.Age = age

	fmt.Println("Student 1:", student1)
	fmt.Println("Student 2:", student2)

	//update student1 name
	student1.Name = "KM"
	fmt.Println("Student 1 Name:", student1.Name)
}
