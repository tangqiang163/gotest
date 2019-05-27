package main

import (
	"fmt"
	"unsafe"
)

func main() {

	var flag = false
	// bool类型
	var flag2 bool = true

	var c1 byte = 'a'
	var c2 byte = '0'
	var c3 byte = 'g'

	// 输出Ascii编码值
	fmt.Println("c1=", c1)
	fmt.Println("c2=", c2)
	fmt.Println("c3=", c3)

	fmt.Println("11 = ", unsafe.Sizeof(flag))

	fmt.Println("flag2 = ", flag2)

	// 高精度转向低精度 - 范围大转向范围小的

}
