package main

import (
	"fmt"
)

type Animal interface {
	Speak() string
}

type Dog struct {
}

func (d Dog) Speak() string {
	return "Woof!"
}

type Cat struct {
}

func (c *Cat) Speak() string {
	return "Meow!"
}

type Llama struct {
}

func (l Llama) Speak() string {
	return "?????"
}

type JavaProgrammer struct {
}

func (j JavaProgrammer) Speak() string {
	return "Design patterns!"
}

func main() {
	animals := []Animal{Dog{}, Cat{}, Llama{}, JavaProgrammer{}}
	// 40: cannot use Cat literal (type Cat) as type Animal
	// in array element:Cat does not implement Animal
	// (Speak method has pointer receiver)
	for _, animal := range animals {
		fmt.Println(animal.Speak())
	}
}
