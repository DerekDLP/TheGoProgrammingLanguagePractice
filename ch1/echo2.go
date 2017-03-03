// 打印命令行参数
package main

import (
	"fmt"
	"os"
	"strconv"
)

func main()  {
	s, sep := "", ""
	for n, arg := range os.Args[1:] {
		// int 转 string
		// str2 := fmt.Sprintf("%d", n)
		sep = " "
		temp := strconv.Itoa(n)
		s += sep + temp + "," + sep + arg + "\n"
	}
	fmt.Println(s)
}