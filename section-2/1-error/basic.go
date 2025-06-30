package main

import (
	"errors"
	"fmt"
)

func validateAge(age int) error {
	if age < 0 {
		return errors.New(fmt.Sprintf("AGE CANNOT BE NEGATIVE: %d", age))
	}
	if age > 120 {
		return fmt.Errorf("AGE IS UNREALISTICALLY HIGH: %d", age)
	}
	return nil
}

func main() {
	ages := []int{-5, 30, 150}

	err := validateAge(ages[0])
	if err != nil {
		fmt.Println("Error:", err)

		//handling bisa dimasukan kedalam log file

		return
	}

	for _, age := range ages {
		if err := validateAge(age); err != nil {
			fmt.Println("Error:", err)
		} else {
			fmt.Printf("Age %d is valid.\n", age)
		}
	}
}
