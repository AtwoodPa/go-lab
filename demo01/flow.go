package main

import "fmt"

func main() {
	num := -10
	// 基本的 if-else 语句
	if num > 0 {
		fmt.Println("num 是正数")
	} else if num == 0 {
		fmt.Println("num 是零")
	} else {
		fmt.Println("num 是负数")
	}

	// if 语句中可以先执行一个语句再判断
	if x := 5; x < 0 {
		fmt.Println("x 是负数")
	} else if x == 0 {
		fmt.Println("x 是零")
	} else {
		fmt.Println("x 是正数")
	}

	// 基本的 for 循环
	for i := 0; i < 5; i++ {
		fmt.Println(i)
	}

	// for 循环作为 while 使用
	j := 0
	for j < 5 {
		fmt.Println(j)
		j++
	}

	// 无限循环
	k := 0
	for {
		if k >= 5 {
			break
		}
		fmt.Println(k)
		k++
	}

	// 使用 range 遍历数组或切片
	nums := []int{1, 2, 3, 4, 5}
	for index, value := range nums {
		fmt.Printf("索引 %d 对应的值是 %d\n", index, value)
	}

	// 使用 range 遍历 map
	person := map[string]int{"Alice": 30, "Bob": 25, "Charlie": 35}
	for key, value := range person {
		fmt.Printf("%s 的年龄是 %d 岁\n", key, value)
	}

	weekday := "Monday"

	// 基本的 switch-case 语句
	switch weekday {
	case "Monday":
		fmt.Println("星期一")
	case "Tuesday":
		fmt.Println("星期二")
	case "Wednesday", "Thursday":
		fmt.Println("星期三或星期四")
	default:
		fmt.Println("其他日期")
	}

	// switch 语句中可以不带表达式，相当于 if-else 的用法
	switch {
	case weekday == "Monday":
		fmt.Println("今天是星期一")
		fallthrough // 强制执行后一条case
	case weekday == "Friday":
		fmt.Println("今天是星期五")
	default:
		fmt.Println("今天不知道是星期几")
	}

	// goto
	fmt.Print("Hello")
	goto sign
	fmt.Println("123123") // 这段代码不执行，跳过
sign:
	fmt.Println(" goto")
}
