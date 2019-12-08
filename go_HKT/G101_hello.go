// 从命令行接收一个输入参数并补到输出结果的最后
package main

import (
	"fmt"
	"os"
)

func main() {
	if len(os.Args) > 1 {
		fmt.Println("hello world " + os.Args[2])
	}
}
