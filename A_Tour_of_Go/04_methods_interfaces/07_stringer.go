package main

import "fmt"

//
// 无处不在的接口 - Stringers （fmt包中）
// One of the most ubiquitous interfaces is Stringer defined by the fmt package.
//
// type Stringer interface {
//     String() string
// }
//
// A Stringer is a type that can describe itself as a string.
// The fmt package (and many others) look for this interface to print values.
//

type Person struct {
	Name string
	Age  int
}

func (p Person) String() string {
	return fmt.Sprintf("%v (%v years)", p.Name, p.Age)
}

func main() {
	a := Person{"Arthur Dent", 42}
	z := Person{"Zaphod Beeblebrox", 9001}
	fmt.Println(a, z)
	//
	// 默认方法打印
	// {Arthur Dent 42} {Zaphod Beeblebrox 9001}
	//
	// 重新实现方法后
	// Arthur Dent (42 years) Zaphod Beeblebrox (9001 years)
	//
}
