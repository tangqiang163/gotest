package main

import "fmt"

func main() {

	/**
	  使用for模仿while的效果
	*/

	i := 1

	for {
		if i > 10 {
			fmt.Println("合格！")
			break
		} else {
			// 栈内存遍历
			fmt.Println("遍历中！！ i=", &i)
			i++
		}
	}

	// goto   break   continue

}
