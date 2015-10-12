package main

import (
	"fmt"
	"math/rand"
	"time"
	//	"encoding/json"
	//	"log"
)

//
// 匿名函数
//
func anonymous_func() {
	x := 5
	fn := func() {
		fmt.Println("x is", x)
	}
	fn()
	x++
	fn()
}

//
// 函数集合 - 自建函数类型
// we define the type "binFunc", which is a binary function;
// it takes two integers, and returns one integer.
// This isn't strictly necessary for this example,
// but typing `binFunc` over and over again is much more
// convenient than typing `func(int, int) int` everywhere you see it.
//
type binFunc func(int, int) int

func func_collections() {
	// seed your random number generator.
	// note:  this does not work on the Go Playground, because time is
	// frozen on the Go Playground.  time.Now() always returns the same
	// value.
	rand.Seed(time.Now().Unix())

	// create a slice of functions
	fns := []binFunc{
		func(x, y int) int { return x + y },
		func(x, y int) int { return x - y },
		func(x, y int) int { return x * y },
		func(x, y int) int { return x / y },
		func(x, y int) int { return x % y },
	}

	// pick one of those functions at random
	f := fns[rand.Intn(len(fns))]

	x, y := 12, 5
	fmt.Println(f(x, y))
}

//
// 函数集合2 - 自建struct类型
//
type op struct {
	name string
	fn   func(int, int) int
}

func func_collections2() {
	// seed your random number generator
	rand.Seed(time.Now().Unix())

	// create a slice of ops
	ops := []op{
		{"add", func(x, y int) int { return x + y }},
		{"sub", func(x, y int) int { return x - y }},
		{"mul", func(x, y int) int { return x * y }},
		{"div", func(x, y int) int { return x / y }},
		{"mod", func(x, y int) int { return x % y }},
	}

	// pick one of those ops at random
	o := ops[rand.Intn(len(ops))]

	x, y := 12, 5
	fmt.Println(o.name, x, y)
	fmt.Println(o.fn(x, y))
}

//
// 递归函数
//
type walkFn func(*int) walkFn

func pickRandom(fns ...walkFn) walkFn {
	return fns[rand.Intn(len(fns))]
}

func walkEqual(i *int) walkFn {
	*i += rand.Intn(7) - 3
	return pickRandom(walkForward, walkBackward)
}

func walkForward(i *int) walkFn {
	*i += rand.Intn(6)
	return pickRandom(walkEqual, walkBackward)
}

func walkBackward(i *int) walkFn {
	*i += -rand.Intn(6)
	return pickRandom(walkEqual, walkForward)
}

func func_recursive() {
	// time is frozen on playground, so this is actually always
	// the same.  The behavior is different when run locally.
	rand.Seed(time.Now().Unix())

	fn, progress := walkEqual, 0
	for i := 0; i < 20; i++ {
		fn = fn(&progress)
		fmt.Println(progress)
	}
}

//
// golang中:
// 1. 一切皆为类型对象
// 2. function为一等公民。
// 3. type是对归属于该type对象的最直接抽象。
// 4. 任意类型可以拥有methods来实现一个interfaces。
//
// golang任意类型对象都有值，对不？
// 对，未初始化的函数类型变量值 -- nil (空值 - 特殊的值)
// 函数类型作为接口值。
// 任意类型都拥有Error() string方法（method）。
// The error built-in interface type is the conventional interface
// for representing an error condition,
// with the nil value representing no error.
//
// error is an interface:
//
type error1 interface {
	Error() string
	Error2() string
}

//
// This means that any type which has a method:
// Error() string fulfills the interface and
// can be assigned to a variable of type error.
//

func (f binFunc) Error() string {
	fmt.Println("binFunc error")
	return "err1"
}

func (f binFunc) Error2() string {
	fmt.Println("binFunc error2")
	return "err2"
}

func add(x, y int) int {
	return x + y
}

func func_type_convert() {
	var err error1
	var i int
	//	fmt.Println(binFunc(add))
	err = binFunc(add) //函数类型convert：add转换成binFunc类型
	fmt.Println(err)
	fmt.Println(i)
}

//
// 从发jsonDoc中分解binFunc
//
var jsonDoc = []byte(`"add"`)

// 函数本身没有任何持久化数据，我们需要一些地方来注册函数名字。

var registry = map[string]binFunc{
	"add": func(x, y int) int { return x + y },
	"sub": func(x, y int) int { return x - y },
	"mul": func(x, y int) int { return x * y },
	"div": func(x, y int) int { return x / y },
	"mod": func(x, y int) int { return x % y },
}

//type binFunc func(int, int) int

// 函数类型对象值 -- nil
// 在binFunc函数类型上实现方法UnmarshalJSON
// 为了让改变能够被函数的调用者看到，
// 我们实现了带pointer receiver的方法UnmarshalJSON
//
// To unmarshal JSON into an interface value,
// Unmarshal stores one of these in the interface value:
//
//	bool, for JSON booleans
//	float64, for JSON numbers
//	string, for JSON strings
//	[]interface{}, for JSON arrays
//	map[string]interface{}, for JSON objects
//	nil for JSON null
//func (fn *binFunc) UnmarshalJSON(b []byte) error {
//	var name string
//	if err := json.Unmarshal(b, &name); err != nil {
//		return err
//	}
//
//	// get the function out of our function registry
//	found := registry[name]
//	if found == nil {
//		// return a descriptive error if we can't find the function
//		return fmt.Errorf("unknown function in (*binFunc)UnmarshalJSON: %s", name)
//	}
//
//	// dereference the pointer receiver, so that the changes are visible to the caller
//	*fn = found
//	return nil
//}
//
//func unmarshal_binFunc() {
//	var fn binFunc
//	fmt.Println("jsonDoc: " + string(jsonDoc))
//	fmt.Println(json.Unmarshal(jsonDoc, &fn))
//	fmt.Println(fn)
//	if err := json.Unmarshal(jsonDoc, &fn); err != nil {
//		log.Fatal(err)
//	}
//
//	x := fn(5, 12)
//	fmt.Println(x)
//}

func main() {
	//	anonymous_func() // 匿名函数
	//	func_collections() // 函数集合，代码很丑，但说明函数是一等公民
	//	func_collections2() // 函数集合，函数可以存在struct里，
	// 也可以存放在map里map[string]binFunc
	//	func_recursive() // 递归函数
	func_type_convert() // 函数类型作为接口值
	//	unmarshal_binFunc() // 从发json doc中分解binFunc

}
