package main

import (
	"fmt"
)

type Animal interface {
	Speak2() string
}

type Dog struct {
}

func (d Dog) Speak2() string {
	return "Woof!"
}

type Cat struct {
}

func (c *Cat) Speak2() string {
	return "Meow!"
}

type Llama struct {
}

func (l Llama) Speak2() string {
	return "?????"
}

type JavaProgrammer struct {
}

func (j JavaProgrammer) Speak2() string {
	return "Design patterns!"
}

func main() {
	animals := []Animal{Dog{}, new(Cat), Llama{}, JavaProgrammer{}}
	// 40: cannot use Cat literal (type Cat) as type Animal
	// in array element:Cat does not implement Animal
	// (Speak2 method has pointer receiver)
	for _, animal := range animals {
		fmt.Println(animal.Speak2())
	}
}
