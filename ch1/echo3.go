// 打印命令行参数
package main

import (
	"fmt"
	"os"
	"strings"
)

func main()  {
	// 用Join函数代替+=操作字符串，节省内存开销
	fmt.Println(strings.Join(os.Args[0:], " "))
	// 如果不关心格式，可直接输出以显示值
	fmt.Println(os.Args[1:])
}