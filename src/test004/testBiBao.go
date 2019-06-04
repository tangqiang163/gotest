package main

import "fmt"

// 闭包
func main() {

	f1 := add()

	f2 := add()

	fmt.Println("sum = ", f1(30))
	fmt.Println("sum = ", f1(10))
	fmt.Println("sum = ", f1(10))

	fmt.Println("sum = ", f2(10))
	fmt.Println("sum = ", f2(10))
}

// 只在函数定义的时候使用了初始化函数
func add() func(int) int {

	var n int = 10

	fmt.Println("执行了！")

	return func(x int) int {
		n = n + x
		return n
	}

}
