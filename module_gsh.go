package module_gsh

import "fmt"

type Person struct {
	Name string
}

func (p Person) SayHello() {
	fmt.Println("Hello " + p.Name)
}

func (p Person) GetName() string {
	return p.Name
}
