package main

import (
	"fmt"
	"time"
)

func main() {
	start := time.Now()
	fmt.Println("hello world")
	fmt.Printf("%.4fs \n", time.Since(start).Seconds())
}
