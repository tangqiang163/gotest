package main

import (
	"fmt"
)

func main() {

	var num1 int = 99

	var num2 float64 = 23.456

	var b bool = true

	//var myChar byte = 'h'

	var str string

	// 使用 fmt.Springtf方法 - 指明转换的类型的参数
	str = fmt.Sprintf("%d", num1)
	fmt.Println(str)

	str = fmt.Sprintf("%f", num2)
	fmt.Println(str)

	str = fmt.Sprintf("%t", b)
	fmt.Println(str)

}
