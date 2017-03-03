// 打印一行标准输入中超过一次出现的文本的出现次数
package main

import (
	"bufio"
	"fmt"
	"os"
)

func main()  {
	counts := make(map[string]int)
	input := bufio.NewScanner(os.Stdin)
	for input.Scan() {
		// line := input.Text()
		// counts[line] = counts[line] + 1
		counts[input.Text()]++
	}
	// 忽略input.Err()的潜在错误
	for line, n := range counts {
		if n > 1 {
			fmt.Printf("%d\t%s\n", n, line)
		}
	}
}