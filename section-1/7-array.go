package main

import "fmt"

func main() {
	var arr1 [5]int
	var arr2 = [5]int{1, 2, 3, 4, 5}
	var arr3 = [...]string{"satu", "dua", "tiga", "empat", "lima"}

	fmt.Println("Array 1:", arr1)
	fmt.Println("Array 2:", arr2)
	fmt.Println("Array 3:", arr3)

	fmt.Println("Array 3-2:", arr3[2])

	var cars = [3]string{"Toyota", "Honda", "Suzuki"}
	fmt.Println(cars)
	fmt.Println("Jumlah mobil:", len(cars))
	fmt.Println("Mobil pertama:", cars[0])
}
