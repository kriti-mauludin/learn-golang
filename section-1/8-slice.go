package main

import "fmt"

func main() {
	sliceMakanan := []string{"Nasi Goreng", "Sate", "Rendang", "Soto", "Bakso"}
	sliceMinuman := []string{"Teh", "Kopi", "Jus", "Air Mineral", "Soda"}

	fmt.Println("Makanan:")
	for i, makanan := range sliceMakanan {
		println(i+1, makanan)
	}

	sliceMakananMinuman := append(sliceMakanan, sliceMinuman...)
	println("\nMakanan dan Minuman:")
	for i, item := range sliceMakananMinuman {
		println(i+1, item)
	}
}
