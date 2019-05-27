package main

import "fmt"

func main() {

	fmt.Println("i plus : ", plus(10))

	// 递归算法走一波
	fmt.Println("斐波那契数:", test001(6))

	// 猴子吃桃问题
	fmt.Println("猴吃桃数:", test003(1))

	// 值传递问题 - 基本数据类型和数组
	var str string = "hello"

	fmt.Println("str point = ", &str)

	fmt.Println("str111 = ", test004(&str))

	fmt.Println("func  = ", str)

}

func test001(n int) int {

	if n == 1 || n == 2 {
		return 1
	} else {
		return test001(n-1) + test001(n-2)
	}

}

// 猴子吃桃
func test003(n1 int) int {

	if n1 == 10 {
		return 1
	} else {
		return (test003(n1+1) + 1) * 2
	}

}

/**
  指针指向的问题
*/
func test004(ptr *string) string {

	var str1 = *ptr + "11"

	// 更改指针指向的值 -
	/**
	错误写法 - 更改不了
		ptr = &str1
	*/
	*ptr = str1

	fmt.Println("ptr", ptr)

	return str1
}

func plus(i int) int {

	return i + 10
}
