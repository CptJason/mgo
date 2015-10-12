package main

import "fmt"

func main() {
	//
	// slice -- 无固定长度数组，即所谓的动态数组
	// [] 方括号中没有数字
	// 一个 slice 会指向一个序列的值，并且包含了长度信息
	// []T 是一个元素类型为 T 的 slice
	//
	s := []int{2, 3, 5, 7, 11, 13}
	fmt.Println("s ==", s)

	for i := 0; i < len(s); i++ {
		fmt.Printf("s[%d] == %d\n", i, s[i])
	}
}
