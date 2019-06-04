package main

import "fmt"

/**
  import优先

  函数执行的顺序
	1.全局变量
	2.init()方法
	3.main()方法
*/
func main() {

	fmt.Println("sum = ", sum(10, 15, 20))

	var num1 = 10
	var num2 = 20

	point1 := &num1
	point2 := &num2

	num3, num4 := swap(point1, point2)

	fmt.Println("num3=", num3)
	fmt.Println("num4=", num4)

}

// 可变参数
func sum(n1 int, args ...int) int {

	sum := n1
	for i := 0; i < len(args); i++ {

		sum += args[i]

	}

	return sum
}

/**
更换指针所指向的地址
*/
func swap(n1 *int, n2 *int) (int, int) {

	t := *n1
	*n1 = *n2
	*n2 = t

	return *n1, *n2

}

/**
在执行main方法前 会首先执行init进行初始化
*/
func init() {

	fmt.Println("init 初始化 ~ ")

}

// 定义全局变量
var age int = test()

/**
  在执行test方法
*/
func test() int {

	fmt.Println("全局变量赋值！")

	return 100
}
