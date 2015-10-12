package main

import "fmt"

type Human struct {
	name  string
	age   int
	phone string
}

type Student struct {
	Human  //匿名字段 (anonymous field)
	school string
	loan   float32
}

type Employee struct {
	Human   //匿名字段
	company string
	money   float32
}

//Human实现SayHi方法
func (h Human) SayHi() {
	fmt.Printf("Hi, I am %s you can call me on %s\n", h.name, h.phone)
}

//Human实现Sing方法
func (h Human) Sing(lyrics string) {
	fmt.Println("La la la la...", lyrics)
}

//Employee重载Human的SayHi方法
func (e Employee) SayHi() {
	fmt.Printf("Hi, I am %s, I work at %s. Call me on %s\n", e.name,
		e.company, e.phone)
}

// Interface Men被Human,Student和Employee实现
// 因为这三个类型都实现了这两个方法
type Men interface {
	SayHi()
	Sing(lyrics string)
}

func test_interface_basic() {
	men := []Men{
		Human{"Mike", 25, "222-222-XXX"},
		Student{Human{"Paul", 26, "111-222-XXX"}, "Harvard", 100},
		Employee{Human{"Sam", 36, "444-222-XXX"}, "Golang Inc.", 1000},
	}
	for _, m := range men {
		m.SayHi()
	}
}

func test_var_interface() {
	mike := Student{Human{"Mike", 25, "222-222-XXX"}, "MIT", 0.00}
	paul := Student{Human{"Paul", 26, "111-222-XXX"}, "Harvard", 100}
	sam := Employee{Human{"Sam", 36, "444-222-XXX"}, "Golang Inc.", 1000}
	Tom := Employee{Human{"Tom", 37, "222-444-XXX"}, "Things Ltd.", 5000}

	//定义Men类型的变量i
	var i Men

	//i能存储Student
	i = mike
	fmt.Println("This is Mike, a Student:")
	i.SayHi()
	i.Sing("November rain")

	//i也能存储Employee
	i = Tom
	fmt.Println("This is Tom, an Employee:")
	i.SayHi()
	i.Sing("Born to be wild")

	//定义了slice Men
	fmt.Println("Let's use a slice of Men and see what happens")
	x := make([]Men, 3)
	//这三个都是不同类型的元素，但是他们实现了interface同一个接口
	x[0], x[1], x[2] = paul, sam, mike

	for _, value := range x {
		value.SayHi()
	}
}

func main() {
	//	test_interface_basic()
	//	test_var_interface()
	test_extending_type()
}

// notifier is an interface that defined notification
// type behavior.
type notifier interface {
	notify2()
}

// user defines a user in the program.
type user struct {
	name  string
	email string
}

// admin represents an admin user with privileges.
type admin struct {
	user
	level string
}

// notify implements a method that can be called via
// a value of type user.
func (u *user) notify2() {
	fmt.Printf("user: Sending user email To %s<%s>\n",
		u.name,
		u.email)
}

// main is the entry point for the application.
func test_extending_type() {
	// Create an admin user.
	ad := admin{
		user: user{
			name:  "john smith",
			email: "john@yahoo.com",
		},
		level: "super",
	}

	// Send the admin user a notification.
	// The embedded inner type's implementation of the
	// interface is "promoted" to the outer type.
	sendNotification2(&ad)
}

// sendNotification accepts values that implement the notifier
// interface and sends notifications.
func sendNotification2(n notifier) {
	n.notify2()
}
