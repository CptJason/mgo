package main

import "fmt"

//
// 空interface
// 空interface(interface{})不包含任何的method，
// 正因为如此，
// 所有的类型都实现了空interface。
// 空interface对于描述起不到任何的作用(因为它不包含任何的method），
// --> 但是空interface在我们需要存储<任意类型的数值>的时候相当有用 <--
// 注意：是任意类型的数值，而不是任意类型，数值！数值！数值！
// --> 一个函数把interface{}作为参数，那么他可以接受任意类型的值作为参数 <--
// --> 如果一个函数返回interface{},那么也就可以返回任意类型的值。<--
// 是不是很有用啊！
// 因为它可以存储任意类型的数值。它有点类似于C语言的void*类型。

// 定义a为空接口
func test_zero_interface() {
	var a interface{}
	var i int = 5
	s := "Hello world"
	numbers := make(map[string]int)
	numbers["one"] = 1  //赋值
	numbers["ten"] = 10 //赋值
	numbers["three"] = 3
	// a可以存储任意类型的数值
	a = i
	fmt.Printf("%d \n", a)
	a = s
	fmt.Printf("%s \n", a)

	// 空接口无法map索引
	// invalid operation: a["one"]
	// (type interface {} does not support indexing)
	a = numbers
	//	fmt.Printf("%d ", a["one"])

	// 空接口不支持slice索引
	// invalid operation: a[0]
	// (type interface {} does not support indexing)
	slice := []byte{'a', 'b', 'c', 'd', 'e', 'f', 'g', 'h', 'i'} // 声明一个slice, 并初始化数据
	a = slice
	//	fmt.Printf("%s ", a[0])
}

func main() {
	test_zero_interface()
}
