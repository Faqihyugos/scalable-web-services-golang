package main

import (
	"fmt"
	"math"
)

type Shape interface {
	area() float64
	perimeter() float64
}

type rectangle struct {
	Width, Height float64
}

type circle struct {
	radius float64
}

// implementasi method circle
func (c *circle) area() float64 {
	return math.Pi * math.Pow(c.radius, 2)
}

func (c *circle) perimeter() float64 {
	return 2 * math.Pi * c.radius
}

// method di luar interface Shape
func (c *circle) volume() float64 {
	return 4 / 3 * math.Pi * math.Pow(c.radius, 3)
}

// implementasi method rectangle
func (r *rectangle) area() float64 {
	return r.Height * r.Width
}

func (r *rectangle) perimeter() float64 {
	return 2 * (r.Height + r.Width)
}

func Print(t string, s Shape) {
	fmt.Printf("%s Area %v\n", t, s.area())
	fmt.Printf("%s Perimeter %v\n", t, s.perimeter())
}

func main() {
	var c1 Shape = &circle{radius: 5}
	var r1 Shape = &rectangle{Width: 3, Height: 2}

	fmt.Printf("Type of c1: %T\n", c1)
	fmt.Printf("Type of r1: %T\n", r1)
	fmt.Println()

	// // call method
	// // circle
	// fmt.Println("Circle Area :", c1.area())
	// fmt.Println("Circle Perimeter :", c1.perimeter())
	// // rectangle
	// fmt.Println("Rectangle Area :", r1.area())
	// fmt.Println("Rectangle Perimeter :", r1.perimeter())

	// Call interface method in func print
	Print("Circle", c1)
	Print("Rectangle", r1)

	// Type Assertion when volume method is not define at Shape interface
	value, ok := c1.(*circle)
	if ok == true {
		fmt.Printf("Circle value: %+v\n", value)
		fmt.Printf("Circle volume: %f", value.volume())
	}

}
