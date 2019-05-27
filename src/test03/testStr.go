package main

import "fmt"

func main() {
	// 字符串不可变
	var str1 = "test"

	var str2 = " gogo"

	fmt.Println("str", str2+str1)
}
