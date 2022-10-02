package main

import (
	"fmt"
	"runtime"
	"time"
)

func FirstProcess(index int) {
	fmt.Println("First process func started")
	for i := 1; i <= index; i++ {
		fmt.Println("i : ", i)
	}
	fmt.Println("First process func ended")
}

func SecondProcess(index int) {
	fmt.Println("Second process func started")
	for i := 1; i <= index; i++ {
		fmt.Println("i : ", i)
	}
	fmt.Println("Second process func ended")
}

func main() {
	fmt.Println("Main Execution")
	go FirstProcess(5)
	SecondProcess(5)

	fmt.Println("No. of Goroutine : ", runtime.NumGoroutine())
	time.Sleep(2 * time.Second)

	fmt.Println("Main Execution Ended")
}
