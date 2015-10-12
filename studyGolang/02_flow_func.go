package main

import (
	"fmt"
	"log"
	"os"
)

// 分组声明
//import "net"
//import "html"

//
// 2.3 流程和函数
// https://github.com/astaxie/build-web-application-with-golang/blob/master/zh/02.3.md
//

func test_if() {
	var x = 1
	var y = 2
	var name = "/tmp/test"

	if true && true {
		println("true && true = true")
	}

	if true && false {
		println("true && false = true")
	} else {
		println("true && false = false")
	}

	if true || true {
		println("true || true = true")
	}

	if true || false {
		println("true || false = true")
	} else {
		println("true || false = false")
	}

	if !false {
		println("true")
	}

	if x > 0 {
		println(x, y)
	}

	if x > 10 {
		fmt.Println("x is greater than 10")
	} else {
		fmt.Println("x is less than 10")
	}

	if a, b := 21, 3; a > b {
		fmt.Println("a>b ? true")
	}

	if a, b := 21, 31; a > b {
		fmt.Println("a>b ? true")
	} else {
		fmt.Println("a>b ? false")
		fmt.Println(a, b) //Ok
	}
	//	fmt.Println(a,b)    //error: undefined a ,undefined b

	// 计算获取值x,然后根据x返回的大小，判断是否大于10。
	if x := 15; x > 10 {
		fmt.Println("x is greater than 10")
	} else {
		fmt.Println("x is less than 10")
	}

	//这个地方如果这样调用就编译出错了，因为x是条件里面的变量
	fmt.Println(x)

	f, err := os.Open(name)
	println("f: ", f)
	if err != nil {
		return
	}

}

func test_multi_if() {
	integer := 10
	if integer == 3 {
		fmt.Println("The integer is equal to 3")
	} else if integer < 3 {
		fmt.Println("The integer is less than 3")
	} else {
		fmt.Println("The integer is greater than 3")
	}
}

func getName(id int) string {
	if id == 1 {
		return "YourName"
	} else {
		return "MyName"
	}
	//	panic("function ends without a return statement")
}

func test_if_func(id int) {
	name := getName(id)
	fmt.Println(name)
}

func test_if_file() (err1 error, err2 error) {
	file, err1 := os.Open("/Users/yinhezhixing/go/src/github.com/prometheus/prometheus/studyGolang/hello.go") // For read access.
	err2 = nil
	if err2 := file.Chmod(0664); err1 != nil && err2 != nil {
		log.Print(err1)
		log.Fatal(err2)
		return err1, err2
	}
	return err1, err2
} // miss return at the end of function

//func test_if_myMap() interface{}  {
//	myMap := make(map[string]int, 3)
//	myMap["one"] = 1
//	myMap["two"] = 2
//	myMap["three"] = 3
//	if v,err :=myMap["one"];err != nil {
//		log.Print(err)
//		return err
//	}else {
//		return v
//	}
//}

func test_goto() {
	n := 1

Here: // Goto标签
	fmt.Println(n)
	n++

	if n == 9 {
		fmt.Println("break goto loop")
		os.Exit(0) // break只能用于退出for循环
	}

	goto Here // 跳回标签Here, 重新开始执行

}

// 测试for, 类似while方式循环
func test_for_while() {
	a := 1
	b := 5
	for a < b {
		a *= 2
		println("test_for1,", a, b)
	}
	println("-----------")
	println("test_for1,", a, b)
}

// 测试正常for循环 break
// for expression1; expression2; expression3 {
// ...
// }
func test_for_break() {
	sum := 0
	for i := 0; i < 10; i++ {
		if i > 5 {
			break // 利用break提前退出循环
		}
		sum += i
	}
	println("test_for2", sum)
}

// 测试正常for循环 continue
func test_for_continue() {
	sum := 0
	for i := 0; i < 10; i++ {
		if i > 5 {
			continue // 利用break提前退出循环
		}
		sum += i
	}
	println("test_for3", sum)
}

// 测试for中平行赋值
func test_for_parallel_assign() {
	var a = []int{1, 2, 3, 4, 5}
	for i, j := 0, len(a)-1; i < j; i, j = i+1, j-1 {
		a[i], a[j] = a[j], a[i]
	}
	println("test_for3", a)
}

func test_for_label() {
J:
	for j := 0; j < 2; j++ {
		for i := 0; i < 10; i++ {
			if i > 5 {
				break J //退出外循环，内循环继续
			}
			println(i)
		}
	}
}

