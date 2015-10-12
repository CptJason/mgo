//switch测试
//
//最好的讲解就是代码例子，现在让我们重写上面的这个实现

package main

import (
	"fmt"
	"strconv"
)

type Element interface{}
type List []Element

type Person2 struct {
	name string
	age  int
}

//打印
func (p Person2) String() string {
	return "(name: " + p.name + " - age: " + strconv.Itoa(p.age) + " years)"
}

func test_interface_type_switch() {
	list := make(List, 3)
	list[0] = 1       //an int
	list[1] = "Hello" //a string
	list[2] = Person2{"Dennis", 70}

	for index, element := range list {
		switch value := element.(type) {
		case int:
			fmt.Printf("me list[%d] is an int and its value is %d\n", index, value)
		case string:
			fmt.Printf("list[%d] is a string and its value is %s\n", index, value)
		case Person2:
			fmt.Printf("list[%d] is a Person and its value is %s\n", index, value)
		default:
			fmt.Println("list[%d] is of a different type", index)
		}
	}
}

// 这里有一点需要强调的是：element.(type)语法不能在switch外的任何逻辑里面使用，
// 如果你要在switch外面判断一个类型就使用comma-ok。

func main() {
	test_interface_type_switch()
}
