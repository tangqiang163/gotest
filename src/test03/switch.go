package main

import "fmt"

func main() {

	/**
	 * swicht
	 */

	var key = "h"

	switch key {

	case "a", "h":
		fmt.Println("aaa")
	case "b":
		fmt.Println("bbb")
	case "c":
		fmt.Println("ccc")
	case "d":
		fmt.Println("ddd")
	default:
		fmt.Println("other")

	}

	for i := 1; i < 10; i++ {
		fmt.Println("hello,go!")
	}

}
