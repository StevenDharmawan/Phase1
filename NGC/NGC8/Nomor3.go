package main

import "fmt"

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
	sisi := rectangle.Length
	area := sisi * sisi
	fmt.Printf("%s luasnya adalah %d\n", nama, area)
}

func circleArea(channel <-chan Shape) {
	circle := <-channel
	nama := circle.ShapeType
	sisi := circle.Length
	area := 22 / 7 * sisi * sisi
	fmt.Printf("%s luasnya adalah %d\n", nama, area)
}

func triangleArea(channel <-chan Shape) {
	triangle := <-channel
	nama := triangle.ShapeType
	sisi := triangle.Length
	area := sisi * sisi / 2
	fmt.Printf("%s luasnya adalah %d\n", nama, area)
}

func main() {
	input := []Shape{
		{RECTANGLE, 5, 0},
		{CIRCLE, 3, 0},
		{TRIANGLE, 5, 0},
		{RECTANGLE, 15, 0},
		{CIRCLE, 5, 0},
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
