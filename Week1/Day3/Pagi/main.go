package main

import "fmt"

type Employee struct {
	Name, Address string
	Age           int
	Salary        Salary
}

type Salary struct {
	nominal int
}

func (employee Employee) Introduce(msg string) string {
	return fmt.Sprintf("%s My name is %s and I'm %d years old", msg, employee.Name, employee.Age)
}

func (employee Employee) ChangeName1() {
	employee.Name = "Reinard"
}

func (employee *Employee) ChangeName2() {
	employee.Name = "Reinard"
}

func main() {
	//Struct 1
	// var andrew Employee
	// andrew.Name = "Andrew"
	// andrew.Address = "Jakarta"
	// andrew.Age = 20
	// fmt.Println(andrew)
	// fmt.Println(andrew.Name)
	// fmt.Println(andrew.Address)
	// fmt.Println(andrew.Age)

	//Struct 2
	// charly := Employee{}
	// charly.Name = "Charly"
	// charly.Address = "Bandung"
	// charly.Age = 25
	// fmt.Println(charly)
	// fmt.Println(charly.Name)
	// fmt.Println(charly.Address)
	// fmt.Println(charly.Age)

	//Struct 3
	// budi := Employee{"Budi", "Wakanda", 18}
	// fmt.Println(budi)
	// fmt.Println(budi.Name)
	// fmt.Println(budi.Address)
	// fmt.Println(budi.Age)

	employee1 := Employee{"Bambang", "Bintaro", 23, Salary{10000000}}

	var employee2 *Employee = &employee1

	fmt.Printf("Value employee1 %+v\n", employee1)
	fmt.Printf("Value employee2 %+v\n", employee2)

	fmt.Println("========================")
	employee2.Name = "Jono"
	fmt.Printf("Value employee1 %+v\n", employee1)
	fmt.Printf("Value employee2 %+v\n", employee2)
	fmt.Println(employee1.Introduce("Hello all,"))
	employee1.ChangeName1()
	fmt.Println("Change name with ChangeName1 method", employee1.Name)
	employee1.ChangeName2()
	fmt.Println("Change name with ChangeName2 method", employee1.Name)

	//Pointer
	// var num int = 4
	// var num2 *int = &num

	// fmt.Println("Alamat dari num = ", &num)
	// fmt.Println("Value dari num = ", num)
	// fmt.Println("Alamat dari num2 = ", num2)
	// fmt.Println("Value dari num2 = ", *num2)

	// age := 17
	// fmt.Println(age, "<--- Before")
	// addNumberNotPointer(age)
	// fmt.Println(age, "<--- After Not Pointer")
	// addNumberPointer(&age)
	// fmt.Println(age, "<--- After Pointer")

}

// func addNumberPointer(num *int) {
// 	*num += 10
// }

// func addNumberNotPointer(num int) {
// 	num += 10
// }
