package main

import (
	"fmt"
	"time"
)

func print(s string) {
	fmt.Printf("Image Processing completed: https://%s\n", s)
}

func processImage(imageURL []string) {
	for _, value := range imageURL {
		fmt.Printf("Processing image: https://%s\n", value)
	}

	for _, value := range imageURL {
		go print(value)
	}

	time.Sleep(3 * time.Second)
}

func main() {
	images := []string{
		"example.com/image1.jpg",
		"example.com/image2.jpg",
		"example.com/image3.jpg",
		"example.com/image4.jpg",
	}
	processImage(images)
}
