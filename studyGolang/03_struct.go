package main

import "fmt"

//
// 2.4 struct
// https://github.com/astaxie/build-web-application-with-golang/blob/master/zh/02.4.md
//

// 声明一个struct如此简单，上面的类型包含有两个字段
//
// 一个string类型的字段name，用来保存用户名称这个属性
// 一个int类型的字段age,用来保存用户年龄这个属性
//如何使用struct呢？请看下面的代码

type person struct {
	name string
	age  int
}

func test_var_struct() {
	var P person // P现在就是person类型的变量了

	P.name = "Astaxie"                              // 赋值"Astaxie"给P的name属性.
	P.age = 25                                      // 赋值"25"给变量P的age属性
	fmt.Printf("The person's name is %s\n", P.name) // 访问P的name属性.

}

func test_init1_struct() {

	P := person{"Tom", 25}                          // 按照field顺序初始化初始值
	fmt.Printf("The person's name is %s\n", P.name) // 访问P的name属性.

}

func test_init2_struct() {

	P := person{age: 24, name: "Tom"}               // 按field:value初始化初始值
	fmt.Printf("The person's name is %s\n", P.name) // 访问P的name属性.

}

func test_init3_struct() {
	P := new(person)                                // 通过new函数分配一个指针，此处P的类型为*person
	P.name = "Astaxie"                              // 赋值"Astaxie"给P的name属性.
	P.age = 25                                      // 赋值"25"给变量P的age属性
	fmt.Printf("The person's name is %s\n", P.name) // 访问P的name属性.

}

// 比较两个人的年龄，返回年龄大的那个人，并且返回年龄差
// struct也是传值的
func Older(p1, p2 person) (person, int) {
	if p1.age > p2.age { // 比较p1和p2这两个人的年龄
		return p1, p1.age - p2.age
	}
	return p2, p2.age - p1.age
}

func test_struct_older() {
	var tom person

	// 赋值初始化
	tom.name, tom.age = "Tom", 18

	// 按照struct定义顺序初始化值
	paul := person{"Paul", 43}

	// 两个字段都写清楚的初始化
	bob := person{age: 25, name: "Bob"}

	tb_Older, tb_diff := Older(tom, bob)
	tp_Older, tp_diff := Older(tom, paul)
	bp_Older, bp_diff := Older(bob, paul)

	fmt.Printf("Of %s and %s, %s is older by %d years\n",
		tom.name, bob.name, tb_Older.name, tb_diff)

	fmt.Printf("Of %s and %s, %s is older by %d years\n",
		tom.name, paul.name, tp_Older.name, tp_diff)

	fmt.Printf("Of %s and %s, %s is older by %d years\n",
		bob.name, paul.name, bp_Older.name, bp_diff)

}

//
// struct的匿名字段
//
type Human1 struct {
	name   string
	age    int
	weight int
}

type Student1 struct {
	Human1     // 匿名字段，那么默认Student就包含了Human的所有字段
	speciality string
}

func test_struct_student1() {
	// 我们初始化一个学生
	mark := Student1{Human1{"Mark", 25, 120}, "Computer Science"}

	// 我们访问相应的字段
	fmt.Println("His name is ", mark.name)
	fmt.Println("His age is ", mark.age)
	fmt.Println("His weight is ", mark.weight)
	fmt.Println("His speciality is ", mark.speciality)
	// 修改对应的备注信息
	mark.speciality = "AI"
	fmt.Println("Mark changed his speciality")
	fmt.Println("His speciality is ", mark.speciality)
	// 修改他的年龄信息
	fmt.Println("Mark become old")
	mark.age = 46
	fmt.Println("His age is", mark.age)
	// 修改他的体重信息
	fmt.Println("Mark is not an athlet anymore")
	mark.weight += 60
	fmt.Println("His weight is", mark.weight)
}

func test_struct_student1_father() {
	// 我们初始化一个学生
	mark := Student1{Human1{"Mark", 25, 120}, "Computer Science"}
	mark.Human1 = Human1{"Marcus", 55, 220}
	mark.Human1.age -= 1
	// 我们访问相应的字段
	fmt.Println("His name is ", mark.Human1)
	fmt.Println("His age is ", mark.Human1.age)
}

//
// 不仅仅是struct字段哦，所有的内置类型和自定义类型都是可以作为匿名字段的。
//
type Skills []string

type Human2 struct {
	name   string
	age    int
	weight int
}

type Student2 struct {
	Human2     // 匿名字段，struct
	Skills     // 匿名字段，自定义的类型string slice
	int        // 内置类型作为匿名字段
	speciality string
}

func test_struct_student2() {
	// 初始化学生Jane
	jane := Student2{Human2: Human2{"Jane", 35, 100}, speciality: "Biology"}
	// 现在我们来访问相应的字段
	fmt.Println("Her name is ", jane.name)
	fmt.Println("Her age is ", jane.age)
	fmt.Println("Her weight is ", jane.weight)
	fmt.Println("Her speciality is ", jane.speciality)
	// 我们来修改他的skill技能字段
	jane.Skills = []string{"anatomy"}
	fmt.Println("Her skills are ", jane.Skills)
	fmt.Println("She acquired two new ones ")
	jane.Skills = append(jane.Skills, "physics", "golang")
	fmt.Println("Her skills now are ", jane.Skills)
	// 修改匿名内置类型字段
	jane.int = 3
	fmt.Println("Her preferred number is", jane.int)
}

// 如果human里面有一个字段叫做phone，而student也有一个字段叫做phone，那么该怎么办呢？
//
type Human struct {
	name  string
	age   int
	phone string // Human类型拥有的字段
}

type Employee struct {
	Human      // 匿名字段Human
	speciality string
	phone      string // 雇员的phone字段
}

func test_struct_employee() {
	Bob := Employee{Human{"Bob", 34, "777-444-XXXX"}, "Designer", "333-222"}
	fmt.Println("Bob's work phone is:", Bob.phone)
	// 如果我们要访问Human的phone字段
	fmt.Println("Bob's personal phone is:", Bob.Human.phone)
}

func main() {
	//	test_var_struct()
	//	test_init1_struct()
	//	test_init2_struct()
	//	test_init3_struct()
	//	test_struct_older()
	//	test_struct_student1()
	//	test_struct_student1_father()
	//	test_struct_student2()
	test_struct_employee()

}
