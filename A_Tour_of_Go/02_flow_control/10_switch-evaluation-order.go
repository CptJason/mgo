package main

import (
	"fmt"
	"time"
)

func main() {
	fmt.Println("When's Saturday?")
	today := time.Now().Weekday()
	fmt.Println("today: ", today)
	fmt.Println("today + 0: ", today+0)
	fmt.Println("today - 1: ", today-1)
	//	fmt.Println("today + 1: ", today + 1) // PANIC=runtime error: index out of range
	fmt.Println("time.Saturday: ", time.Saturday)
	// switch 的条件从上到下的执行
	// 当匹配成功的时候停止
	switch time.Saturday {
	case today + 0:
		fmt.Println("Today.")
	case today + 1:
		fmt.Println("Tomorrow.")
	case today + 2:
		fmt.Println("In two days.")
	default:
		fmt.Println("Too far away.")
	}
}
