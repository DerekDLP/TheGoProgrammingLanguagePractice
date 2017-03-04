// 小的 "echo"服务器 + 计数功能
package main
import (
	"fmt"
	"log"
	"net/http"
	"sync"
)

var (
	mu sync.Mutex
	count int
)

func main()  {
	// 每个请求的处理句柄
	http.HandleFunc("/", handler)
	http.HandleFunc("/count", counter)
	log.Fatal(http.ListenAndServe("localhost:8000", nil))
}

// 返回请求的路径部分
func handler(w http.ResponseWriter, r *http.Request)  {
	mu.Lock()
	count++
	mu.Unlock()
	fmt.Fprintf(w, "URL.Path = %q\n", r.URL.Path)
}

// 记录服务器被请求的次数
func counter(w http.ResponseWriter, r *http.Request)  {
	mu.Lock()
	fmt.Fprintf(w, "Count %d\n", count)
	mu.Unlock()
}