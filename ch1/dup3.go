// 读取文件(全部读取)，计算文件中单词的重复频率
package main

import (
	"fmt"
	"io/ioutil"
	"os"
	"strings"
)

func main()  {
	counts := make(map[string]int)
	for _, filename := range os.Args[1:] {
		// ReadFile返回字节slice，必须将其转换位string才可以使用Split
		data, err := ioutil.ReadFile(filename)
		if err != nil {
			fmt.Fprintf(os.Stderr, "dup3: %v\n", err)
			continue
		}
		for _, line := range strings.Split(string(data), "\n") {
			counts[line]++
		}
	}
	for line, n := range counts {
		if n > 1 {
			fmt.Printf("%d\t%s\n", n, line)
		}
	}
}
