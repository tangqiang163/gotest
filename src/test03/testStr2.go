package main

import "fmt"

func main() {

	var str = "abc\nabc"

	fmt.Println(str)

	var str2 = `just \n test`

	fmt.Println(str2)

	var str3 = "hello world"

	str3 += " sixceo"

	fmt.Println(str3)

}
