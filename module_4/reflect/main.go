package main

import (
	"fmt"
	"reflect"
)

type student struct {
	Name  string
	Grade int
}

func (s *student) SetName(name string) {
	s.Name = name
}

func main() {
	number := 12
	reflectValue := reflect.ValueOf(number)
	fmt.Println("Type Variable : ", reflectValue.Type())
	fmt.Println("Nilai Variable : ", reflectValue.Interface())

	// Indentifying Method Information
	s1 := &student{Name: "Faqih yugos", Grade: 4}
	fmt.Println("Nama : ", s1.Name)

	reflectValueNew := reflect.ValueOf(s1)
	method := reflectValueNew.MethodByName("SetName")
	method.Call([]reflect.Value{
		reflect.ValueOf("Faqih"),
	})

	fmt.Println("Nama : ", s1.Name)

}
