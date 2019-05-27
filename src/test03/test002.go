package main

import "fmt"

func main() {

	var n1 int32 = 20

	var n2 = int8(n1) + 127

	// 编译不通过
	//var n3 = int8(n1)+128

	fmt.Println(n2)

	//fmt.Println(n3)

}
