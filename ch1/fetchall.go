// 打印 所有网页的大小和响应时间
package main

import (
	"fmt"
	"io"
	"io/ioutil"
	"net/http"
	"os"
	"time"
	"strings"
)

func main()  {
	start := time.Now()
	ch := make(chan string)
	for _, url := range os.Args[1:] {
		// 开启一个协程
		go fetch(url, ch)
	}
	for range os.Args[1:] {
		// 从通道ch中接收数据<-ch
		fmt.Println(<-ch)
	}
	fmt.Printf("%.2fs elapsed\n", time.Since(start).Seconds())
}

func fetch(url string, ch chan<- string)  {
	start := time.Now()
	const preHTTP = "http://"
	if !strings.HasPrefix(url, preHTTP) {
		url = preHTTP + url
	}
	resp, err := http.Get(url)
	if err != nil {
		// 发送数据到通道ch <-
		ch <- fmt.Sprint(err)
		return
	}

	// Copy()返回字节数，错误信息
	nbytes, err := io.Copy(ioutil.Discard, resp.Body)
	resp.Body.Close()
	if err != nil {
		ch <- fmt.Sprintf("while reading %s: %v", url, err)
		return
	}
	secs := time.Since(start).Seconds()
	ch <- fmt.Sprintf("%.2fs\t%7d\t%s", secs, nbytes, url)
}