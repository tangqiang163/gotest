package main

import "fmt"

/**
  -  函数也是一种引用类型
*/
func main() {

	a := add

	fmt.Println("a的类型%T = ", a)

	var n3 = a(1, 2)

	fmt.Println("n3 = ", n3)

	fmt.Println("n4 = ", addPlus(a, 1, 2))

}

func add(n1 int, n2 int) int {
	return n1 + n2
}

/**
  函数作为参数传递
*/
func addPlus(funvar func(int, int) int, n1 int, n2 int) int {
	return funvar(n1, n2) + 4
}
