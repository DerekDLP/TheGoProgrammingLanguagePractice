// 有多种方法可以格式化结构
// 本示例为最简单的一种：以一个固定宽度打印每个issue

// 演示示例 go run issues.go repo:golang/go is:open json decoder
//         查询Go语言项目中和JSON解码相关的问题

// GitHub的Web服务接口 https://developer.github.com/v3/

package main

import (
	"fmt"
	"log"
	"os"

	"GoBook/TheGoProgrammingLanguagePractice/ch4/json/github"
)

func main() {
	result, err := github.SearchIssues(os.Args[1:])
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("%d issues:\n", result.TotalCount)
	for _, item := range result.Items {
		fmt.Printf("#%-5d %9.9s %.55s\n",
			item.Number, item.User.Login, item.Title)
	}
}
