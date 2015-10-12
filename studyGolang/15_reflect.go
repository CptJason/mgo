package main

import (
	"fmt"
	"reflect"
)

//反射
//
//Go语言实现了反射，所谓反射就是能检查程序在运行时的状态。我们一般用到的包是reflect包。如何运用reflect包，官方的这篇文章详细的讲解了reflect包的实现原理，laws of reflection
//
// https://github.com/astaxie/build-web-application-with-golang/blob/master/zh/02.6.md
//使用reflect一般分成三步，下面简要的讲解一下：要去反射是一个类型的值(这些值都实现了空interface)，首先需要把它转化成reflect对象(reflect.Type或者reflect.Value，根据不同的情况调用不同的函数)。这两种获取方式如下：
func main() {

	var i int = 10
	t := reflect.TypeOf(i)  //得到类型的元数据,通过t我们能获取类型定义里面的所有元素
	v := reflect.ValueOf(i) //得到实际的值，通过v我们获取存储在里面的值，还可以去改变值
	//转化为reflect对象之后我们就可以进行一些操作了，也就是将reflect对象转化成相应的值，例如
	fmt.Printf("%v,%v\n---\n", t, v)

	//	tag := t.Elem().Field(0).Tag  //获取定义在struct里面的标签
	//	name := v.Elem().Field(0).String()  //获取存储在第一个字段里面的值
	//	//获取反射值能返回相应的类型和数值
	//	fmt.Printf("%v, %v", tag , name)

	var x float64 = 3.4
	x_value := reflect.ValueOf(x)
	fmt.Printf("%v\n", x_value)
	fmt.Printf("%v\n", x_value.Type())
	fmt.Printf("%v\n", x_value.Kind() == reflect.Float64)
	fmt.Printf("%v\n", x_value.Float())

	// 最后，反射的话，那么反射的字段必须是可修改的，我们前面学习过传值和传引用，
	// 这个里面也是一样的道理。反射的字段必须是可读写的意思是，如果下面这样写，那么会发生错误

	var y float64 = 3.4
	y_value := reflect.ValueOf(y)
	//	y_value.SetFloat(7.1)
	fmt.Printf("%v", y, "%v", y_value)
	//如果要修改相应的值，必须这样写
	//
	//	var z float64 = 3.4
	//	z_value := reflect.ValueOf(&z)
	//	v4 := z_value.Elem()
	//	v4.SetFloat(7.1)
	//	// 上面只是对反射的简单介绍，更深入的理解还需要自己在编程中不断的实践。
	//	fmt.Printf("%v", z, "%v", z_value)
}
