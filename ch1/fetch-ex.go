// 通过URL打印对应的内容
package main

import (
	"fmt"
	"io"
	"net/http"
	"os"
	"strings"
)

func main()  {
	for _, url := range os.Args[1:] {
		// ex1.8 使用strings.HasPrefix
		const preHTTP = "http://"
		if !strings.HasPrefix(url, preHTTP) {
			url = preHTTP + url
		}
		resp, err := http.Get(url)
		// ex1.9 显示HTTP的状态信息resp.Status
		fmt.Printf("The page %s 's Status : %v\n", url, resp.Status)
		if err != nil {
			fmt.Fprintf(os.Stderr, "fetch: %v\n", err)
			os.Exit(1)
		}
		// ex1.7使用io.Copy(dst,src)代替ioutil.ReadAll
		b, err := io.Copy(os.Stdout, resp.Body)
		// 关闭，避免内存泄露
		resp.Body.Close()
		if err != nil {
			fmt.Fprintf(os.Stderr, "fetch: reading %s: %v\n", url, err)
			os.Exit(1)
		}
		fmt.Printf("%s", b)
	}
}