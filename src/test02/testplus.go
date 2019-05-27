package main

import "fmt"

/**

	基本数据类型
			数值类型int int8 int16 int32 int64
                   uint uint8 uint32 uint64 byte

            浮点类型float32 float64

            字符型 没有专门的字符行，使用byte来保存单个字母字符




*/
func main() {

	var a = 10
	var b = 20
	var r = a + b
	// 使用byte保存单个字母字符
	var byte = '1'

	var str1 = " hello "
	var str2 = " golang "
	var sum = str1 + str2

	fmt.Println("r = ", r)
	fmt.Println("sum = ", sum)

}
