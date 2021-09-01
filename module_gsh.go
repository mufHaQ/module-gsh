package module_gsh

import "fmt"

type Person struct {
	Name string
	Age  int
}

func (p Person) SayHello() {
	fmt.Println("Hello " + p.Name)
}

func (p Person) GetName() string {
	return p.Name
}

func (p Person) GetAge() int {
	return p.Age
}
