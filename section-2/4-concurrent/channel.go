package main

import (
	"fmt"
	"time"
)

func getOngkirJNE(ch chan string) {
	time.Sleep(2 * time.Second) // Simulate a delay
	ch <- "Ongkir JNE: Rp 10.000"
}

func getOngkirJNT(ch chan string) {
	time.Sleep(3 * time.Second) // Simulate a delay
	ch <- "Ongkir JNT: Rp 15.000"
}

func getOngkirSiCepat(ch chan string) {
	time.Sleep(1 * time.Second) // Simulate a delay
	ch <- "Ongkir SiCepat: Rp 12.000"
}

func main() {
	chJNE := make(chan string)
	chJNT := make(chan string)
	chSiCepat := make(chan string)

	go getOngkirJNE(chJNE)
	go getOngkirJNT(chJNT)
	go getOngkirSiCepat(chSiCepat)

	ongkirJNE := <-chJNE
	ongkirJNT := <-chJNT
	ongkirSiCepat := <-chSiCepat

	fmt.Println(ongkirJNE)
	fmt.Println(ongkirJNT)
	fmt.Println(ongkirSiCepat)
	fmt.Println("Semua ongkir telah diterima.")
}
