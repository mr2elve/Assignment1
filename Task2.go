package main

import (
	"fmt"
	"math"
)

type Shape interface {
	Area() float64
	Perimeter() float64
}

type Circle struct {
	Radius float64
}

func (c Circle) Area() float64 {
	return math.Pi * c.Radius * c.Radius
}

func (c Circle) Perimeter() float64 {
	return 2 * math.Pi * c.Radius
}

type Rectangle struct {
	Length float64
	Width  float64
}

func (r Rectangle) Area() float64 {
	return r.Length * r.Width
}

func (r Rectangle) Perimeter() float64 {
	return 2 * (r.Length + r.Width)
}

// PrintShapeDetails function
func PrintShapeDetails(s Shape) {
	fmt.Printf("\nShape: %T\n", s)
	fmt.Printf("Area: %.2f\n", s.Area())
	fmt.Printf("Perimeter: %.2f\n", s.Perimeter())
}

func main() {
	var choice int

	fmt.Println("Choose a shape:")
	fmt.Println("1. Circle")
	fmt.Println("2. Rectangle")
	fmt.Print("Enter your choice: ")
	fmt.Scan(&choice)

	switch choice {
	case 1:
		var radius float64
		fmt.Print("Enter the radius of the circle: ")
		fmt.Scan(&radius)
		circle := Circle{Radius: radius}
		PrintShapeDetails(circle)
	case 2:
		var length, width float64
		fmt.Print("Enter the length of the rectangle: ")
		fmt.Scan(&length)
		fmt.Print("Enter the width of the rectangle: ")
		fmt.Scan(&width)
		rectangle := Rectangle{Length: length, Width: width}
		PrintShapeDetails(rectangle)
	default:
		fmt.Println("Invalid choice!")
	}
}
