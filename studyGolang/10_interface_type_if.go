//interface变量存储的类型
//
//我们知道interface的变量里面可以存储任意类型的数值(该类型实现了interface)。那么我们怎么反向知道这个变量里面实际保存了的是哪个类型的对象呢？目前常用的有两种方法：
//
//Comma-ok断言
//
//Go语言里面有一个语法，可以直接判断是否是该类型的变量： value, ok = element.(T)，这里value就是变量的值，ok是一个bool类型，element是interface变量，T是断言的类型。
//
//如果element里面确实存储了T类型的数值，那么ok返回true，否则返回false。
//
//让我们通过一个例子来更加深入的理解。

package main

import (
	"fmt"
	"strconv"
)

type Element interface{}
type List []Element

type Person struct {
	name string
	age  int
}

//定义了String方法，实现了fmt.Stringer
func (p Person) String() string {
	return "(name: " + p.name + " - age: " + strconv.Itoa(p.age) + " years)"
}

func test_interface_type_if() {
	list := make(List, 3)
	list[0] = 1       // an int
	list[1] = "Hello" // a string
	list[2] = Person{"Dennis", 70}

	for index, element := range list {
		if value, ok := element.(int); ok {
			fmt.Printf("list[%d] is an int and its value is %d\n", index, value)
		} else if value, ok := element.(string); ok {
			fmt.Printf("list[%d] is a string and its value is %s\n", index, value)
		} else if value, ok := element.(Person); ok {
			fmt.Printf("list[%d] is a Person and its value is %s\n", index, value)
		} else {
			fmt.Printf("list[%d] is of a different type\n", index)
		}
	}
}

//是不是很简单啊，同时你是否注意到了多个if里面，还记得我前面介绍流程时讲过，if里面允许初始化变量。
//
//也许你注意到了，我们断言的类型越多，那么if else也就越多，所以才引出了下面要介绍的switch。

func main() {
	test_interface_type_if()
}
