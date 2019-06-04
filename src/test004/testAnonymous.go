package main

import "fmt"

/**
  1.使用匿名函数的两种方式
  2.使用全局匿名函数
*/
func main() {

	/**
	在函数定义的时候 直接赋值调用
	*/
	res1 := func(n1 int, n2 int) int {
		return n1 + n2
	}(10, 20)

	fmt.Println("sum1 = ", res1)
	fmt.Println("sum2 = ", res1)

	fmt.Println("----------------------")

	/**
	在函数定义的时候 使用引用变量获取函数

	操作函数类型的引用
	*/
	f := func(n1 int, n2 int) int {
		return n2 - n1
	}

	fmt.Println("sub1 = ", f(1, 100))
	fmt.Println("sub2 = ", f(50, 100))

	fmt.Println("乘积 = ", Fun1(5, 12))

}

// 全局匿名函数 -e
var (
	Fun1 = func(n1 int, n2 int) int {
		return n1 * n2
	}
)
