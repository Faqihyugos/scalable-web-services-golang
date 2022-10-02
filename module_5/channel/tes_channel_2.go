package main

import (
	"fmt"
	"sync"
)

var wg *sync.WaitGroup

func SendData(data map[string]interface{}, c chan<- map[string]interface{}) {
	c <- data
}

func ReceiveData(data map[string]interface{}, c <-chan map[string]interface{}) {
	data = <-c
	if data["Identity"].(map[string]interface{})["isMahasiswa"] == true {
		fmt.Println("Name:", data["Identity"].(map[string]interface{})["Name"], "- Age:", data["Identity"].(map[string]interface{})["Age"], "is student")
	} else {
		fmt.Println("Name:", data["Identity"].(map[string]interface{})["Name"], "- Age:", data["Identity"].(map[string]interface{})["Age"], "is not student")
	}
	wg.Done()
}

// Test Slice of Map of Interface{} Channel with Function
func main() {
	data := []map[string]interface{}{
		{
			"Identity": map[string]interface{}{
				"Name":        "Dias",
				"Age":         23,
				"isMahasiswa": true,
			},
		},
		{
			"Identity": map[string]interface{}{
				"Name":        "Nayla",
				"Age":         11,
				"isMahasiswa": false,
			},
		},
	}
	wg = &sync.WaitGroup{}
	c := make(chan map[string]interface{}, len(data))

	for _, value := range data {
		wg.Add(1)
		go SendData(value, c)
		go ReceiveData(value, c)
	}

	wg.Wait()
	close(c)
}
