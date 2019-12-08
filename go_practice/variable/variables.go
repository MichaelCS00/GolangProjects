package main

import "fmt"

/**
 *	变量声明的 5 种方式
 */

// 1.全局变量声明可以不在本文件内使用 列表方式
var (
	first  int
	second string
	third  float32
)

func main() {

	// 局部变量声明后必须使用

	// 2.短变量声明 只能函数内局部变量声明
	fifth := " "
	// 3.默认初始化
	var fifth1 string
	// 4.
	var fifth2 = ""
	// 5.
	var fifth3 string = ""

	// 实际使用中较多使用前两种声明方式

	fmt.Print(fifth)
}
