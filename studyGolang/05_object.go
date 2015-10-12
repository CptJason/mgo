package main

import (
	"fmt"
	"math"
)

//
// 2.5 object
// https://github.com/astaxie/build-web-application-with-golang/blob/master/zh/02.5.md
//

// 是否想过函数当作struct的字段一样来处理呢？

//
// test area non receiver
//
type Rectangle struct {
	width, height float64
}

func area(r Rectangle) float64 {
	return r.width * r.height
}

func test_area_non_receiver() {
	r1 := Rectangle{12, 2}
	r2 := Rectangle{9, 4}
	fmt.Println("Area of r1 is: ", area(r1))
	fmt.Println("Area of r2 is: ", area(r2))
}

// 实现当然没有问题咯，
// 但是当需要增加圆形、正方形、五边形甚至其它多边形的时候，
// 你想计算他们的面积的时候怎么办啊？那就只能增加新的函数咯，
// 但是函数名你就必须要跟着换了，
// 变成area_rectangle, area_circle, area_triangle...
// 这样的实现并不优雅，并且从概念上来说"面积"是"形状"的一个属性，
// 它是属于这个特定的形状的，就像长方形的长和宽一样。

// 基于上面的原因所以就有了method的概念，
// method是附属在一个给定的类型上的，
// 他的语法和函数的声明语法几乎一样，
// 只是在func后面增加了一个receiver(也就是method所依从的主体)。
// 用上面提到的形状的例子来说，method area() 是依赖于某个形状
// (比如说Rectangle)来发生作用的。
// Rectangle.area()的发出者是Rectangle， area()是属于Rectangle的方法，
// 而非一个外围函数。
//
// 更具体地说，Rectangle存在字段length 和 width,
// 同时存在方法area(), 这些字段和方法都属于Rectangle。
//
// 用Rob Pike的话来说就是：
//
// "A method is a function with an implicit first argument, called a receiver."
// method的语法如下：
//
// func (r ReceiverType) funcName(parameters) (results)
// 下面我们用最开始的例子用method来实现：
// 在使用method的时候重要注意几点
//
// 虽然method的名字一模一样，但是如果接收者不一样，那么method就不一样
// method里面可以访问接收者的字段
// 调用method通过.访问，就像struct里面访问字段一样

//
// test area with receiver
//
type Circle struct {
	radius float64
}

func (r Rectangle) area() float64 {
	return r.width * r.height
}

func (c Circle) area() float64 {
	return c.radius * c.radius * math.Pi
}

func test_area_with_receiver() {
	r1 := Rectangle{12, 2}
	r2 := Rectangle{9, 4}
	c1 := Circle{10}
	c2 := Circle{25}

	fmt.Println("Area of r1 is: ", r1.area())
	fmt.Println("Area of r2 is: ", r2.area())
	fmt.Println("Area of c1 is: ", c1.area())
	fmt.Println("Area of c2 is: ", c2.area())
}

// 自定义类型
// 可以在自己的代码里面定义有意义的类型了
// 实际上只是一个定义了一个别名

//
// test custom type
//
func test_custom_type() {
	type ages int

	type money float32

	type months map[string]int

	m := months{
		"January":   31,
		"February":  28,
		"March":     31,
		"April":     30,
		"May":       31,
		"June":      30,
		"July":      31,
		"August":    31,
		"September": 30,
		"October":   31,
		"Novenber":  30,
		"December":  31,
	}
	fmt.Printf("月份天数是 %d", m["January"])
}

//
// test area with pointer receiver
//
const (
	WHITE = iota
	BLACK
	BLUE
	RED
	YELLOW
)

type Color byte

type Box struct {
	width, height, depth float64
	color                Color
}

type BoxList []Box //a slice of boxes

func (b Box) Volume() float64 {
	return b.width * b.height * b.depth
}

func (b *Box) SetColor(c Color) {
	b.color = c
}

func (bl BoxList) BiggestColor() Color {
	v := 0.00
	k := Color(WHITE)
	for _, b := range bl {
		if bv := b.Volume(); bv > v {
			v = bv
			k = b.color
		}
	}
	return k
}

func (bl BoxList) PaintItBlack() {
	for i, _ := range bl {
		bl[i].SetColor(BLACK)
	}
}

func (c Color) String() string {
	strings := []string{"WHITE", "BLACK", "BLUE", "RED", "YELLOW"}
	return strings[c]
}

func test_area_with_pointer_receiver() {
	boxes := BoxList{
		Box{4, 4, 4, RED},
		Box{10, 10, 1, YELLOW},
		Box{1, 1, 20, BLACK},
		Box{10, 10, 1, BLUE},
		Box{10, 30, 1, WHITE},
		Box{20, 20, 20, YELLOW},
	}

	fmt.Printf("We have %d boxes in our set\n", len(boxes))
	fmt.Println("The volume of the first one is", boxes[0].Volume(), "cm³")
	fmt.Println("The color of the last one is", boxes[len(boxes)-1].color.String())
	fmt.Println("The biggest one is", boxes.BiggestColor().String())

	fmt.Println("Let's paint them all black")
	boxes.PaintItBlack()
	fmt.Println("The color of the second one is", boxes[1].color.String())

	fmt.Println("Obviously, now, the biggest one is", boxes.BiggestColor().String())
}

//
// override method
//
type Human struct {
	name  string
	age   int
	phone string
}

type Student struct {
	Human  //匿名字段
	school string
}

type Employee struct {
	Human   //匿名字段
	company string
}

//在human上面定义了一个method
func (h *Human) SayHi1() {
	fmt.Printf("Hi, I am %s you can call me on %s\n", h.name, h.phone)
}

func test_method_inherit() {
	mark := Student{Human{"Mark", 25, "222-222-YYYY"}, "MIT"}
	sam := Employee{Human{"Sam", 45, "111-888-XXXX"}, "Golang Inc"}

	mark.SayHi1()
	sam.SayHi1()
}

//
// override method
//

//Human定义method
func (h *Human) SayHi2() {
	fmt.Printf("Hi, I am %s -> you can call me on %s\n", h.name, h.phone)
}

//Employee的method重写Human的method
func (e *Employee) SayHi2() {
	fmt.Printf("Hi, I am %s, -> I work at %s. -> Call me on %s\n", e.name,
		e.company, e.phone) //Yes you can split into 2 lines here.
}

func test_method_override() {
	mark := Student{Human{"Mark", 25, "222-222-YYYY"}, "MIT"}
	sam := Employee{Human{"Sam", 45, "111-888-XXXX"}, "Golang Inc"}

	mark.SayHi2()
	sam.SayHi2()
}

func main() {
	//	test_area_non_receiver()
	//	test_area_with_receiver()
	//	test_custom_type()
	//	test_area_with_pointer_receiver()
	//	test_method_inherit()
	test_method_override()

}

//
// 我们可以设计出基本的面向对象的程序了，
// 但是Go里面的面向对象是如此的简单，
// 没有任何的私有、公有关键字，
// 通过大小写来实现 (大写开头的为公有，
// 小写开头的为私有)
// 方法也同样适用这个原则。
//
