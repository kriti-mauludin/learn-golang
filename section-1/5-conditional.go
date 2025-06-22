package main

import "fmt"

func main() {
	agent := "KM"
	age := 20

	if agent == "KM" {
		fmt.Println("Hello KM")
	} else {
		fmt.Println("Hello ", agent)
	}

	switch age {
	case 20:
		fmt.Println("Umurmu 20 tahun")
	case 30:
		fmt.Println("Umurmu 30 tahun")
	default:
		fmt.Println("Umurmu bukan 20 atau 30 tahun")
	}
}
