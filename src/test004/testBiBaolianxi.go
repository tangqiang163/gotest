package main

import (
	"fmt"
	"strings"
)

/**

闭包练习

*/
func main() {

	// 闭包处理
	a := judge(".jpg")

	fmt.Println("判断处理", a("a.jpg"))
	fmt.Println("判断处理", a("b.e"))

}

func judge(suffix string) func(string) string {

	return func(name string) string {

		if !strings.HasSuffix(name, suffix) {

			return name + suffix

		}

		return name
	}
}
