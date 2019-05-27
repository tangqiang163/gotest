package main

import "fmt"

func main() {

	// 指针 - 指向对象引用地址的变量
	var str string = "hello"

	fmt.Println("str:", &str)

	// 指针对象
	var ptr *string = &str

	fmt.Println("ptr:", ptr)

	// 更改的指针变量的指向
	var num int = 9

	// 赋值指针变量
	var ptr1 *int = &num

	// 更改指针的指向值
	*ptr1 = 10

	// 输出
	fmt.Println(10)

}
