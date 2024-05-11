package main

import "fmt"

// APP 常量声明可以不使用，但是变量一定要使用
const APP string = "EE" // 显示声明格式
const VERSION = "1.1"   // 隐式声明格式
// 常量组
const (
	a = 3.14
	b
	c
	d = 100
	e
)

// 定义枚举类型
type Weekday int

// 定义枚举值
const (
	Sunday    Weekday = iota // 0
	Monday                   // 1
	Tuesday                  // 2
	Wednesday                // 3
	Thursday                 // 4
	Friday                   // 5
	Saturday                 // 6
)

func main() {

	fmt.Println(APP, VERSION)

	fmt.Println("a:", a) // 输出显式赋值的常量 a 的值
	fmt.Println("b:", b) // 输出默认值，即前一个常量 a 的值（3.14）
	fmt.Println("c:", c) // 输出默认值，即前一个常量 b 的值（3.14）
	fmt.Println("d:", d) // 输出显式赋值的常量 d 的值
	fmt.Println("e:", e)

	// go实现枚举
	fmt.Println(Sunday)  // 输出: Sunday
	fmt.Println(Monday)  // 输出: Monday
	fmt.Println(Tuesday) // 输出: Tuesday

	// 使用枚举类型作为函数参数
	weekday := Wednesday
	fmt.Println("Today is", weekday) // 输出: Today is Wednesday

	a := 9
	b := 13

	// 按位与 (&)：对应位上的两个数都为1时，结果位为1，否则为0
	fmt.Println("a&b =", a&b) // 输出: 9 & 13 = 9 (二进制 1001 & 1101 = 1001)

	// 按位或 (|)：对应位上的两个数只要有一个为1，结果位就为1，否则为0
	fmt.Println("a|b =", a|b) // 输出: 9 | 13 = 13 (二进制 1001 | 1101 = 1101)

	// 按位异或 (^)：对应位上的两个数不相同时，结果位为1，否则为0
	fmt.Println("a^b =", a^b) // 输出: 9 ^ 13 = 4 (二进制 1001 ^ 1101 = 0100)

	// 左移 (<<)：将一个数的各二进制位全部左移若干位，高位丢弃，低位补0
	fmt.Println("a<<2 =", a<<2) // 输出: 9 << 2 = 36 (二进制 1001 左移2位变为 100100，即十进制的36)

	// 右移 (>>)：将一个数的各二进制位全部右移若干位，保持符号位不变
	fmt.Println("b>>2 =", b>>2) // 输出: 13 >> 2 = 3 (二进制 1101 右移2位变为 11，即十进制的3)
}

func (d Weekday) String() string {
	names := [...]string{"Sunday", "Monday", "Tuesday", "Wednesday", "Thursday", "Friday", "Saturday"}
	if d < Sunday || d > Saturday {
		return "Unknown"
	}
	return names[d]
}
