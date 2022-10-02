package main

import (
	"fmt"
	"sync"
)

var group *sync.WaitGroup

func PrintFruit(index int, fruit string) {
	fmt.Println("Index:", index, "- Fruit:", fruit)
	group.Done()
}

func main() {
	fruits := []string{"Banana", "Apple", "Orange"}
	group = &sync.WaitGroup{}

	for index, fruit := range fruits {
		group.Add(1)
		go PrintFruit(index, fruit)
	}

	group.Wait()
}
