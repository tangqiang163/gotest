package main

import "fmt"

func main() {

	var a = 1 >> 2

	var b = -11 >> 2

	fmt.Println("a=", a)

	fmt.Println("b=", b)

	var c = -11 << 2

	var d = -5 << 2

	fmt.Println("c=", c)

	fmt.Println("d=", d)

	// 同时为1成立
	fmt.Println("按位与", 2&3)

	// 有一位为1成立
	fmt.Println("按位或", 2|3)

	// 有一位为1 有一位为0
	fmt.Println("按位异或", 2^3)

	fmt.Println("按位异或", -2^2)

}
