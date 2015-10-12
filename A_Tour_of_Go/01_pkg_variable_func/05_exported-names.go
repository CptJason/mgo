package main

import (
	"fmt"
	"math"
)

func main() {
	//	fmt.Println(math.pi) // cannot refer to unexported name math.pi
	fmt.Println(math.Pi) // 首字母大写的名称是被导出的
}
