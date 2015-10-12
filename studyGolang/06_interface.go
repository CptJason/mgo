package main

import (
	"fmt"
)

//
// 2.6 interface
// https://github.com/astaxie/build-web-application-with-golang/blob/master/zh/02.6.md
//

// 什么是interface
//
// 简单的说，interface是一组method的组合，
// 我们通过interface来定义对象的一组行为。
// 我们前面一章最后一个例子中Student和Employee都能SayHi，
// 虽然他们的内部实现不一样，但是那不重要，重要的是他们都能say hi
// 让我们来继续做更多的扩展，Student和Employee实现另一个方法Sing，
// 然后Student实现方法BorrowMoney而Employee实现SpendSalary。
// 这样Student实现了三个方法：SayHi、Sing、BorrowMoney；
// 而Employee实现了SayHi、Sing、SpendSalary。
// 上面这些方法的组合称为interface(被对象Student和Employee实现)。
// 例如Student和Employee都实现了interface：SayHi和Sing，
// 也就是这两个对象是该interface类型。
//

type Human struct {
	name  string
	age   int
	phone string
}

type Student struct {
	Human  //匿名字段Human
	school string
	loan   float32
}

type Employee struct {
	Human   //匿名字段Human
	company string
	money   float32
}

//Human对象实现Sayhi方法
func (h *Human) SayHi3() {
	fmt.Printf("Hi, I am %s you can call me on %s\n", h.name, h.phone)
}

// Human对象实现Sing方法
func (h *Human) Sing1(lyrics string) {
	fmt.Println("La la, la la la, la la la la la...", lyrics)
}

//Human对象实现Guzzle方法
func (h *Human) Guzzle(beerStein string) {
	fmt.Println("Guzzle Guzzle Guzzle...", beerStein)
}

// Employee重载Human的Sayhi方法
func (e *Employee) SayHi3() {
	fmt.Printf("Hi, I am %s, I work at %s. Call me on %s\n", e.name,
		e.company, e.phone) //此句可以分成多行
}

//Student实现BorrowMoney方法
func (s *Student) BorrowMoney(amount float32) {
	s.loan += amount // (again and again and...)
}

//Employee实现SpendSalary方法
func (e *Employee) SpendSalary(amount float32) {
	e.money -= amount // More vodka please!!! Get me through the day!
}

// 定义interface
// interface可以被任意的对象实现。
// 我们看到上面的Men interface被Human、Student和Employee实现。
type People interface {
	SayHi3()
}

type Men interface {
	SayHi3()
	Sing(lyrics string)
	Guzzle(beerStein string)
}

type YoungChap interface {
	SayHi3()
	Sing(song string)
	BorrowMoney(amount float32)
}

type ElderlyGent interface {
	SayHi3()
	Sing(song string)
	SpendSalary(amount float32)
}

// 同理，一个对象可以实现任意多个Employee，
// 例如上面的Student实现了Men和YoungChap两个interface。

// Sample program to show how to use an interface in Go. In this case,
// a pointer is used to support the interface call.

// notifier is an interface that defined notification
// type behavior.
type notifier interface {
	notify1()
}

// user defines a user in the program.
type user struct {
	name  string
	email string
}

// notify implements a method that can be called via
// a value of type user.
func (u *user) notify1() {
	fmt.Printf("user: Sending user Email To %s<%s>\n",
		u.name,
		u.email)
}

// sendNotification accepts values that implement the notifier
// interface and sends notifications.
func sendNotification(n notifier) {
	n.notify1()
}

// main is the entry point for the application.
func main() {
	// Create a value of type user.
	u := &user{"Jill", "jill@email.com"}

	// Pass a pointer of of type user to the function.
	sendNotification(u)
}
