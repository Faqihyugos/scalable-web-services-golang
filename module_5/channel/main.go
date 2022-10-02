package main

import (
	"fmt"
	"sync"
)

func main() {
	// c := make(chan string)

	// go Introduce("Airell", c)
	// go Introduce("Nanda", c)
	// go Introduce("Faqih", c)

	// msg1 := <-c
	// fmt.Println(msg1)

	// msg2 := <-c
	// fmt.Println(msg2)

	// msg3 := <-c
	// fmt.Println(msg3)

	// close(c)

	// Channel using Anonymous Function
	group = &sync.WaitGroup{}
	dataChannel := make(chan string)

	dataStudents := []string{"A", "B", "C", "D"}
	for _, value := range dataStudents {
		resultDataStudents := value
		go func() {
			group.Add(1)
			result := fmt.Sprintf("Hi, my name is %s", resultDataStudents)
			dataChannel <- result
			group.Done()
		}()
	}

	for i := 1; i <= len(dataStudents); i++ {
		Print(dataChannel)
	}

	group.Wait()
	close(dataChannel)
}

func Introduce(Student string, c chan string) {
	result := fmt.Sprintf("Hai, my name is %s", Student)

	c <- result
}

func Print(c <-chan string) {
	fmt.Println(<-c)
}

var group *sync.WaitGroup
