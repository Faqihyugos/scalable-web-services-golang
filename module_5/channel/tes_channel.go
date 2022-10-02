package main

import (
	"fmt"
	"sync"
)

var WaitGroup *sync.WaitGroup

// Test Slice of Map of Interface{} Channel with Anonymous Function
func main() {
	c := make(chan map[string]interface{}, 3)
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

	go func(c chan map[string]interface{}) {
		for _, value := range data {
			c <- value
		}
		close(c)
	}(c)

	for v := range c {
		if v["Identity"].(map[string]interface{})["isMahasiswa"] == true {
			fmt.Println("Name:", v["Identity"].(map[string]interface{})["Name"], "- Age:", v["Identity"].(map[string]interface{})["Age"], "is student")
		} else {
			fmt.Println("Name:", v["Identity"].(map[string]interface{})["Name"], "- Age:", v["Identity"].(map[string]interface{})["Age"], "is not student")
		}
	}
}
