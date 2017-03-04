// 小的 "echo"服务器
package main
import (
	"fmt"
	"log"
	"net/http"
)

func main()  {
	// 每个请求的处理句柄
	http.HandleFunc("/", handler)
	log.Fatal(http.ListenAndServe("localhost:8000", nil))
}

// 返回请求的路径部分
func handler(w http.ResponseWriter, r *http.Request)  {
	fmt.Fprintf(w, "URL.Path = %q\n", r.URL.Path)
}