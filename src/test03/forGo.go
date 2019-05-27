package main

import "fmt"

func main() {

	for i := 1; i < 10; i++ {
		fmt.Println("遍历开始 i=:", i)
	}

	var team [3]string
	team[0] = "jjj"
	team[1] = "kkk"
	team[2] = "lll"

	// 数据遍历
	for idx, val := range team {
		fmt.Println("index:", idx)
		fmt.Println("val:", val)

	}

}
