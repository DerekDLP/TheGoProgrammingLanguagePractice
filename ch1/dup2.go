// 打印一行标准输入中超过一次出现的文本的出现次数
// 或者 从一系列文件名的列表
package main

import (
	"bufio"
	"fmt"
	"os"
)

func main()  {
	counts := make(map[string]int)
	files := os.Args[1:]
	if len(files) == 0 {
		countLines(os.Stdin, counts, "")
	} else {
		for _, arg :=range files {
			// 打开文件返回值[文件句柄，错误代码(正常返回nil)]
			f, err := os.Open(arg)
			if err != nil {
				fmt.Fprintf(os.Stderr, "dup2: %v\n", err)
				continue
			}
			countLines(f, counts, arg)
			f.Close()
		}
	}
	for line, n := range counts {
		if n > 1 {
			fmt.Printf("%d\t%s\n", n, line)
		}
	}
}

func countLines(f *os.File, counts map[string]int, fileName string) {
	input := bufio.NewScanner(f)
	for input.Scan() {
		counts[input.Text() + fileName]++
	}
	// 忽略input.Err()的潜在错误
}