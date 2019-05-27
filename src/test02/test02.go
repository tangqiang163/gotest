package main

import (
	"fmt"
)

func main() {

	var a int = 10

	a = 20

	a = 40

	// 没有严格执行语言规范
	a = 12.00

	fmt.Print("a=", a)

	var i int = 10
	//i:=8
	/**
	i:=8

	等价于
	var i int = 8

	*/

	fmt.Println("i=", i)

}
