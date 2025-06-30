package main

import "fmt"

func main() {
	withoutPointer := 3000

	var withPointer *int
	withPointer = &withoutPointer

	fmt.Println("Without Pointer:", withoutPointer)
	fmt.Println("With Pointer:", *withPointer)
	fmt.Println("Address of withoutPointer:", &withoutPointer)
	fmt.Println("Address of withPointer:", withPointer)

	length := 5
	width := 10
	area, err := calculateAreaOfRectangle(length, width)
	if err != nil {
		fmt.Println("Error:", err)
	} else {
		fmt.Printf("Area of rectangle with length %d and width %d is: %d\n", length, width, *area)
	}
}

func calculateAreaOfRectangle(length, width int) (*int, error) {
	//calculate area of rectangle
	if length == 0 || width == 0 {
		return nil, fmt.Errorf("LENGTH AND WIDTH MUST BE GREATER THAN ZERO: LENGTH=%d, WIDTH=%d", length, width)
	}
	area := length * width
	return &area, nil
}
