package main

import (
	"fmt"
	"strconv"
)

func main() {

	var num3 int = 99
	//var num4 float64 = 23.456
	//var b2 bool = true

	var str string

	str = strconv.FormatInt(int64(num3), 10)

	fmt.Println(str)

	var str1 string = "15"

	var mynum int64

	// pasreInt返回两个值
	mynum, _ = strconv.ParseInt(str1, 10, 64)

	fmt.Println(mynum)

	var str2 string = "true"

	var b bool

	b, _ = strconv.ParseBool(str2)

	fmt.Println(b)

}
