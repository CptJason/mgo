//
//答案：如下
//package main
//
//import "fmt"
//
//type IPAddr [4]byte
//
// TODO: Add a "String() string" method to IPAddr.
//func (p IPAddr) String() string {
//    return fmt.Sprintf("%d.%d.%d.%d",p[0], p[1], p[2], p[3])
//
//}
//
//func main() {
//	addrs := map[string]IPAddr{
//		"loopback":  {127, 0, 0, 1},
//		"googleDNS": {8, 8, 8, 8},
//	}
//	for n, a := range addrs {
//		fmt.Printf("%v: %v\n", n, a)
//	}
//}
//

// 自己完成的
package main

import "fmt"
import "reflect"
import "bytes"
import "strings"
import "strconv"

type IPAddr [4]byte

//
// TODO: Add a "String() string" method to IPAddr.
// 我们使用了 bytes 包的 Buffer 类型，避免大量的字符串连接操作
// （因为 Go 中字符串是不可变的）。
// 我们再看一下标准库的实现：
//

//func Join(str []string, sep string) string {
//    // 特殊情况应该做处理
//    if len(str) == 0 {
//        return ""
//    }
//    if len(str) == 1 {
//        return str[0]
//    }
//    buffer := bytes.NewBufferString(str[0])
//    for _, s := range str[1:] {
//        buffer.WriteString(sep)
//        buffer.WriteString(s)
//    }
//    return buffer.String()
//}
//

func (ip IPAddr) String() string {
	var ss []string
	for i, v := range ip {
		fmt.Println("i:", i)
		fmt.Println("v:", v)
		fmt.Println("str v:", string(v))
		// 查看v的类型
		fmt.Println("type of v:", reflect.TypeOf(v))
		fmt.Println("type of v:", reflect.TypeOf(string(v)))
		s := strconv.FormatInt(int64(v), 10)
		ss = append(ss, s)
		fmt.Println("################")
	}
	fmt.Println(ss)
	fmt.Sprintf("ss: %v", ss[:])
	return fmt.Sprintf(strings.Join(ss, "."))
	// return fmt.Sprintf("%d.%d.%d.%d",ip_all[0], ip_all[1], ip_all[2], ip_all[3])
}

func main() {
	addrs := map[string]IPAddr{
		"loopback":  {127, 0, 0, 1},
		"googleDNS": {8, 8, 8, 8},
	}
	for n, a := range addrs {
		fmt.Printf("--------%v: %v\n", n, a)
	}
	type MyBytes []byte
	fmt.Println(string(MyBytes{'h', 'e', 'l', 'l', '\xc3', '\xb8'})) // "hellø"
	// join方法
	s := []string{"foo", "bar", "baz"}
	fmt.Println(strings.Join(s, ", "))
	// 注意: string(hello)中 - 0编码字符并没有打印出来
	fmt.Println("Hello, playground")
	hello := []byte{0, 'h', 'e', 'l', 'l', 'o'}
	fmt.Println(hello)
	fmt.Println(string(hello))
	fmt.Println(string(hello) == "hello")
	//
	// 为啥"127" 和 string(127)是不一样的
	//
	// string(127) means the character whose code is 127,
	// it is not printable. use strconv
	// string(127) 代表编码是127的字符，不可打印。
	//
	fmt.Println(strings.Join([]string{"name=xxx", "127"}, "&"))
	fmt.Println(strings.Join([]string{"name=xxx", string(127)}, "&"))
	// 正确做法如下
	fmt.Println(strings.Join([]string{"name=xxx", strconv.FormatInt(127, 10)}, "&"))
	a := bytes.Trim([]byte{'h', 'e', 'l', 'l', 'o'}, "lo")
	fmt.Println(a)
	bb := []byte{'h', 'e', 'l', 'l', 'o', 127}
	ss := string(bb)
	fmt.Println(ss)
}
