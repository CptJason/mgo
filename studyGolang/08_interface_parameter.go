package main

import (
	"fmt"
	"strconv"
)

//interface函数参数
//
// interface的变量可以持有任意实现该interface类型的对象，
// 这给我们编写函数(包括method)提供了一些额外的思考，
// 我们是不是可以通过定义interface参数，让函数接受各种类型的参数。

// 举个例子：fmt.Println是我们常用的一个函数，
// 但是你是否注意到它可以接受任意类型的数据。打开fmt的源码文件，你会看到这样一个定义:

type Stringer interface {
	String() string
}

//也就是说，任何实现了String方法的类型都能作为参数被fmt.Println调用,让我们来试一试

type Human struct {
	name  string
	age   int
	phone string
}

// 通过这个方法 Human 实现了 fmt.Stringer
func (h Human) String() string {
	return "❰" + h.name + " - " + strconv.Itoa(h.age) + " years -  ✆ " + h.phone + "❱"
}

func test_interface_Stringer() {
	Bob := Human{"Bob", 39, "000-7777-XXX"}
	fmt.Println("This Human is : ", Bob)
	// 现在我们再回顾一下前面的Box示例，
	// 你会发现Color结构也定义了一个method：String。
	// 其实这也是实现了fmt.Stringer这个interface，
	// 即如果需要某个类型能被fmt包以特殊的格式输出，
	// 你就必须实现Stringer这个接口。
	// 如果没有实现这个接口，fmt将以默认的方式输出。

	//实现同样的功能
	// fmt.Println("The biggest one is", boxes.BiggestsColor().String())
	// fmt.Println("The biggest one is", boxes.BiggestsColor())
	//注：实现了error接口的对象（即实现了Error() string的对象），使用fmt输出时，会调用Error()方法，因此不必再定义String()方法了。
}

func PrintAll(vals []interface{}) {
	for _, val := range vals {
		fmt.Println(val)
	}
}

//func test_interface_parameter() {
//	names := []string{"stanley", "david", "oscar"}
//    PrintAll(names)

//	// 56: cannot use names (type []string)
//	// as type []interface {} in argument to PrintAll
//}

func test_interface_parameter2() {
	names := []string{"stanley", "david", "oscar"}
	vals := make([]interface{}, len(names)) // vals
	for i, v := range names {
		vals[i] = v // 这样的代码特别丑
	}
	PrintAll(vals)
}

func main() {
	//	test_interface_Stringer()
	//	test_interface_parameter()
	test_interface_parameter2()
}
