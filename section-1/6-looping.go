package main

import "fmt"

func main() {
	for i := 1; i <= 5; i++ {
		fmt.Println("Perulangan ke-", i)
	}

	fmt.Println("--------------------")

	buah := []string{"apel", "jeruk", "mangga", "pisang"}
	for i := 0; i < len(buah); i++ {
		fmt.Println("Buah ke-", i+1, "adalah", buah[i])
	}
}
