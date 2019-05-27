package main

import (
	"fmt"
	"strconv"
)

func main() {

	// 类型转化异常情况 - 字符串转整型  失败为0
	var str string = "yes"
	var num, _ = strconv.ParseInt(str, 10, 64)
	fmt.Println(num)

	// 字符串转布尔类型 - 失败为false
	var b, _ = strconv.ParseBool(str)
	fmt.Println(b)

}
