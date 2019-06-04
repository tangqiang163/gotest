package main

import "fmt"

/**
  常用于释放资源
*/
func main() {

	deferfunc()
}

func deferfunc() {

	// 暂不执行 会压入等待栈中 等待函数执行完毕 然后再从栈中按照先进后出的原则执行
	defer fmt.Println("n1 = ", 100)
	defer fmt.Println("n2 = ", 200)

	fmt.Println("n3 = ", 300)

}
