package main

import "fmt"

func main() {
	pow := make([]int, 10)
	fmt.Println("pow", pow)
	fmt.Printf("-----\n")
	// 取元素下标 -- i, 元素值 -- v
	for i, v := range pow {
		fmt.Printf("pow[%d] == %d\n", i, v)
	}
	// 只取元素下标 -- i
	for i := range pow {
		fmt.Printf("%d\n", i)
	}
	// 只取元素值 -- v
	for _, v := range pow {
		fmt.Printf("%d\n", v)
	}
	// 求值平方
	for i := range pow {
		pow[i] = 1 << uint(i)
	}
}
