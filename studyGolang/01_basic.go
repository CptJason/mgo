package main

import (
	"errors"
	"fmt"
)

// 分组声明
//import "net"
//import "html"

//
// 2.2 学习go基础
// https://github.com/astaxie/build-web-application-with-golang/blob/master/zh/02.2.md
//

func test_var_byte() {
	// rune (int32别称), int8, int16, int32, int64
	// 等于
	// byte (unit8),    unit8, unit16, unit32, unit64
	// 这些类型的变量之间不允许相互赋值或操作，
	// 否则编译报错，int(==int32),但int 与int32不可以互用。
	fmt.Println("----------byte----------")
	// 定义一个名称为a <variableName>，类型那个为int <type>的变量
	var a byte // 声明变量
	a = 15     // 变量赋值
	var d []byte
	d = append(d, 9)

	const (
		WHITE = iota
		BLACK
		BLUE
		RED
		YELLOW
	)

	type Color byte

	k := Color(WHITE)
	j := Color(BLACK)
	l := Color(YELLOW)

	//	var b byte // 声明变量
	//	b = false // byte类型变量不支持存放布尔值
	//	var c byte // 声明变量
	//	c = "test_byte" // byte类型变量不支持存放字符串

	fmt.Printf("整数a = %d\nslice第一个整数成员d = %d\n", a, d[0])
	fmt.Printf("color1 is: %d\ncolor2 is: %d\ncolor3 is :%d", k, j, l)

}

func test_var_int() {
	// rune (int32别称), int8, int16, int32, int64
	// byte (unit8),    unit8, unit16, unit32, unit64
	// 0 -- 2^(8)-1
	// 这些类型的变量之间不允许相互赋值或操作，
	// 否则编译报错，int(==int32),但int 与int32不可以互用。
	fmt.Println("----------int----------")
	// 定义一个名称为a <variableName>，类型那个为int <type>的变量
	var a int   // 声明变量
	var b bool  // 声明变量
	a = 15      // 变量赋值
	b = false   // 变量赋值
	c1 := 15    // 声赋 = 声明变量并赋值
	d1 := false // 声赋 = 声明变量并赋值
	fmt.Printf("整数a = %d\n布尔b = %t\n整数c1 = %d\n布尔d1 = %t\n", a, b, c1, d1)
}

var isActive bool                   // 全局变量声明
var enabled, disabled = true, false // 忽略类型的声明
func test_var_bool() {
	var available bool //一般声明
	valid := false     //简短声明
	available = true   // 复制操作
	fmt.Println("布尔isActive = %t\n布尔enabled = %t\n布尔disabled = %t\n布尔available = %t\n布尔valid = %t\n", isActive, enabled, disabled, available, valid)
}

func test_var_rune_clx() {
	var a1 rune = 16           // rune == int32
	var clx complex64 = 5 + 5i // 定义复数变量
	fmt.Printf("rune整数a1 = %d\n复数clx = %v\n", a1, clx)
}

func test_var_err() {
	// 定义一个错误
	fmt.Println("----------errors----------")
	err := errors.New("emit macho dwarf: elf header corrupted")
	if err != nil {
		fmt.Printf("error字串 = %s\n", err)
	}
}

func test_var_parallel_assign() {
	// 平行赋值
	fmt.Println("----------characters----------")
	var x1, y1 = "x1", "y1"
	var x, y string = "x", "y"
	x2, y2 := "x2", "y2"
	_, z := 34, 35
	fmt.Printf("字符串x = %s\n字符串y = %v\n", x, y)
	fmt.Printf("字符串x1 = %s\n字符串y1 = %v\n", x1, y1)
	fmt.Printf("字符串x2 = %s\n字符串y2 = %v\n", x2, y2)
	fmt.Printf("整数z = %d\n", z)
}

func test_var_conn_chars() {
	// 用 + 号连接两个字符串
	fmt.Println("----------use + conn characters----------")
	s4 := "hello"
	s5 := "world"
	sa := s4 + s5
	fmt.Printf("字符串s4 = %s\n字符串s5 = %s\n字符串sall = %s\n", s4, s5, sa)
}

func test_var_mul_lines_chars() {
	// 多行字符串
	fmt.Println("----------multiple lines characters----------")
	s1 := "starting part" +
		" ending part" // + 并不能实现多行字符串
	s2 := `starting part
	       ending part` // 输出多行字符串，包括换行符
	s3 := `starting part ending part` // 单行字符串
	fmt.Printf("字符串s1 = %s\n字符串s2 = %s\n字符串s3 = %s\n", s1, s2, s3)
}

