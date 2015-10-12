package main

import (
	"fmt"
	"math"
)

//
// Interfaces
// An interface type is defined by a set of methods.
//
// 重点理解下面这句话:
// 一个接口类型值可以存放任意实现这些方法的值。
// A value of interface type can hold any value that implements those methods.
//
// Note:
//
// There is an error in the example code on line 22.
// Vertex (the value type) doesn't satisfy Abser
// because the Abs method is defined only on *Vertex (the pointer type).
//

type Abser interface {
	Abs3() float64
}

type MyFloat float64

func (f MyFloat) Abs3() float64 {
	if f < 0 {
		return float64(-f)
	}
	return float64(f)
}

type Vertex struct {
	X, Y float64
}

// receiver的类型，决定接口存放的类型实现
func (v *Vertex) Abs3() float64 {
	return math.Sqrt(v.X*v.X + v.Y*v.Y)
}

func main() {
	var a Abser
	f := MyFloat(-math.Sqrt2)
	v := Vertex{3, 4}
	// 或者如下
	// v := new(Vertex)
	// v.X = 2
	// v.Y = 8

	a = f  // a MyFloat implements Abser
	a = &v // a *Vertex implements Abser

	// In the following line, v is a Vertex (not *Vertex)
	// and does NOT implement Abser.
	//
	// cannot use v (type Vertex) as type Abser in assignment:
	// Vertex does not implement Abser (Abs3 method has pointer receiver)
	//
	// a = v  // 重点理解该语句上面报错的原因

	fmt.Println(a.Abs3())
}
