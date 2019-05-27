package main

import "fmt"

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

func swap(n1 *int, n2 *int) (int, int) {

	t := *n1
	*n1 = *n2
	*n2 = t

	return *n1, *n2
}
