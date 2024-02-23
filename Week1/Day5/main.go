package main

import "fmt"

type Shape interface {
	area() float64
	perimeter() float64
}

type Circle struct {
}

func (c Circle) area() float64 {
	return 1.0
}

func (c Circle) perimeter() float64 {
	return 2.0
}

type Rectangle struct {
}

func (r Rectangle) area() float64 {
	return 3.0
}

func (r Rectangle) perimeter() float64 {
	return 4.0
}

func display(s Shape) {
	fmt.Println("Luas =", s.area())
	fmt.Println("Keliling =", s.perimeter())
}

func main() {
	newCircle := Circle{}
	display(newCircle)
	newRectangle := Rectangle{}
	display(newRectangle)

	var randomNumber interface{}
	randomNumber = 12
	fmt.Println(randomNumber)
}
