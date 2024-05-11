package main

import "fmt"

func main() {
	// 指针
	var intP *int
	intP = new(int)
	*intP = 1
	fmt.Println(*intP)

	num := 1
	var p *int
	p = &num
	fmt.Println("num变量的地址为：", p)
	fmt.Println("指针变量p的地址为：", &p)

	// 获取指针中的内容
	// 指针变量存储的值为地址值，通过在指针变量前面加上“*”符号可以获取指针所指向地址值 的内容
	fmt.Println("指针变量p的值为", *p)
	*p = 100
	fmt.Println("修改后指针变量p的值为", *p)

}