func test_switch() {
	i := 10
	switch i {
	case 1:
		fmt.Println("i is equal to 1")
	case 2, 3, 4:
		fmt.Println("i is equal to 2, 3 or 4")
	case 10:
		fmt.Println("i is equal to 10")
	default:
		fmt.Println("All I know is that i is an integer")
	}
}

func test_switch_fallthrough() {
	integer := 6
	switch integer {
	case 4:
		fmt.Println("The integer was <= 4")
		fallthrough
	case 5:
		fmt.Println("The integer was <= 5")
		fallthrough
	case 6:
		fmt.Println("The integer was <= 6")
		fallthrough
	case 7:
		fmt.Println("The integer was <= 7")
		fallthrough
	case 8:
		fmt.Println("The integer was <= 8")
		fallthrough
	default:
		fmt.Println("default case")
	}
}

// 函数
//	func funcName(input1 type1, input2 type2) (output1 type1, output2 type2) {
//		//这里是处理逻辑代码
//		//返回多个值
//		return value1, value2
//	}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func test_func_max() {
	var x, y int = 3, 5
	max_xy := max(x, y)
	fmt.Printf("max(%d, %d) = %d\n", x, y, max_xy)
}

//返回 A+B 和 A*B
func SumAndProduct(A, B int) (int, int) {
	return A + B, A * B
}

// 官方建议：最好命名返回值，因为不命名返回值，
// 虽然使得代码更加简洁了，但是会造成生成的文档可读性差。
func SumAndProduct2(A, B int) (add int, Multiplied int) {
	add = A + B
	Multiplied = A * B
	return
}

func test_multi_return() {
	x := 3
	y := 4

	xPLUSy, xTIMESy := SumAndProduct(x, y)

	fmt.Printf("%d + %d = %d\n", x, y, xPLUSy)
	fmt.Printf("%d * %d = %d\n", x, y, xTIMESy)
}

// 变参
// func myfunc(arg ...int) {}
func print_var_args(arg ...int) {
	for _, n := range arg {
		fmt.Printf("And the number is: %d\n", n)
	}
}

func test_print_var_args() {
	print_var_args(1, 2, 3, 4)
}

//简单的一个函数，实现了参数+1的操作
func add1(a int) int {
	a = a + 1 // 我们改变了a的值
	return a  //返回一个新值
}

func test_parameter_value() {
	x := 3

	fmt.Println("x = ", x) // 应该输出 "x = 3"

	x1 := add1(x) //调用add1(x)

	fmt.Println("x1: x+1 = ", x1) // 应该输出"x+1 = 4"
	fmt.Println("x = ", x)        // 应该输出"x = 3"
}

//简单的一个函数，实现了参数+1的操作
func add2(a *int) int { // 请注意，
	*a = *a + 1 // 修改了a的值
	return *a   // 返回新值
}

// 那么到底传指针有什么好处呢？
// 传指针使得多个函数能操作同一个对象。
// 传指针比较轻量级 (8bytes),只是传内存地址，
// 我们可以用指针传递体积大的结构体。如果用参数值传递的话,
// 在每次copy上面就会花费相对较多的系统开销（内存和时间）。
// 所以当你要传递大的结构体的时候，用指针是一个明智的选择。
// Go语言中channel，slice，map
// 这三种类型的实现机制类似指针，所以可以直接传递，
// 而不用取地址后传递指针。
// （注：若函数需改变slice的长度，则仍需要取地址传递指针）
func test_parameter_pointer() {
	x := 3

	fmt.Println("x = ", x) // 应该输出 "x = 3"

	x1 := add2(&x) // 调用 add1(&x) 传x的地址

	fmt.Println("x+1 = ", x1) // 应该输出 "x+1 = 4"
	fmt.Println("x = ", x)    // 应该输出 "x = 4"
}

//func ReadWrite() bool {
//	file.Open("/tmp/test.txt")
//	// 做一些工作
//	if failureX {
//		file.Close()
//		return false
//	}
//
//	if failureY {
//		file.Close()
//		return false
//	}
//
//	file.Close()
//	return true
//}
//
//func ReadWrite() bool {
//	file.Open("/tmp/test.txt")
//	defer file.Close()
//	if failureX {
//		return false
//	}
//	if failureY {
//		return false
//	}
//	return true
//}

func test_defer_for() {
	for i := 0; i < 5; i++ {
		defer fmt.Printf("%d ", i)
	}
}

