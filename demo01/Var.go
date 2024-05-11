package main

import (
	"fmt"
	"reflect"
)

func main() {
	fmt.Println("========================================")
	var num1 int
	fmt.Printf("num1: %d\n", num1)
	var num int = 10
	// 打印输出地址
	fmt.Printf("内存地址：%p\n", &num)

	// 使用 fmt.Printf 打印输出
	fmt.Printf("num的值为：%d\n", num)

	// 使用 fmt.Println 打印输出
	fmt.Println("num的值为：", num)
	fmt.Println("========================================")
	// 批量创建变量
	var (
		a int
		b string
		c float32
		d float64
		e bool
	)
	// 打印输出 a b c d e
	fmt.Println(a)
	fmt.Println(b)
	fmt.Println(c)
	fmt.Println(d)
	fmt.Println(e)
	// 编译器推导格式
	fmt.Println("========================================")
	var age = 20
	fmt.Println(age)
	// 短变量声明初始化
	fmt.Println("========================================")
	temp := 123
	fmt.Println(temp)

	name, age := "Tom", 18
	fmt.Println("name:", name)
	fmt.Println("age:", age)

	var name2 string
	name2 = "盼盼"
	fmt.Println("name2:", name2)

	fmt.Println("========================================")
	fmt.Println(result())
	a, _ = result()
	fmt.Println(a)
	// float32、float64的区别
	// float32浮点数的最大范围约为3.4e38，float64浮点数 最大范围约为1.8e308
	var tempf float64
	tempf = 2.0 / 3.0
	fmt.Println(tempf)
	tempf = float64(float32(tempf))
	fmt.Println(tempf)
	fmt.Println("变量tempf的类型:", reflect.TypeOf(tempf))

	fmt.Println("========================================")
	str1 := "你好：Hello"
	fmt.Println(str1)
	// 多行字符串
	str2 := `这是第一行
第二行。。。
第三行。。。
`
	fmt.Println(str2)
	// 字符和转义字符
	english := 'a'
	china := '我'
	fmt.Println(english, china)
	// 使用转义字符创建带有换行和制表符的字符串
	fmt.Println("第一行\n第二行\t制表符")

	// 使用转义字符表示双引号和反斜杠
	fmt.Println("这是一个双引号：\"")
	fmt.Println("这是一个反斜杠：\\")

	// 布尔
	var flag bool
	flag = false
	if flag {
		fmt.Println("true")
	} else {
		fmt.Println("false")
	}

	// 数据类型推断
	var sum int = 1
	fmt.Println("sum的数据类型为:", reflect.TypeOf(sum))
	// 类型转换
	var aa int16 = 65
	fmt.Println("变量a值为：", aa, ", 变量类型为：", reflect.TypeOf(aa))
	// 将 int16 类型的变量 a 转换为 int32 类型的变量 b
	bb := int32(aa)
	// 输出变量 b 的值和类型
	fmt.Println("变量b值为：", bb, ", 变量类型为：", reflect.TypeOf(bb))

	// 将 int32 类型的变量 b 转换为 string 类型（注意：直接将整数转换为 string 并不会得到预期的字符）
	// 需要将数字转换为对应的 Unicode 字符来获取对应的字符
	str3 := string(rune(bb))
	// 输出转换后的字符（注意：这里输出的是字符而不是数字的字符串表示）
	fmt.Println("转换变量b类型为string：", str3)

}

/*
*
匿名变量
*/
func result() (int, int) {
	return 10, 20
}