func test_var_declare_in_group() {
	// 分组声明
	// 同时定义多个变量
	fmt.Println("----------declare multiple variables at the same time----------")
	var (
		h int  = 10000
		k bool = false
	)
	fmt.Printf("整数h = %d\n布尔k = %t\n", h, k)
}

func test_modify_chars_way1() {
	// 如何修改字符串方式1
	fmt.Println("----------change the chars way1----------")
	s := "hello"
	fmt.Printf("字符串s = %s\n", s)
	c := []byte(s) // 将字符串s 转换成 []byte类型数组
	fmt.Printf("s []byte is: %v\n", c)
	c[0] = 'c'
	s0 := string(c) // 再转换回string类型
	fmt.Printf("字符串s0 = %s\n", s0)
}

func test_modify_chars_way2() {
	// 如何修改字符串方式2
	fmt.Println("----------change the chars way2----------")
	s7 := "hello"
	s6 := "c" + s7[1:]
	fmt.Printf("字符串s = %s\n字符串s6 = %s\n", s7, s6)
}

func test_const() {
	const constantName = "value"
	const Pi float32 = 3.1415926
	const pi = 3.1415926
	const h = 10000
	const MaxThread = 10
	const prefix = "astaxie_"
	const (
		i = iota // 第一次为0
		l = iota // 第二次+1
		o = iota // 第三次又+1
	)
	const k = iota // 没遇到一次const关键字，iota就会重置
	fmt.Printf("常量constantName = %s\n常量Pi = %f\n常量pi = %f\n常量h = %d\n常量MaxThread = %d\n常量prefix = %s\n",
		constantName, Pi, pi, h, MaxThread, prefix)
	fmt.Printf("常量i = %d\n常量l = %d\n常量l = %d\n", i, l, o)
	fmt.Printf("iota被重置为0, 查看变量k:\n常量k = %d\n", k)
}

func test_array() {
	// 定义方式：var arrayName [n]type
	// 注长度是数组类型一部分,所以[3]int不同于[4]int
	// 数组不能改变长度
	// 数组之间的赋值是值的赋值 - 当把数组作为参数传入函数，传入的是副本，不是指针
	var arr [10]int // 声明一个int类型的数据
	arr[0] = 42     //数组下标从0开始
	arr[1] = 13     //复制操作
	fmt.Printf("The first element is %d\n", arr[0])
	fmt.Printf("The last element is %d\n", arr[9]) // 返回未赋值的最后一个元素，默认返回0
	a := [3]int{1, 2, 3}                           //声明一个长度为3的int数组
	b := [10]int{1, 2, 3}                          // 声明一个长度为10的int数组，其中前三个元素初始化为1、2、3, 其它默认为0
	c := [...]int{4, 5, 6}                         // 可以省略长度而采用...的方式, Go语言会自动根据元素个数来计算长度
	fmt.Printf("The first element is %d\n", a[0])
	fmt.Printf("The first element is %d\n", b[0])
	fmt.Printf("The first element is %d\n", c[0])
}

func test_double_array() {
	//	embebed_arr := [4]int{9, 10, 11, 12}
	doubleArray := [2][4]int{[4]int{1, 2, 3, 4}, [4]int{5, 6, 7, 8}}
	fmt.Printf("The first element is %d\n", doubleArray[0][0])
	fmt.Printf("The last element is %d\n", doubleArray[1][3])
}