// 函数作为值、类型
// 在Go中函数也是一种变量，我们可以通过type来定义它，
// 它的类型就是所有拥有相同的参数，相同的返回值的一种类型
// type typeName func(input1 inputType1 , input2 inputType2 [, ...]) (result1 resultType1 [, ...])
// 函数作为类型到底有什么好处呢？
// 那就是可以把这个类型的函数当做值来传递，请看下面的例子
type testInt func(int) bool // 声明了一个函数类型

func isOdd(integer int) bool {
	if integer%2 == 0 {
		return false
	}
	return true
}

func isEven(integer int) bool {
	if integer%2 == 0 {
		return true
	}
	return false
}

// 声明的函数类型在这个地方当做了一个参数
func filter(slice []int, f testInt) []int {
	var result []int
	for _, value := range slice {
		if f(value) {
			result = append(result, value)
		}
	}
	return result
}

func filter2(slice []int, f testInt) ([]int, []int) {
	var result []int
	var result2 []int
	for index, value := range slice {
		if f(value) {
			result = append(result, index)
			result2 = append(result2, value)
		}
	}
	return result, result2
}

// 函数当做值和类型在我们写一些通用接口的时候非常有用，
// 通过下面例子我们看到testInt这个类型是一个函数类型，
// 然后两个filter函数的参数和返回值与testInt类型是一样的，
// 但是我们可以实现很多种的逻辑，这样使得我们的程序变得非常的灵活。
func test_func_as_parameter() {
	slice := []int{1, 2, 3, 4, 5, 7}
	fmt.Println("slice = ", slice)
	odd := filter(slice, isOdd) // 函数当做值来传递了
	fmt.Println("Odd elements of slice are: ", odd)
	even := filter(slice, isEven) // 函数当做值来传递了
	fmt.Println("Even elements of slice are: ", even)

	odd_ok, odd := filter2(slice, isOdd) // 函数当做值来传递了
	fmt.Println("Odd2 index of slice are: ", odd_ok, "Odd2 elements of slice are: ", odd)
	even_ok, even := filter2(slice, isEven) // 函数当做值来传递了
	fmt.Println("Even2 index of slice are: ", even_ok, "Even2 elements of slice are: ", even)
}

func test_import() {
	// import 我们在写Go代码的时候经常用到import这个命令
	// 用来导入包文件，而我们经常看到的方式参考如下：
	//	import(
	//	"fmt"
	//	)
	//	然后我们代码里面可以通过如下的方式调用
	//	fmt.Println("hello world")
	//	上面这个fmt是Go语言的标准库，其实是去GOROOT环境变量指定目录下去加载该模块，当然Go的import还支持如下两种方式来加载自己写的模块：
	//
	//	相对路径
	//	import “./model” //当前文件同一目录的model目录，
	// 但是不建议这种方式来import
	//
	//	绝对路径
	//	import “shorturl/model”
	//  加载gopath/src/shorturl/model模块
	//	上面展示了一些import常用的几种方式，
	// 但是还有一些特殊的import，让很多新手很费解，
	// 下面我们来一一讲解一下到底是怎么一回事
	//
	//	点操作
	//	我们有时候会看到如下的方式导入包
	//	import(
	//	. "fmt"
	//	)
	//	这个点操作的含义就是这个包导入之后在你调用这个包的函数时，
	//  你可以省略前缀的包名，
	//  也就是前面你调用的fmt.Println("hello world")可以省略的写成Println("hello world")
	//
	//	别名操作
	//	别名操作顾名思义我们可以把包命名成另一个我们用起来容易记忆的名字
	//	import(
	//	f "fmt"
	//	)
	//	别名操作的话调用包函数时前缀变成了我们的前缀，即f.Println("hello world")
	//
	//	_操作
	//	这个操作经常是让很多人费解的一个操作符，请看下面这个import
	//	import (
	//	"database/sql"
	//	_ "github.com/ziutek/mymysql/godrv"
	//	)
	//	_操作其实是引入该包，而不直接使用包里面的函数，
	// 而是调用了该包里面的init函数。

}

func main() {
	//	test_if()
	//	test_multi_if()
	//	test_if_func(1)
	err1, err2 := test_if_file()
	fmt.Println(err1, err2)
	//	test_goto()
	//	test_for_while()
	//	test_for_break()
	//	test_for_continue()
	//	test_for_parallel_assign()
	//	test_for_label()
	//	test_switch()
	//	test_switch_fallthrough()
	//	test_func_max()
	//	test_multi_return()
	//	test_print_var_args()
	//	test_parameter_value()
	//	test_parameter_pointer()
	//	test_defer_for()
	//	test_func_as_parameter()
}
