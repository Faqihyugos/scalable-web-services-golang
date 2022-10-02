package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"text/template"
)

type Employee struct {
	ID       int
	Name     string `json:"name"`
	Age      int    `json:"age"`
	Division string `json:"division"`
}

var employess []Employee = []Employee{
	{
		ID:       1,
		Name:     "Ayat Maulana",
		Age:      20,
		Division: "Engineering",
	},
	{
		ID:       2,
		Name:     "Faqih",
		Age:      20,
		Division: "Finance",
	},
	{
		ID:       3,
		Name:     "Labib",
		Age:      20,
		Division: "Operation",
	},
}

func main() {
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprint(w, "hello world")
	})

	http.HandleFunc("/employees", func(w http.ResponseWriter, r *http.Request) {
		tpl, _ := template.ParseFiles("views/template.html")

		tpl.Execute(w, employess)
		return
	})

	http.HandleFunc("/employees/create", func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")

		if r.Method == "POST" {
			var newEmployee Employee
			json.NewDecoder(r.Body).Decode(&newEmployee)

			newEmployee.ID = len(employess) + 1
			employess = append(employess, newEmployee)

			json.NewEncoder(w).Encode(newEmployee)
			return
		}

		http.Error(w, "Invalid Method", http.StatusBadRequest)
	})

	http.ListenAndServe(":8080", nil)
}
