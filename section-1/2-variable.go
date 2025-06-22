package main

import "fmt"

func main() {
	var name string
	name = "KM"

	fmt.Println("Hello", name)

	umur := 20
	fmt.Println("Umur", umur)

	var noRumah uint
	noRumah = 123
	fmt.Println("No Rumah", noRumah)

	fmt.Printf("Nama: %s, Umur: %d, No Rumah: %d\n", name, umur, noRumah)

	var alamat any
	alamat = "Jalan Raya"
	fmt.Println("Alamat:", alamat)
}
