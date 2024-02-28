package main

import (
	"fmt"
	"math"
)

type Shape struct {
	ShapeType string
	Length    int
	Area      float32
}

const (
	RECTANGLE = "rectangle"
	CIRCLE    = "circle"
	TRIANGLE  = "triangle"
)

func rectangleArea(channel <-chan Shape) {
	rectangle := <-channel
	nama := rectangle.ShapeType
	sisi := float32(rectangle.Length)
	area := sisi * sisi
	fmt.Printf("%s luasnya adalah %.2f\n", nama, area)
}

func circleArea(channel <-chan Shape) {
	circle := <-channel
	nama := circle.ShapeType
	sisi := float32(circle.Length)
	area := math.Pi * sisi * sisi
	fmt.Printf("%s luasnya adalah %.2f\n", nama, area)
}

func triangleArea(channel <-chan Shape) {
	triangle := <-channel
	nama := triangle.ShapeType
	sisi := float32(triangle.Length)
	area := sisi * sisi / 2
	fmt.Printf("%s luasnya adalah %.2f\n", nama, area)
}

func main() {
	input := []Shape{
		{ShapeType: RECTANGLE, Length: 5},
		{ShapeType: CIRCLE, Length: 3},
		{ShapeType: TRIANGLE, Length: 5},
		{ShapeType: RECTANGLE, Length: 15},
		{ShapeType: CIRCLE, Length: 5},
	}
	channelRectangle := make(chan Shape)
	channelCircle := make(chan Shape)
	channelTriangle := make(chan Shape)
	go func() {
		channelRectangle <- input[0]
	}()
	rectangleArea(channelRectangle)
	go func() {
		channelCircle <- input[1]
	}()
	circleArea(channelCircle)
	go func() {
		channelTriangle <- input[2]
	}()
	triangleArea(channelTriangle)
	go func() {
		channelRectangle <- input[3]
	}()
	rectangleArea(channelRectangle)
	go func() {
		channelCircle <- input[4]
	}()
	circleArea(channelCircle)
}
