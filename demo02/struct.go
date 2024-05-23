package main

import "fmt"

// 类型定义
type NewInt int

// 类型别名
type MyInt = int

// 定义结构体 person
type person struct {
	name string
	city string
	age  int8
}

//	type person1 struct {
//		name, city string
//		age        int8
//	}
//
// 嵌套结构体
// Address 地址结构体
type Address struct {
	Province string
	City     string
}

// User 用户结构体
type User struct {
	Name    string
	Gender  string
	Address Address
}

func main() {
	var a NewInt
	var b MyInt

	fmt.Printf("type of a:%T\n", a) //type of a:main.NewInt
	fmt.Printf("type of b:%T\n", b) //type of b:int

	// 结构体实例化
	var p1 person
	p1.name = "pp"
	p1.city = "SuQian"
	p1.age = 18
	fmt.Printf("type of p1:%T\n", p1)
	fmt.Printf("p1=%v\n", p1)
	fmt.Printf("p1=%#v\n", p1)

	// 匿名结构体
	var user struct {
		Name string
		Age  int
	}
	user.Name = "pp1111"
	user.Age = 18
	fmt.Printf("%#v\n", user)
	// 指针类型结构体
	//var p2 = new(person)
	//fmt.Printf("%T\n", p2)     //*main.person
	//fmt.Printf("p2=%#v\n", p2) //p2=&main.person{name:"", city:"", age:0}
	// 在Go语言中支持对结构体指针直接使用.来访问结构体的成员
	var p2 = new(person)
	p2.name = "测试"
	p2.age = 18
	p2.city = "北京"
	fmt.Printf("p2=%#v\n", p2)
	// &对结构体进行new实例化
	p3 := &person{}
	fmt.Printf("%T\n", p3)     //*main.person
	fmt.Printf("p3=%#v\n", p3) //p3=&main.person{name:"", city:"", age:0}
	p3.name = "博客"
	p3.age = 30
	p3.city = "成都"
	fmt.Printf("p3=%#v\n", p3) //p3=&main.person{name:"博客", city:"成都", age:30}

	var p4 person
	fmt.Printf("p4=%#v\n", p4)
	// 使用键值对初始化
	p5 := person{
		name: "pprof.cn",
		city: "北京",
		age:  18,
	}
	fmt.Printf("p5=%#v\n", p5)
	// 嵌套结构体
	user1 := User{
		Name:   "pprof",
		Gender: "女",
		Address: Address{
			Province: "黑龙江",
			City:     "哈尔滨",
		},
	}
	fmt.Printf("user1=%#v\n", user1)
}
