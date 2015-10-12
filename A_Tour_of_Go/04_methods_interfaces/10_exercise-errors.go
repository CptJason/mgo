//package main
//
//import (
//	"fmt"
//)
//
//type ErrNegativeSqrt float64
//
//func (e ErrNegativeSqrt) Error() string {
//	return fmt.Sprintf("cannot Sqrt negativ number: %g", float64(e))
//}
//
//func Sqrt(f float64) (float64, error) {
//	if f < 0 {
//		return 0, ErrNegativeSqrt(f)
//	}
//	return 0, nil
//}
//
//func main() {
//	fmt.Println(Sqrt(2))
//	fmt.Println(Sqrt(-2))
//}

package main

import (
	"fmt"
	//	"math"
)

//
// Exercise: Errors
// Copy your Sqrt function from the earlier exercise and
// modify it to return an error value.
//
// Sqrt should return a non-nil error value when given a negative number,
// as it doesn't support complex numbers.
//
// Create a new type
//
// type ErrNegativeSqrt float64
// and make it an error by giving it a
//
// func (e ErrNegativeSqrt) Error() string
// method such that ErrNegativeSqrt(-2).Error()
// returns "cannot Sqrt negative number: -2".
//
// Note:
// a call to fmt.Sprint(e) inside the Error method
// will send the program into an infinite loop.
// You can avoid this by converting e first: fmt.Sprint(float64(e)). Why?
//
// Change your Sqrt function to return an ErrNegativeSqrt value
// when given a negative number.
//
// 自己实现代码
//
type ErrNegativeSqrt float64

func (e ErrNegativeSqrt) Error() string {
	return fmt.Sprintf("cannot Sqrt negative number:  %g", e)
}

func Sqrt(x ErrNegativeSqrt) (float64, error) {
	if x < 0 {
		return 0, ErrNegativeSqrt(x)
	}
	return 0, nil
}

func Sqrt2(x ErrNegativeSqrt) (ErrNegativeSqrt, error) {
	if x < 0 {
		return 0, ErrNegativeSqrt(x)
	}
	return 0, nil
}

func main() {
	// 理解Sqrt和Sqrt2输出之间为啥有这种区别，来理解interface类型返回值
	fmt.Println(Sqrt(2))
	fmt.Println(Sqrt(-2))
	fmt.Println(Sqrt2(2))
	fmt.Println(Sqrt2(-2))
}
