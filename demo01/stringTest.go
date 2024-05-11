package main

import (
	"bytes"
	"fmt"
	"reflect"
	"strings"
)

/*
*
定义一个结构体
*/
type Person struct {
	Name string
	Age  int
}

func main() {
	// 字符串拼接
	a := "123546"
	b := "789"
	c := a + b
	fmt.Println(c)

	//转换为字节缓冲流
	var d bytes.Buffer
	d.WriteString(a)
	d.WriteString(b)

	fmt.Println("Buffer => ", d.String())
	fmt.Println("reflect.TypeOf(d) => ", reflect.TypeOf(d))

	// 字符串截取
	str := "GO 语言"
	index := strings.Index(str, "G")
	fmt.Println(index)
	fmt.Println("下标之后的数据", str[index:])
	fmt.Println("下标数据", string(str[index]))
	fmt.Println("下标数据", str[index])
	fmt.Println("下标之前的数据", str[:index])
	// 最后一次出现的下标
	str1 := "Go语言,Python语言"
	index1 := strings.LastIndex(str1, "语")
	fmt.Println(index1)
	fmt.Println(string(str1[index1]))

	// 修改字符串
	// Go语言无法对字符串直接进行修改，只能将字符串转换为字节数组后再进行操作
	bytes := []byte(str1)
	for i := 0; i < len(bytes); i++ {
		fmt.Printf("%X ", bytes[i])
		bytes[i] = ' '
	}
	fmt.Println(string(bytes))

	// 格式化输出 Sprintf
	var day = 1
	var hour = 24
	strFmt := fmt.Sprintf("%d天包含%d个小时", day, hour)
	fmt.Println(strFmt)
	fmt.Println("============================================================")
	// 定义一个结构体实例
	person := Person{Name: "Alice", Age: 30}
	// %v - 按值的本来值输出 {Alice 30}
	fmt.Printf("%%v: %v\n", person)
	// %+v - 在 %v 基础上，对结构体字段名和值进行展开 {Name:Alice Age:30}
	fmt.Printf("%%+v: %+v\n", person)
	// %#v - 输出 Go 语言语法格式的值
	fmt.Printf("%%#v: %#v\n", person)
	// %T - 输出 Go 语言语法格式的类型
	fmt.Printf("%%T: %T\n", person)
	// %% - 输出 % 本体
	fmt.Printf("%%%%: %%\n")
	// %b - 整型以二进制方式显示
	num := 42
	fmt.Printf("%%b: %b\n", num)
	// %o - 整型以八进制方式显示
	fmt.Printf("%%o: %o\n", num)
	// %d - 整型以十进制方式显示
	fmt.Printf("%%d: %d\n", num)

	// %x - 整型以十六进制方式显示（字母小写）
	fmt.Printf("%%x: %x\n", num)

	// %X - 整型以十六进制方式显示（字母大写）
	fmt.Printf("%%X: %X\n", num)

	// %U - Unicode 字符
	char := 'A'
	fmt.Printf("%%U: %U\n", char)

	// %f - 浮点数
	floatNum := 3.14159
	fmt.Printf("%%f: %f\n", floatNum)

	// %p - 指针，以十六进制方式显示
	ptr := &person
	fmt.Printf("%%p: %p\n", ptr)
}
