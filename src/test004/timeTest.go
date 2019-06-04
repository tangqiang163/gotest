package main

import (
	"fmt"
	"time"
)

func main() {

	start := time.Now()

	fmt.Println("start: ", start)

	test()

	end := time.Now()

	fmt.Println("end: ", end)

}

func test() {

	for i := 0; i < 1000; i++ {

		fmt.Println("i: ", i)

	}

}
