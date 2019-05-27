package main

import "fmt"

/**
- 自定义类型
*/
func main() {

	type myint int
	var num1 myint
	var num2 int
	num1 = 10
	num2 = 20

	// 自定义类型  类似别名 实际上还是需要做一次类型转换
	fmt.Println("num1 + num2 = ", int(num1)+num2)

	// 函数有点类似多态 别名引用
	sum := myFun2(getSum, 20, 20)

	fmt.Println("sum = ", sum)

}

// 针对函数 - 取别名
type myFunType func(int, int) int

func myFun2(funvar myFunType, n1 int, n2 int) int {
	return funvar(n1, n2)

}

func getSum(i1 int, i2 int) int {

	return i1 + i2
}
