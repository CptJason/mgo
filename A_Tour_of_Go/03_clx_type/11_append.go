package main

import (
	"fmt"
)

// 将数据追加到切片的尾部
func AppendByte(slice []byte, data ...byte) []byte {
	m := len(slice)
	n := m + len(data)
	if n > cap(slice) { // if necessary, reallocate
		// allocate double what's needed, for future growth.
		newSlice := make([]byte, (n+1)*2)
		copy(newSlice, slice)
		slice = newSlice
	}
	slice = slice[0:n]
	copy(slice[m:n], data)
	return slice
}

func printSlice2(s string, x []int) {
	fmt.Printf("%s len=%d cap=%d %v\n",
		s, len(x), cap(x), x)
}

func main() {
	//
	// 向 slice 添加元素
	// 向 slice 添加元素是一种常见的操作，因此 Go 提供了一个内建函数 append 。
	// 内建函数的文档对 append 有详细介绍。
	// func append(s []T, vs ...T) []T
	//
	// append 的第一个参数 s 是一个类型为 T 的数组，
	// 其余类型为 T 的值将会添加到 slice。
	//
	// append 的结果是一个包含原 slice 所有元素加上新添加的元素的 slice。
	// 如果 s 的底层数组太小，而不能容纳所有值时，会分配一个更大的数组。
	// 返回的 slice 会指向这个新分配的数组。
	//（了解更多关于 slice 的内容。）
	//
	// 参阅文章Go 切片：用法和本质:
	// https://blog.go-zh.org/go-slices-usage-and-internals
	//

	var a []int
	printSlice2("a", a)

	// append works on nil slices.
	a = append(a, 0)
	printSlice2("a", a)

	// the slice grows as needed.
	a = append(a, 1)
	printSlice2("a", a)

	// we can add more than one element at a time.
	a = append(a, 2, 3, 4)
	printSlice2("a", a)

	s := make([]byte, 5)
	//
	s = AppendByte(s, 7, 11, 13, 'g', 0)
	fmt.Println("s", s)
}
