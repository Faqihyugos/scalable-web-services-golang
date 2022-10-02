package main

import "fmt"

func main() {
	var V interface{}
	V = 20

	// Use type assertion to reassign value
	if value, ok := V.(int); ok == true {
		V = value * 9
		fmt.Println(V)
	}

	// Map & Slice with Interface Kosong (Tanya live sesi)
	// sliceOfInterface := []interface{}{
	// 	1, "Airell", true,
	// 	2, "Ananda", true,
	// }

	mapOfInterface := []map[string]interface{}{
		{
			"Name":   "Dias",
			"Status": true,
			"Age":    23,
		},
		{
			"Name":   "Nayla",
			"Status": false,
			"Age":    11,
		},
		{
			"Name":   "Cita",
			"Status": true,
			"Age":    22,
		},
	}
	// Loop mapOfInterface
	for _, value := range mapOfInterface {
		if value["Status"] == true {
			fmt.Println("Name:", value["Name"], "- Age:", value["Age"], "is student")
		} else {
			fmt.Println("Name:", value["Name"], "- Age:", value["Age"], "is not student")
		}
	}
}