func test_slice() {
	// 定义方式：var sliceName []type
	// 和声明array一样，只是无需定义长度
	// 初始定义数组并不知道需要多大 -- 这就需要"动态数组"
	// golang用slice数据结构来实现，但它并不是真的动态数组，而是一个引用类型
	var fslise []int                                                    //
	slice := []byte{'a', 'b', 'c', 'd'}                                 // 声明一个slice, 并初始化数据
	var ar = [10]byte{'a', 'b', 'c', 'd', 'e', 'f', 'g', 'h', 'i', 'j'} // 声明一个含有10个元素元素类型为byte的数组
	var a, b []byte                                                     // 声明两个含有byte的slice
	a = ar[2:5]                                                         // a指向数组的第三个元素开始, 第5个元素结束; 现有a含有的元素：ar[2]、ar[3]、 ar[4]
	b = ar[3:5]                                                         // b是数组ar的另一个slice; b的元素是 ar[3] ar[4]
	fmt.Println("----------printf----------")
	fmt.Printf("fslice = %s\nslice = %s\n", fslise, slice)
	fmt.Println("----------println----------")
	fmt.Println("fslice =", fslise, "\nslice =", slice)
	fmt.Println("----------printf slice a & b----------")
	fmt.Printf("slice a = %s\n", a)
	fmt.Printf("slice b = %s\n", b)
	fmt.Println("----------println slice a & b----------")
	fmt.Println("slice a = ", a)
	fmt.Println("slice b = ", b)
	fmt.Println("----------printf slice a----------")
	fmt.Printf("slice a 's first element is %d\n", a[0])
	fmt.Printf("slice a 's first element is %s\n", string(a[0]))
	fmt.Println("----------printf slice b----------")
	fmt.Printf("slice b 's first element is %d\n", b[0])
	fmt.Printf("slice b 's first element is %s\n", string(b[0]))
}

func test_slice_op() {
	// slice的默认开始位置是0, ar[:n] = ar[0:n]
	// slice的第二个序列默认是数组的长度，ar[n:] = ar[n:len(ar)]
	// 从一个数组直接获取slice, ar[:] = ar[0:len(ar)]
	var array = [10]byte{'a', 'b', 'c', 'd', 'e', 'f', 'g', 'h', 'i', 'j'} // 声明一个含有10个元素元素类型为byte的数组
	var aSlice, bSlice []byte                                              //声明两个slice
	aSlice = array[:3]                                                     // 等价于aSlice = array[0:3]  aSlice包含元素：a,b,c
	aSlice = array[5:]                                                     // 等价于aSlice = array[5:10]  aSlice包含元素：f,g,h,i,j
	aSlice = array[:]                                                      // 等价aSlice = array[0:10] 这样aSlice包含了全部的元素
	//从Slice中获取Slice
	aSlice = array[3:7] // aSlice包含元素：d,e,f,g, len=4, cap=7
	fmt.Println("----------println bSlice----------")
	bSlice = aSlice[1:3] // bSlice包含元素：aSlice[1], aSlice[2] 也就是含有：e,f
	fmt.Println("slice b = ", bSlice)
	bSlice = aSlice[:3] // bSlice包含元素：aSlice[0], aSlice[1], aSlice[2] 也就是含有：d,e,f
	fmt.Println("slice b = ", bSlice)
	bSlice = aSlice[:] // bSlice包含所有aSlice的元素：d,e,f,g
	fmt.Println("slice b = ", bSlice)
	bSlice = aSlice[0:5] // 对slice的slice可以在cap范围内扩展，此时bSlice包含元素：d,e,f,g,h
	fmt.Println("slice b = ", bSlice)
	bSlice = array[2:4:7] // 支持三个参数slice
	fmt.Printf("three slice b 暂时不支持 = %s\n", bSlice)

	fmt.Println("----------printf bSlice b----------")
	fmt.Printf("slice b 's first element is %d\n", bSlice[0])
	fmt.Printf("slice b 's first element is %s\n", string(bSlice[0]))
	fmt.Println("----------printf bSlice----------")
	fmt.Printf("slice b = %s\n", bSlice)
	fmt.Println("----------println bSlice----------")
	fmt.Println("slice b = ", bSlice)
	// slice是引用类型
	// 当引用改变其中元素的值时，其他的所有引用都会改变该值，如，aSlice, bSlice
}

func test_map() {
	//	map也就是Python中字典的概念，
	// 格式为
	// map[keyType]valueType
	//	我们看下面的代码，map的读取和设置也类似slice一样，通过key来操作
	// 只是slice的index只能是｀int｀类型，而map多了很多类型，
	// 可以是int，可以是string及所有完全定义了==与!=操作的类型。
	//	var numbers map[string]int // 声明一个key是字符串，值为int的字典,
	// 这种方式的声明需要在使用之前使用make初始化
	// 另一种map的声明方式
	numbers := make(map[string]int)
	numbers["one"] = 1  //赋值
	numbers["ten"] = 10 //赋值
	numbers["three"] = 3

	fmt.Println("第三个数字是: ", numbers["three"]) // 读取数据
	// 打印出来如:第三个数字是: 3
	//	这个map就像我们平常看到的表格一样，左边列是key，右边列是值
	//
	//	使用map过程中需要注意的几点：
	//
	//	1. map是无序的，每次打印出来的map都会不一样，它不能通过index获取，而必须通过key获取
	//	2. map的长度是不固定的，也就是和slice一样，也是一种引用类型
	//	3. 内置的len函数同样适用于map，返回map拥有的key的数量
	//	4. map的值可以很方便的修改，通过numbers["one"]=11可以很容易的把key为one的字典值改为11
	//	5. map和其他基本型别不同，它不是thread-safe，在多个go-routine存取时，必须使用mutex lock机制
	//	6. map的初始化可以通过key:val的方式初始化值，同时map内置有判断是否存在key的方式
}

