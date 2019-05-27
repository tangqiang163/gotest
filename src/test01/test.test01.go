package test01

import (
	"fmt"
)

func getVal(num1 int, num2 int) (int, int) {
	sum := num1 + num2
	sub := num1 - num2

	return sum, sub
}

func main() {
	sum, sub := getVal(30, 20)
	fmt.Print("sum=", sum, " sub=", sub)
}
