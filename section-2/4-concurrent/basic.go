package main

import (
	"fmt"
	"runtime"
	"sync"
	"time"
)

func TaskA(wg *sync.WaitGroup, index int) {
	defer wg.Done()
	time.Sleep(1 * time.Second)
	fmt.Println("Task A is running:", index)
}

func TaskB(wg *sync.WaitGroup, index int) {
	defer wg.Done()
	time.Sleep(2 * time.Second)
	fmt.Println("Task B is running:", index)
}

func TaskC(wg *sync.WaitGroup, index int) {
	defer wg.Done()
	fmt.Println("Task C is running:", index)
}

func main() {
	runtime.GOMAXPROCS(runtime.NumCPU()) // Set the number of OS threads to use
	fmt.Println("Number of CPU cores:", runtime.NumCPU())
	startTime := time.Now()

	var wg sync.WaitGroup

	for i := 1; i < 100; i++ {
		wg.Add(3)
		go TaskA(&wg, i) // Launching TaskA as a goroutine
		go TaskB(&wg, i) // Launching TaskB as a goroutine
		go TaskC(&wg, i) // Launching TaskC as a goroutine
	}

	wg.Wait() // Wait for all goroutines to finish
	fmt.Println("All tasks completed in:", time.Since(startTime))
}