func test_map_delete() {
	rating := map[string]float32{"C": 5, "Go": 4.5, "Python": 4.5, "C++": 2} // 初始化一个字典
	csharpRating, ok := rating["C#"]                                         // map有两个返回值，第二个返回值，
	// 如果不存在key，那么ok为false，如果存在ok为true
	if ok {
		fmt.Println("C# is in the map and its rating is ", csharpRating)
	} else {
		fmt.Println("We have no rating associated with C# in the map")
	}

	delete(rating, "C") // 删除key为C的元素
}

func test_map_assign() {
	m := make(map[string]string)
	m["Hello"] = "Bonjour"
	m1 := m
	m1["Hello"] = "Salut" // 现在m["hello"]的值已经是Salut了
	fmt.Println("what is m[\"Hello\"] now: ", m["Hello"])
}

func make_vs_new() {
	// make用于 <内建类型>（map、slice 和channel）的内存分配 - make(T, args)
	// new用于  <各种类型> 的内存分配 - new(T)
	//
	// 内建函数new本质上说跟其它语言中的同名函数功能一样：
	// new(T)分配了零值填充的T类型的内存空间，并且返回其地址，即一个*T类型的值。
	// 用Go的术语说，它返回了一个指针，指向新分配的类型T的零值。
	// 有一点非常重要：
	//	*******  new返回指针  *******

	//	内建函数make(T, args)与new(T)有着不同的功能，
	//  make只能创建slice、map和channel，并且返回一个有初始值(非零)的T类型，而不是*T。

	//  本质来讲，导致这三个类型有所不同的原因是指向数据结构的引用在使用前必须被初始化。
	//  例如，一个slice，是一个包含指向数据（内部array）的指针、长度和容量的三项描述符；
	//  在这些项目被初始化之前，slice为nil。对于slice、map和channel来说，
	//  make初始化了内部的数据结构，填充适当的值。
	//
	//	*******   make返回初始化后的（非零）值 *******
	//	下面这个图详细的解释了new和make之间的区别。
}

func test_zero() {
	// 关于“零值”，所指并非是空值，而是一种“变量未填充前”的默认值，通常为0。
	// 此处罗列 部分类型 的 “零值”
	//
	//	int     0 // int32
	//	int8    0
	//	int32   0
	//	int64   0
	//	uint    0x0
	//	rune    0 //rune的实际类型是 int32
	//	byte    0x0 // byte的实际类型是 uint8
	//	float32 0 //长度为 4 byte
	//	float64 0 //长度为 8 byte
	//	complex64 //长度为 8 byte
	//	bool    false
	//	string  ""

}

//func test_unused() {
//	var i int
//}

func main() {
	//	test_var_byte() // 定义byte类型
	//	test_var_int()  // 定义整数
	//	test_var_bool() // 定义布尔值
	//	test_var_rune_clx() // 定义复数
	//	test_var_err() // 定义错误
	//	test_var_parallel_assign() // 定义平行赋值
	//	test_var_conn_chars() // 定义连接字符串
	//	test_var_mul_lines_chars() // 定义多行字符串
	//	test_var_declare_in_group() // 分组定义
	test_modify_chars_way1() // 修改字符方式1
	//	test_modify_chars_way2()// 修改字符方式2

	//	test_const() // 定义常量
	//	test_array()
	//	test_double_array()
	//	test_slice()
	//	test_slice_op()
	//	test_map()
	//	test_map_delete()
	//	test_map_assign()
	//	make_vs_new()
	//	test_zero()
	//	test_unused()
	// go之所以简洁，是因为他有一些默认的行为
	// 大写字母开头变量为public变量, 小写字母开头为private变量
	// 大写字母开头函数为public函数, 小写字母开头为private函数

}
