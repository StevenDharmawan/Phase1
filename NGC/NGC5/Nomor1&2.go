package main

import (
	"fmt"
)

type Person struct {
	Name, Job string
	Age       int
}

func (person Person) GetInfo() {
	fmt.Println("Name:", person.Name)
	fmt.Println("Age:", person.Age)
	fmt.Println("Job:", person.Job)
}

func (person *Person) AddYear() {
	person.Age++
	if person.Age == 50 {
		person.Job = "Retired"
	}
}

func main() {
	//Nomor 1
	bambang := Person{"Bambang", "Gambler", 20}
	bambang.GetInfo()
	budi := Person{"Budi", "Gambler", 45}
	budi.AddYear()
	budi.AddYear()
	budi.AddYear()
	budi.AddYear()
	budi.AddYear()
	budi.GetInfo()

	//Nomor2
	person := []Person{
		{"Bambang", "Job 1", 20},
		{"Budi", "Job 2", 40},
		{"Herry", "Job 3", 30},
	}

	for _, value := range person {
		value.GetInfo()
	}
}
